package main

import (
	"fmt"
	"strings"
)

func main() {
	// Colsure on variable
	evenNumbers := func(numbers ...int) []int {
		var result []int
		for _, element := range numbers {
			if element%2 == 0 {
				result = append(result, element)
			}
		}
		return result
	}

	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 15, 16,
	}
	fmt.Println("Even numbers: ", evenNumbers(numbers...))

	// Closure in variable with direct invoke
	isPlindrome := func(str string) bool {
		temp := ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp == str
	}("kataks")
	fmt.Println("is palindrome: ", isPlindrome)

	// Closure as a return value
	students := []string{
		"Yosua",
		"Richel",
		"SImanjuntak",
	}
	var find = findStudent(students)
	fmt.Printf(find("yosuas"))

	numbersOdd := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}

	findOdd := findOddNumber(numbersOdd, func(number int) bool {
		return number%2 == 0
	})
	fmt.Println("Total odd number = ", findOdd)

}

// Closure as a return value
func findStudent(students []string) func(string) string {

	return func(s string) string {
		var student string
		var position int

		for index, value := range students {
			if strings.ToLower(value) == strings.ToLower(s) {
				student = value
				position = index
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s does not exist\n", s)
		}
		return fmt.Sprintf("found %s at position %d\n", s, position+1)
	}
}

func findOddNumber(numbers []int, callback func(int) bool) int {
	var totalOddNumber int

	for _, value := range numbers {
		if callback(value) {
			totalOddNumber++
		}
	}
	return totalOddNumber
}
