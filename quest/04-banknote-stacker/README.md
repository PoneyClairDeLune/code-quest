## Quest 04: Banknote Stacker
### Task
You're owner of a small business operating in Manehatten. There is an experimental economical reform taking place, with paper banknotes substituting the good old bits.

The Equestrian central bank haven't decided which face values to set on however, and the size of your business is varying. To pay your employees off efficiently, you decided to develop a simple but adaptive program to cram out the exact number of banknotes you need.

#### Details
* The face values of notes avaliable now are: 500, 200, 100, 50, 20, 10, 5, 2 and 1.
* You currently have 5 employees.
* The program will ask the salary of each employee, then show the calculated results in the console.
* Example output:
```
Salary of employee #1: 203
Salary of employee #2: 884
Salary of employee #3: 75
Salary of employee #4: 167
Salary of employee #5: 679

  Salary  500  200  100   50   20   10    5    2    1
     203    0    1    0    0    0    0    0    1    1
     884    1    1    1    1    1    1    0    2    0
      75    0    0    0    1    1    0    1    0    0
     167    0    0    1    1    0    1    1    1    0
     679    1    0    1    1    1    0    1    2    0
   Total    2    2    3    4    3    2    3    6    1
```

#### Requirements
* The number of employees, and the exact face values must be able to change.