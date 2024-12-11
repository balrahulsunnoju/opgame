/*
Package main implements a calculator that finds all possible expressions
from a given set of integers that evaluate to a target number. The input
should be a series of integers where the last number is the target.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var operators = []string{"+", "-", "*", "/"}

// Entry point
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue // Skip empty lines
		}
		nums, err := parseInput(line)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1) // Exit with 1 for invalid input
		}
		solutions := findSolutions(nums)
		if len(solutions) == 0 {
			fmt.Println() // Print an empty line if no solutions found
		} else {
			sort.Strings(solutions) // Ensure solutions are sorted
			fmt.Println(strings.Join(solutions, ", "))
		}
	}
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1) // Exit with 1 for reading errors
	}
}

// Parses input from a line of text into integers
func parseInput(input string) ([]int, error) {
	fields := strings.Fields(input)
	if len(fields) < 2 {
		return nil, fmt.Errorf("not enough numbers provided")
	}
	nums := make([]int, len(fields))
	for i, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, fmt.Errorf("invalid input: %s", field)
		}
		nums[i] = num
	}
	return nums, nil
}

// Recursively finds all valid solutions
func findSolutions(nums []int) []string {
	if len(nums) < 3 {
		return nil // Need at least 2 operands and 1 target
	}
	target := nums[len(nums)-1] // The last number is the target result
	nums = nums[:len(nums)-1]   // The rest are operands

	var solutions []string
	search(nums, strconv.Itoa(nums[0]), nums[0], 1, target, &solutions)
	return solutions
}

// Recursive search for all possible valid expressions
func search(nums []int, expr string, current int, index int, target int, solutions *[]string) {
	if index == len(nums) {
		if current == target {
			*solutions = append(*solutions, expr+" = "+strconv.Itoa(target))
		}
		return
	}

	for _, op := range operators {
		nextExpr := fmt.Sprintf("%s %s %d", expr, op, nums[index])
		nextValue := applyOperator(current, nums[index], op)
		if nextValue != nil {
			search(nums, nextExpr, *nextValue, index+1, target, solutions)
		}
	}
}

// Applies the given operator to two operands and returns the result (if valid)
func applyOperator(a, b int, op string) *int {
	switch op {
	case "+":
		result := a + b
		return &result
	case "-":
		result := a - b
		return &result
	case "*":
		result := a * b
		return &result
	case "/":
		if b != 0 && a%b == 0 { // Ensure integer division
			result := a / b
			return &result
		}
	}
	return nil
}
