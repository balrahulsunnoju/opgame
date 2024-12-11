 OpGame

![Math operators](https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/Basic_arithmetic_operators.svg/330px-Basic_arithmetic_operators.svg.png)

## Overview

In this assignment you will write a Go program to solve a simple mathematical puzzle.  The program must efficiently find the binary operators that complete the equation.  For this problem the only operators you need to consider are addition, subtraction, multiplication, and division (when the divisor is non-zero and the remainder is zero).  For example, given `1 □ 2 = 3` the only operator that works to put in the square is "+".  So `1 + 2 = 3` solves the problem.

For simplicity we will not use the normal operator precedence.  Instead there will be no precedence thus expressions are evaluated strictly from left to right.  For example, given `1 □ 2 □ 3 = 1`, your goal is to find two operators that satisfy the equation.  The solution would be `1 + 2 / 3 = 1`.  Notice that this is computed as `(1+2) / 3 = 1` and not `1 + (2/3)`.

Generating your own example problems is easy.  Write an equation (being careful to evaluate it strictly left to right) and then remove the operators and supply that to your program.

## Learning Objectives

- Writing functions
- Tree searching
- Recursion

## Requirements

Write a Go program to do the following:

- Reads problems from STDIN as a list of positive integers separated by whitespace, one problem per line.
- Output a solution to the problem to STDOUT that makes the equation valid.  Each +, -, *, /, and = must be passed by one space on each side.
- If no solution is found simply output a blank line to STDOUT.
- Find **all** solutions to each problem.
- When outputting multiple solutions, separate the solutions by a comma followed by a space, in sorted increasing order as strings (see `sort.Strings()` or `slices.Sort()`).
- Errors in parsing STDIN exit with a non-zero status code after writing the error to STDERR.
- **Your program must use recursion to search for solutions.**  Specifically your program should break the problem down into a set of smaller problems to solve.  This is known as [Dynamic Programming](https://en.wikipedia.org/wiki/Dynamic_programming).
- Do not use any library other than the Go standard library.
- The source code must compile with the most recent version of the Go compiler.
- The program must not panic under any circumstances.
- Make sure your code is "gofmt'd".  See "gofmt" or better use "goimports" or better yet configure IDE to do this formatting on file save.

## Example Output

```shell
$ echo "3 1 2" | go run ./cmd/opgame
3 - 1 = 2

$ echo "5 4 2 22" | go run ./cmd/opgame
5 * 4 + 2 = 22

$ (echo "3 1 2" && echo "9 2 18") | go run ./cmd/opgame
3 - 1 = 2
9 * 2 = 18

$ echo "6 2 3 4" | go run ./cmd/opgame
6 * 2 / 3 = 4

$ echo "7 3 3 7" | go run ./cmd/opgame
7 * 3 / 3 = 7, 7 + 3 - 3 = 7, 7 - 3 + 3 = 7

$ go run ./cmd/opgame < testdata/all/basic_in.txt
3 - 1 = 2

9 + 0 = 9, 9 - 0 = 9
2 * 3 = 6
4 + 5 = 9
4 + 5 + 6 = 15
2 * 3 * 1 = 6, 2 * 3 / 1 = 6, 2 + 3 + 1 = 6
1 * 1 * 1 = 1, 1 * 1 / 1 = 1, 1 + 1 - 1 = 1, 1 - 1 + 1 = 1, 1 / 1 * 1 = 1, 1 / 1 / 1 = 1
```

## Dynamic Programming

[Dynamic Programming](https://en.wikipedia.org/wiki/Dynamic_programming) is a useful technique for finding solutions to the class of problem in this assignment.  The basic idea is that you successively break the proble down into smaller probelm(s) until they are small enough to actually solve.  Then you build back up the solution to the original problem.  Below is an example of how dynamic programming can be used for this assignment.

```txt
3 + 2 - 4 = 1   # original problem
3 2 4 1         # problem 1
 +              # action 1, recurse
 5 4 1          # problem 2
  -             # action 2, recurse
  1 1           # tautology -> solution found, return
                # back up the tree to construct the solution
  5 - 4 = 1     # reconstruct the problem, solution to problem 2, return
3 + 2 - 4 = 1   # solution to problem 1, return
```

If instead you terminate at a contradiction (e.g., 2 = 1) then disregard (i.e., prune the leaf of the tree).

Dynamic programming is an efficient method of finding solutions.  It is much more efficient than exhaustive grid search over the entire solution space.  Function memoization may be used to make dynamic programming even more efficient.  Specifically one would memoize a function that solves the problems such that the same sub-problem would never need to be solved twice.  For this assignment, any solver that is more efficient than searching over the entire solution space is considered "efficient" and meets the requirements.  

## Hints

- Commit and push often so you do not loose any of your work.
- Use the `Makefile`.  For example, you can run tests with `make test -B`.  To make sure your code is properly formatted run `make fix`.


