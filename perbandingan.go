package main

import "fmt"

func main() {
	var name string = "yosua"
	var name2 string = "richel"

	var perbandingan = name == name2
	fmt.Println("perbandingan: ", perbandingan)

	var lebihBesar = name > name2
	fmt.Println("lebih besar: ", lebihBesar)

}
