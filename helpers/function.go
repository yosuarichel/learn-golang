package helpers

import (
	"fmt"
	"math"
	"strings"
)

func functionFunc() {
	Greet("Yosua Richel", 24)
	address("Yosua", "Bukit Kencana")
	names := []string{
		"Yosua",
		"Richel",
	}
	hello := hello("Haii", names)
	fmt.Println(hello)

	diameter := 10
	area, around := circle(float64(diameter))
	fmt.Printf("area: %v, around: %v \n", area, around)

	side := 10
	squareArea, squareAround := square(float64(side))
	fmt.Printf("area: %v, around: %v \n", squareArea, squareAround)

	students := show("Yosua", "Richel", "Simanjuntak")
	fmt.Println("List of students: ", students)

	// Variadic function with slice data type input
	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}
	fmt.Println("sum: ", sum(numbers...))

	foods := []string{
		"indomie",
		"pizza",
		"bakpao",
	}
	fmt.Println(profile("yosua", foods...))
}

// Basic function
func Greet(name string, age int8) {
	fmt.Printf("Hei, my name is %s, i'm %v years old\n", name, age)
}

// Basic function
func greet() {
	fmt.Println("ini dari function greet")
}

// Shorthand for same data type of arguments
func address(name, address string) {
	fmt.Printf("Hei, my name is %s, i live in %v\n", name, address)
}

// Return value with data specific data type
func hello(msg string, names []string) string {
	joinName := strings.Join(names, " ")
	result := fmt.Sprintf("%s %s", msg, joinName)

	return result
}

// Return multiple value
func circle(diameter float64) (float64, float64) {
	area := math.Pi * math.Pow(diameter/2, 2)
	around := math.Pi * diameter

	return area, around
}

// Predifined return value
func square(side float64) (squareArea float64, squareAround float64) {
	squareArea = side * side
	squareAround = 4 * side
	return
}

// Variadic function
func show(names ...string) []map[string]string {
	var result []map[string]string

	for index, value := range names {
		key := fmt.Sprintf("student%d", index)
		tmp := map[string]string{
			key: value,
		}
		result = append(result, tmp)
	}

	return result
}

// Variadic function
func sum(numbers ...int) int {
	result := 0

	for _, value := range numbers {
		result += value
	}

	return result
}

// Variadic function with multiple arguments
func profile(name string, foods ...string) string {
	joinFoods := strings.Join(foods, ",")

	result := fmt.Sprintf("hello my name is %s and i love eat %s", name, joinFoods)
	return result
}
