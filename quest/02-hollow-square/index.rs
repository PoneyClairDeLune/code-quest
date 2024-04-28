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

const USE_CHARS: [&str; 2] = ["*", " "];

fn used_char(size: u16, line: u16, col: u16) -> String {
	let crit: u16 = size - 1;
	if line == 0 || col == 0 || line == crit || col == crit {
		return String::from(USE_CHARS[0]);
	} else {
		return String::from(USE_CHARS[1]);
	};
}

fn main() {
	let stdin: Stdin = stdin();
	let mut size: u16;
	loop {
		print("Length of the square's side (1~20): ");
		size = scanln(&stdin).parse::<u16>().unwrap();
		if size > 0 && size < 21 {
			break;
		};
	};
	let mut line: u16 = 0;
	while line < size {
		let mut col: u16 = 0;
		while col < size {
			print(&used_char(size, line, col));
			col += 1;
		};
		print("\n");
		line += 1;
	};
}
