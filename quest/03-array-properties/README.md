## Quest 03: Array Properties
### Task
Write a program to ask for values, filling up a 10-element array with criteria met.

### Details
* Ask users for 10 numbers to fill the array up.
* After filling up the 10 figures, the program will print all 10 within a single line.
* The program will find both the value and the location of the maximum value.
* The program will print the mean of this array out.

### Requirement
* The elements should be no less than 0, and no more than 100. If the user entered a value exceeding this range, ask them to enter again.
* You're required to implement the actual functionality in four seperate functions below.
```
void tableFill(char targetArray[], char length);
void tablePrint(char targetArray[], char length);
void printMax(char targetArray[], char length);
void printMean(char targetArray[], char length);
```
* Example output:
```
Provide figure #01 (0~100): 1
Provide figure #02 (0~100): 6
Provide figure #03 (0~100): 3
Provide figure #04 (0~100): 6
Provide figure #05 (0~100): 3
Provide figure #06 (0~100): 7
Provide figure #07 (0~100): 0
Provide figure #08 (0~100): 5
Provide figure #09 (0~100): 7
Provide figure #10 (0~100): 4
Among the array: 1 6 3 6 3 7 0 5 7 4
There exists a maximum of 7 at #06
And a mean of 4.2
```