package main

import "fmt"

func main() {
	// 1
	person := map[string]string{}

	person["name"] = "yosua"
	person["age"] = "24"
	person["gender"] = "male"

	fmt.Println("person: ", person)

	// 2
	var user = map[string]string{
		"name":   "yosua",
		"age":    "24",
		"gender": "male",
	}

	fmt.Println("user: ", user)
	fmt.Println("user[name]: ", user["name"])
	fmt.Println("user[age]: ", user["age"])
	fmt.Println("user[gender]: ", user["gender"])

	// Looping with Map
	for key, value := range user {
		fmt.Printf("key %s, with value %s \n", key, value)
	}

	// Delete map element
	delete(person, "gender")
	fmt.Println("after deleteing: ", person)

	// Detecting a value
	value, exist := person["gender"]
	fmt.Printf("value %s, is exist %v \n", value, exist)
	if exist {
		fmt.Println("Value: ", value)
	} else {
		fmt.Println("Element not found")
	}

	// Combining array and map
	cars := []map[string]string{
		{"name": "ferrari", "color": "white"},
		{"name": "mustang", "color": "black"},
		{"name": "ducati", "color": "yellow"},
	}
	fmt.Println("cars: ", cars)

	for index, element := range cars {
		fmt.Printf("index %v with value name: %s, color: %s ", index, element["name"], element["color"])
	}
}
