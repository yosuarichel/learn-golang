package main

import (
	"fmt"
	"learn-golang/helpers"
)

func main() {
	helpers.Greet("Yosua Richel", 24)

	person := helpers.Person{}

	person.Invokegreet()

	var c1 helpers.Shape = helpers.Circle{
		Radius: 5,
	}
	var r1 helpers.Shape = helpers.Rectangle{
		Width:  10,
		Height: 10,
	}
	fmt.Printf("Type of c1 = %T\n", c1)
	fmt.Printf("Type of r1 = %T\n", r1)

	fmt.Println("Circle area", c1.Area())
	fmt.Println("Circle perimter", c1.Perimeter())

	fmt.Println("Rectangle area", r1.Area())
	fmt.Println("Rectangle perimter", r1.Perimeter())

	// type assertion
	c1.(helpers.Circle).Volume()

	value, ok := c1.(helpers.Circle)

	if ok {
		fmt.Println("Circle value", value)
		fmt.Println("Circle volume", value.Volume())
	}

}
