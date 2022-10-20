package main

import (
	"fmt"

	"github.com/bochanikc-calculator-project/calculator"
)

func main() {
	fmt.Printf("Start calculator\n")
	result, err := calculator.DoOperation(5, 6, "+")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	result, err = calculator.DoOperation(5, 6, "-")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
