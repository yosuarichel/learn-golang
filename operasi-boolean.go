package main

import "fmt"

func main() {
	var value1 = 100
	var value2 = 200

	fmt.Println("perbandingan &&: ", value1 > value2 && value1 == value2)
	fmt.Println("perbandingan || ", value1 > value2 || value1 == value2)
	fmt.Println("perbandingan ! ", value1 != value2)
}
