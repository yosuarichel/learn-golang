package main

import "fmt"

func main() {
	var months = []string{}
	fmt.Println("original month: ", months)
	monthAppend := append(months, "baru")
	fmt.Println("append month: ", monthAppend)

	// var slice1 = months[0:6]
	// fmt.Println("slice1: ", slice1)
	// fmt.Println("slice len: ", len(slice1))
	// fmt.Println("slice capacity: ", cap(slice1))

	days := [...]string{
		"minggu",
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
	}
	fmt.Println("days original: ", days)
	daysSlice := days[1:5]
	fmt.Println("day slice original: ", daysSlice)
	daysSlice[0] = "jumat baru"
	daysSlice[1] = "sabtu baru"
	fmt.Println("day slice edited: ", daysSlice)
	fmt.Println("days edited: ", days)

	append := append(daysSlice, "test")
	fmt.Println("append: ", append)
	fmt.Println("appended day slice: ", daysSlice)
	fmt.Println("days: ", days)

	string := string(rune(0))
	fmt.Println("string: ", string)

}
