package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Yosua"
	names[1] = "Richel"

	fmt.Println("names: ", names)
	fmt.Println("name0: ", names[0])
	fmt.Println("name1: ", names[1])
	fmt.Println("name2: ", names[2])

	var array = [3]string{
		"Yosua",
		"Richel",
	}
	fmt.Println("array: ", array)

	fmt.Println("len array: ", len(array))

	// multi dimensional array
	arrays := [][]interface{}{{"string 1", 2, true}, {4, 5, 6}}
	fmt.Println("arrays: ", arrays)

	for _, element1 := range arrays {
		// fmt.Printf("index %v \n", i)
		for j, element2 := range element1 {
			fmt.Printf("index %v, element %v \n", j, element2)
		}
		fmt.Println()
	}
}
