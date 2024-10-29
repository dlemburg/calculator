package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

// negative
// gofmt
// separate prompt and calculator into different packages
// REST api
// history with redis
// goroutines

type Operator string

const (
	Add       Operator = "+"
	Subtract  Operator = "-"
	Divide    Operator = "/"
	Multiply  Operator = "*"
	Calculate Operator = "Calculate"
	Clear     Operator = "Clear"
	Delete    Operator = "Delete"
)

func isOperator(value string) bool {
	if value == string(Divide) || value == string(Multiply) || value == string(Subtract) || value == string(Add) {
		return true
	}

	return false
}

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

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Exec(input []string) int {
	var current string
	var next string
	orderedOperations := make([]string, 0)

	c.input = input

	for i := 0; i < len(c.input); i++ {
		current = c.input[i]

		if i+1 < len(c.input) {
			next = c.input[i+1]
		} else {
			next = ""
		}

		// if priority operator, immediately do operation and append result into orderedOperations
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

func convertStringToArray(r string) []string {
	return strings.Split(r, " ")
}

func (c *Calculator) Prompt() []string {
	// this is how i'd love for the prompt to show up
	options := []string{
		"Calculate", "Clear",
		"+", "-", "*", "/",
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
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
			return convertStringToArray(result)
		}

		if promptResult == string(Calculate) {
			return convertStringToArray(result)
		}

		// TODO: convert to switch
		if promptResult == string(Clear) {
			result = ""
		} else if isOperator(promptResult) {
			// operators not allowed if there's no input
			if promptResult == string(Subtract) {
				// first character: negative operator
				if len(result) == 0 {
					result = "-"
				} else {
					// negative operator
					if len(result) > 2 && isOperator(result[len(result)-2:len(result)-1]) {
						result = result + "-"
					}

					// subtract operator
					result = result + " " + promptResult + " "
				}
				// if current result is empty, do not allow non (-) operator
			} else if len(result) == 0 {
				result = ""
				// append sanitized operator to result
			} else {
				result = result + " " + promptResult + " "
			}
		} else {
			result = result + promptResult
		}

		// keep result inlined
		fmt.Print("\033[u") // restore the cursor position
		fmt.Printf("\n> %q\n", result)
	}
}

func main() {
	calc := NewCalculator()
	input := calc.Prompt()
	result := calc.Exec(input)

	fmt.Println(result)
}
