package helpers

import "fmt"

func constantFunc() {
	const name string = "yosua richel"
	fmt.Printf("name: %v\n", name)

	const (
		dob = "09-11-1997"
		age = 25
	)
	fmt.Printf("dob: %v\n", dob)
	fmt.Printf("age: %v\n", age)
}
