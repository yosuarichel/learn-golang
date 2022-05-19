package helpers

import (
	"fmt"
	"strings"
)

func pointerFunc() {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("value of firstnumber = ", firstNumber)
	// Use & to get memory address from declared variable
	fmt.Println("memory address of firstnumber = ", &firstNumber)
	// Use * to get value of variable from variable that copy another variable memory address
	fmt.Println("value of secondnumber = ", *secondNumber)
	// To get memory address use the variable name instead if the variable has copied the memory address from another variable
	fmt.Println("memory address of secondnumber = ", secondNumber)

	fmt.Println("\n", strings.Repeat("=", 30))

	var firstPerson string = "john"
	var secondPerson *string = &firstPerson

	fmt.Println("value of firstperson = ", firstPerson)
	fmt.Println("memory address of firstperson = ", &firstPerson)
	fmt.Println("value of secondperson = ", *secondPerson)
	fmt.Println("memory address of secondperson = ", secondPerson)

	fmt.Println("\n", strings.Repeat("=", 30))

	*secondPerson = "Doe"

	fmt.Println("value of firstperson = ", firstPerson)
	fmt.Println("memory address of firstperson = ", &firstPerson)
	fmt.Println("value of secondperson = ", *secondPerson)
	fmt.Println("memory address of secondperson = ", secondPerson)

	// Pointer as parameter
	var number int = 100
	fmt.Println("before change = ", number)
	changeValue(&number)
	fmt.Println("after change = ", number)
}

// Pointer as paramter
func changeValue(number *int) {
	*number = 20
}
