package main

import (
	"fmt"

	calc "github.com/dlemburg/calculator/internal/calculator"
	cli "github.com/dlemburg/calculator/internal/prompt"
)

// implement calculator state
// separate prompt and calculator into different packages
// REST api
// history with redis
// goroutines
// web ui with htmx
// tests

func main() {
	p := cli.NewPrompt()
	input := p.Run()

	c := calc.NewCalculator()
	result := c.Exec(input)

	fmt.Println(result)
}
