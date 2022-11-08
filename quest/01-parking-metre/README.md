## Quest 01: Parking Metre
### Task
With increasing demand of vehicles in Equestria like wagons and carriges, more and more parking lots are being built. However, they need an efficient way to calculate how much fee to charge.

In this quest, you're asked to plot a CLI program out, calculating charges for parking metres inside these parking lots.

#### Details
* For the first three hours of parking, the charge is always 2 bits.
* If over three hours, a single bit will be charged every two exceeding hours.
* The maximum charge allowed is 10 bits, and each vehicle can only be parked for at most 24 hours.
* The program designed will ask parking time for three customers, calculate the charges of each, then show the results in the format shown below.
```
Customer   Hours   Charge
       1    2.00     2.00
       2   11.00     6.00
       3   20.00    10.00
   TOTAL   33.00    18.00
```

#### Requirements
* Calculation of charges must be implemented in a seperate function named `calculateFee`.
* The function accepts an `f64`(`double`) value as time parked in hours, and returns an `f64`(`double`) value as fee need to be charged.
* The results must be shown as a correctly-aligned table.

### Target
In this lesson, you'll be able to learn...
1. To declare a dedicated funtion, reusing code whenever needed.
1. Return values in a function using the `return` statement, and the behaviour of returning a value in a function.
1. `if`, `else if` and `else` statements.
1. `for` and/or `while` conditional loops.
1. (C) Basic output formatting with `printf`, and input normalizing with `scanf`.
1. (C, Go, Rust, ...) Understanding pointers and the use of them.
1. (C, Go, Rust, ...) Some of the basic types within your language of choice.