package main

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

// cli stdin
// stdout: stdin = answer
// example:
// ["9", +, 2, *, 2, +, 3, *, 4]
// [9, +, 4, +, ]
// 9 + 2 * 2 + 3 * 4
// 9 + 4 + 12

// types:
// Parser Object -> Parse, Validate; Operations
// Calculator Object -> Add, Subtract, Multiply, Divide, Clear; CurrentValue
// History Map -> key: operation, value: result

// split the string into an array of strings
// 2 loops for operations
// loop through and find * or / and do the operation on the previous and next values and push to new array
// loop through and find + or - and do each operation
// return the result

// input from cli
// REST api
// goroutines

var input = []string{"1", "+", "2", "*", "3", "+", "-3"}

type Operator string

const (
	Add      Operator = "+"
	Subtract Operator = "-"
	Divide   Operator = "/"
	Multiply Operator = "*"
)

func isPriorityOperator(value string) bool {
	if value == string(Divide) || value == string(Multiply) {
		return true
	}

	return false
}

func isNonPriorityOperator(value string) bool {
	if value == string(Add) || value == string(Subtract) {
		return true
	}

	return false
}

type Calculator struct {
	input []string
}

func NewCalculator(input []string) *Calculator {
	return &Calculator{input: input}
}

func (c *Calculator) Exec() int {
	var current string
	var next string
	orderedOperations := make([]string, 0)

	for i := 0; i < len(input); i++ {
		current = input[i]

		if i+1 < len(input) {
			next = input[i+1]
		} else {
			next = ""
		}

		// if priority operator, immediately do operation
		if isPriorityOperator(current) {
			lastIdx := len(orderedOperations) - 1
			lastInt, _ := strconv.Atoi(orderedOperations[lastIdx])
			nextInt, _ := strconv.Atoi(next)

			if current == string(Multiply) {
				orderedOperations[lastIdx] = strconv.Itoa(c.Multiply(lastInt, nextInt))
			} else if current == string(Divide) {
				orderedOperations[lastIdx] = strconv.Itoa(c.Divide(lastInt, nextInt))
			}

			// we looked ahead already to complete the above operation, so skip the next element
			i++
		} else {
			orderedOperations = append(orderedOperations, current)
		}
	}

	var counter int
	var operator string

	for i, el := range orderedOperations {
		if i == 0 {
			counter, _ = strconv.Atoi(el)
			continue
		}

		if isNonPriorityOperator(el) {
			operator = el
			continue
		}

		nextInt, err := strconv.Atoi(el)

		if err != nil {
			panic(err)
		}

		if operator == string(Add) {
			counter = c.Add(counter, nextInt)
		} else if el == string(Subtract) {
			counter = c.Subtract(counter, nextInt)
		}
	}

	fmt.Println(orderedOperations)

	return counter
}

func (c *Calculator) Add(values ...int) int {
	counter := 0
	for _, x := range values {
		counter += x
	}

	return counter
}

func (c *Calculator) Subtract(values ...int) int {
	counter := 0
	for _, x := range values {
		counter -= x
	}

	return counter
}

func (c *Calculator) Multiply(values ...int) int {
	counter := values[0]
	for i := 1; i < len(values); i++ {
		counter *= values[i]
	}

	return counter
}

func (c *Calculator) Divide(values ...int) int {
	counter := values[0]
	for i := 1; i < len(values); i++ {
		counter /= values[i]
	}

	return counter
}

func prompt() {
	// TODO
	options := []string{"Calculate", "Clear", "Del", "+", "-", "*", "/", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	result := ""

	for {
		prompt := promptui.Select{
			Label:        "Select an option",
			Items:        options,
			Size:         16,
			HideSelected: true,
		}

		_, promptResult, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if promptResult == "Calculate" {
			fmt.Println("Exiting...")
			break
		}

		result = result + promptResult

		fmt.Print("\033[u") // restore the cursor position
		fmt.Printf("Operation: %q\n", result)
	}
}

func main() {
	calc := NewCalculator(input)
	result := calc.Exec()

	fmt.Println(result)
}
