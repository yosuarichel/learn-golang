package main

import "fmt"

func main() {
	var bilangan int32 = 1000
	var bilangan2 int64 = int64(bilangan)
	var bilangan3 int8 = int8(bilangan)

	fmt.Println("bilangan: ", bilangan)
	fmt.Println("bilangan2: ", bilangan2)
	fmt.Println("bilangan3: ", bilangan3)

	var name string = "yosua"
	var element byte = name[0]
	fmt.Println("element: ", element)

	var formatElement string = string(element)
	fmt.Println("format element: ", formatElement)
}
