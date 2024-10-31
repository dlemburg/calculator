package Prompt

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

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

type PromptState struct {
	lastInput          string
	lastInputSanitized string
	currentInput       string
	result             string
}

type Prompt struct {
	state *PromptState
}

// 	state *CalculatorState

func isOperator(value string) bool {
	if value == string(Divide) || value == string(Multiply) || value == string(Subtract) || value == string(Add) {
		return true
	}

	return false
}

func convertStringToArray(r string) []string {
	return strings.Split(r, " ")
}

func NewPrompt() *Prompt {
	return &Prompt{}
}

func (p *Prompt) Run() []string {
	// this is how i'd love for the prompt to show up
	options := []string{
		"Calculate", "Clear",
		"+", "-", "*", "/",
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	result := ""

	p.state = &PromptState{
		lastInput:          "",
		lastInputSanitized: "",
		currentInput:       "",
		result:             "",
	}

	for {
		prompt := promptui.Select{
			Label:        "Select an option",
			Items:        options,
			Size:         16,
			HideSelected: true,
		}

		_, promptResult, err := prompt.Run()

		/**
		* validate and append to result
		**/

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
						// subtract operator
					} else {
						result = result + " " + promptResult + " "
					}
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
