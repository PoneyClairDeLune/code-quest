use std::io::{stdin, stdout, Stdin, Write};

fn print(string: &str) {
    // Just to print a string out
    print!("{}", string);
    let _ = stdout().flush();
}

fn scanln(source: &Stdin) -> String {
    // Return a single line of user input
    let mut input = String::new();
    source.read_line(&mut input).expect("Incorrect string.");
    input.truncate(input.trim_end_matches(&['\r', '\n'][..]).len()); // Suggested by Mg138
    return input;
}

fn main() {
    let stdin: Stdin = stdin();
    let mut tree_house: Vec<u8> = vec![];
    let mut tree_count: u16 = 0;
    loop {
        print("Input: ");
        let input: String = scanln(&stdin);
        if input.len() > 0 {
            let char_len = input.len() as u16;
            if tree_count < char_len {
                tree_count = char_len;
            };
            for tree in input.bytes() {
                tree_house.push(tree);
            }
        } else {
            break;
        };
    }
    /* let mut i0: usize = 0;
    while i0 < tree_house.len() {
        print!("{}", String::from_utf8(vec![tree_house[i0]]).unwrap());
        if (i0 + 1) % usize::from(tree_count) == 0 {
            print("\n");
        };
        i0 += 1;
    }; */
    let mut observable: Vec<bool> = vec![];
    // Raycasting from the sides, with left and top first
    let mut max_left: u8 = 0;
    let mut max_top: Vec<u8> = vec![];
    let lc = usize::from(tree_count); // Local count
    let mut i0: usize = 0;
    print("First raycast started.\n");
    while i0 < tree_house.len() {
        let x = i0 % lc;
        let y = i0 / lc;
        let e0 = tree_house[i0];
        let mut observed: bool = false;
        if x == 0 {
            max_left = e0;
            observed = true;
        };
        if y == 0 {
            max_top.push(e0);
            observed = true;
        };
        if x > 0 && y > 0 && x < (lc - 1) && y < (lc - 1) {
            //print!("Testing observability r {} c {}...", y, x);
            if e0 > max_left {
                observed = true;
                max_left = e0;
                //print(" left");
            };
            if e0 > max_top[x] {
                observed = true;
                max_top[x] = e0;
                //print(" top");
            };
            //print("\n");
        };
        observable.push(observed);
        i0 += 1;
    }
    // Then right and bottom
    print("Second raycast started.\n");
    let mut max_right: u8 = 0;
    let mut max_bottom: Vec<u8> = vec![];
    while i0 > 0 {
        i0 -= 1;
        let x = lc - 1 - i0 % lc; // Reverse X
        let y = lc - 1 - (i0 / lc); // Reverse Y
        let e0 = tree_house[i0];
        if x == 0 {
            observable[i0] = true;
            max_right = e0;
        };
        if y == 0 {
            observable[i0] = true;
            max_bottom.push(e0);
        };
        if x > 0 && y > 0 && x < (lc - 1) && y < (lc - 1) {
            //print!("Testing observability r {} c {}...", y, x);
            if e0 > max_right {
                max_right = e0;
                observable[i0] = true;
                //print(" right");
            };
            if e0 > max_bottom[x] {
                max_bottom[x] = e0;
                observable[i0] = true;
                //print(" bottom");
            };
            //print("\n");
        };
    }
    print("Total trees observable from the sides: ");
    let mut observed: u16 = 0;
    for e0 in observable {
        if e0 {
            observed += 1;
        };
    }
    println!("{}", observed);
    let mut leftward: Vec<u32> = vec![];
    let mut topward: Vec<u32> = vec![];
    let mut rightward: Vec<u32> = vec![];
    let mut bottomward: Vec<u32> = vec![];
    let mut y_cache: Vec<u32> = vec![];
    let mut x_nearest: [u32; 10] = [0; 10];
    let mut y_nearest: Vec<[u32; 10]> = vec![];
    // Third raycasting with cached results, from left and top
    print("Third raycasting started.\n");
    while i0 < tree_house.len() {
        let x = i0 % lc;
        let y = i0 / lc;
        let e0 = u32::from(tree_house[i0]);
        let mut x_store: u32 = 0;
        let mut y_store: u32 = 0;
        if y_nearest.len() <= x {
            y_nearest.push([0; 10]);
        };
        if x == 0 {
            let mut i1: usize = 0;
            while i1 < x_nearest.len() {
                x_nearest[i1] = 0;
                i1 += 1;
            }
            x_nearest[usize::from(tree_house[i0] - 48)] = x as u32;
        } else {
            if e0 > u32::from(tree_house[i0 - 1]) {
                let mut start_height: usize = (e0 as usize) - 48;
                let mut nearest_block: u32 = 0;
                while start_height < 10 {
                    if nearest_block < x_nearest[start_height] {
                        nearest_block = x_nearest[start_height];
                    };
                    start_height += 1;
                }
                x_store = (x as u32) - nearest_block;
            } else {
                x_store = 1;
            };
        };
        x_nearest[usize::from(tree_house[i0] - 48)] = x as u32;
        if y == 0 {
            y_cache.push(0);
            y_nearest[x][usize::from(tree_house[i0] - 48)] = y as u32;
        } else {
            if e0 > u32::from(tree_house[i0 - lc]) {
                let mut start_height: usize = (e0 as usize) - 48;
                let mut nearest_block: u32 = 0;
                while start_height < 10 {
                    if nearest_block < y_nearest[x][start_height] {
                        nearest_block = y_nearest[x][start_height];
                    };
                    start_height += 1;
                }
                y_cache[x] = (y as u32) - nearest_block;
            } else {
                y_cache[x] = 1;
            };
            y_store = y_cache[x];
        };
        y_nearest[x][usize::from(tree_house[i0] - 48)] = y as u32;
        leftward.push(if x_store > 0 { x_store } else { 1 });
        topward.push(if y_store > 0 { y_store } else { 1 });
        i0 += 1;
    }
    // Fourth raycasting with cached results, from left and top
    print("Fourth raycasting started.\n");
    while i0 > 0 {
        i0 -= 1; // Real index
        let x = lc - 1 - i0 % lc; // Reverse X
        let y = lc - 1 - (i0 / lc); // Reverse Y
        let e0 = u32::from(tree_house[i0]);
        let mut x_store: u32 = 0;
        let mut y_store: u32 = 0;
        if x == 0 {
            let mut i1: usize = 0;
            while i1 < x_nearest.len() {
                x_nearest[i1] = 0;
                i1 += 1;
            }
            x_nearest[usize::from(tree_house[i0] - 48)] = x as u32;
        } else {
            if e0 > u32::from(tree_house[i0 + 1]) {
                let mut start_height: usize = (e0 as usize) - 48;
                let mut nearest_block: u32 = 0;
                while start_height < 10 {
                    if nearest_block < x_nearest[start_height] {
                        nearest_block = x_nearest[start_height];
                    };
                    start_height += 1;
                }
                x_store = (x as u32) - nearest_block;
            } else {
                x_store = 1;
            };
        };
        x_nearest[usize::from(tree_house[i0] - 48)] = x as u32;
        if y == 0 {
            y_cache[x] = 0;
            let mut i1: usize = 0;
            while i1 < 10 {
                y_nearest[x][i1] = 0;
                i1 += 1;
            }
            y_nearest[x][usize::from(tree_house[i0] - 48)] = y as u32;
        } else {
            if e0 > u32::from(tree_house[i0 + lc]) {
                let mut start_height: usize = (e0 as usize) - 48;
                let mut nearest_block: u32 = 0;
                while start_height < 10 {
                    if nearest_block < y_nearest[x][start_height] {
                        nearest_block = y_nearest[x][start_height];
                    };
                    start_height += 1;
                }
                y_cache[x] = (y as u32) - nearest_block;
            } else {
                y_cache[x] = 1;
            };
            y_store = y_cache[x];
        };
        y_nearest[x][usize::from(tree_house[i0] - 48)] = y as u32;
        rightward.push(if x_store > 0 { x_store } else { 1 });
        bottomward.push(if y_store > 0 { y_store } else { 1 });
    }
    // Debugging window
    /* let mut i4: usize = 0;
    println!("Casting from left:");
    while i4 < tree_house.len() {
        print!("{0:0>2}", leftward[i4]);
        if (i4 + 1) % usize::from(tree_count) == 0 {
            print("\n");
        } else {
            print(" ");
        };
        i4 += 1;
    }
    i4 = 0;
    println!("Casting from top:");
    while i4 < tree_house.len() {
        print!("{0:0>2}", topward[i4]);
        if (i4 + 1) % usize::from(tree_count) == 0 {
            print("\n");
        } else {
            print(" ");
        };
        i4 += 1;
    }
    i4 = 0;
    println!("Casting from right:");
    while i4 < tree_house.len() {
        print!("{0:0>2}", rightward[tree_house.len() - i4 - 1]);
        if (i4 + 1) % usize::from(tree_count) == 0 {
            print("\n");
        } else {
            print(" ");
        };
        i4 += 1;
    }
    i4 = 0;
    println!("Casting from bottom:");
    while i4 < tree_house.len() {
        print!("{0:0>2}", bottomward[tree_house.len() - i4 - 1]);
        if (i4 + 1) % usize::from(tree_count) == 0 {
            print("\n");
        } else {
            print(" ");
        };
        i4 += 1;
    } */
    // Scenic score calculation
    let mut highest_score: u32 = 0;
    //println!("Final result:");
    while i0 < tree_house.len() {
        let r_index = tree_house.len() - i0 - 1;
        let score = leftward[i0] * topward[i0] * rightward[r_index] * bottomward[r_index];
        /* print!("{0:0>2}", score);
        if (i0 + 1) % usize::from(tree_count) == 0 {
            print("\n");
        } else {
            print(" ");
        }; */
        if score > highest_score {
            highest_score = score;
        };
        i0 += 1;
    }
    println!("Highest scenic score: {}", highest_score);
}
