package main

import (
	"fmt"
	"strconv"
)

func main() {

	array := []int{
		5, 4, 3, 2, 1,
	}
	n := len(array)

	result := ``
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			result += strconv.Itoa(j+(n-i)) + " "
		}
		result += "\n"
	}
	fmt.Println(result)

	// loop through the element
	fruits := []string{
		"apel",
		"mangga",
		"durian",
	}
	for i, v := range fruits {
		fmt.Printf("Index: %v, Value: %s \n", i, v)
	}
}
