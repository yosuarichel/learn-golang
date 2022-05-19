package helpers

import "fmt"

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Person struct {
	name string
	age  int
}

// Embeded struct
type Employee struct {
	person   Person
	division string
}

func structFunc() {
	var employee Employee

	// Giving value into a struct
	employee.person.name = "Yosua Richel"
	employee.person.age = 24
	employee.division = "Software Engineer"

	fmt.Println("employee name = ", employee.person.name)
	fmt.Println("employee age = ", employee.person.age)
	fmt.Println("employee division = ", employee.division)
	fmt.Printf("employee structure = %+v \n", employee)

	// Initializing as struct
	employee1 := Employee{}
	employee1.person.name = "Yosua Richel"
	employee1.person.age = 24
	employee1.division = "Software Engineer"

	employee2 := Employee{
		person: Person{
			name: "Yosua Richel",
			age:  25,
		},
		division: "Software Engineer",
	}

	fmt.Printf("employee1 structure = %+v \n", employee1)
	fmt.Printf("employee2 structure = %+v \n", employee2)

	// Pinter to a struct
	structEmployee := &employee1

	structEmployee.person.name = "Simanjuntak"

	fmt.Println("employee1 name = ", employee1.person.name)
	fmt.Println("employee struct name = ", structEmployee.person.name)

	// Anonymous struct without filed insertion
	employee3 := struct {
		person   Person
		division string
	}{}
	employee3.person.name = "Yosua"
	employee3.person.age = 24
	employee3.division = employee.division

	fmt.Printf("employee 3 = %+v\n", employee3)

	// Anonymous struct with filed insertion
	employee4 := struct {
		person   Person
		division string
	}{
		person: Person{
			name: "Yosua Richel",
			age:  25,
		},
		division: "Software Engineer",
	}

	fmt.Printf("employee 4 = %+v\n", employee4)

	// Slice of struct
	employeeSlice := []Employee{
		{
			person: Person{
				name: "Richel",
				age:  24,
			},
			division: "Backend Engineer",
		},
		{
			person: Person{
				name: "Yosua",
				age:  25,
			},
			division: "Frontend Engineer",
		},
		{
			person: Person{
				name: "Simanjuntak",
				age:  26,
			},
			division: "Software Engineer",
		},
	}

	for index, element := range employeeSlice {
		fmt.Printf("employee with index %d = %+v\n", index, element)
	}

	fmt.Println("employee 4 = ", employeeSlice)

	// Slice of anonymous struct
	anonymousStructSlice := []struct {
		person   Person
		division string
	}{
		{
			person: Person{
				name: "Richel",
				age:  24,
			},
			division: "Backend Engineer",
		},
		{
			person: Person{
				name: "Yosua",
				age:  25,
			},
			division: "Frontend Engineer",
		},
		{
			person: Person{
				name: "Simanjuntak",
				age:  26,
			},
			division: "Software Engineer",
		},
	}
	for index, element := range anonymousStructSlice {
		fmt.Printf("employee anon with index %d = %+v\n", index, element)
	}

}
