package helpers

import "fmt"

func conditionFunc() {
	currentYear := 2022
	bornYear := 2005

	// if else condition
	if age := currentYear - bornYear; age < 17 {
		fmt.Println("blm bisa buat ktp")
	} else if age >= 17 && age < 21 {
		fmt.Println("gk boleh beli miras")
	} else {
		fmt.Println("bebas")
	}

	// switch case condition 1
	age := currentYear - bornYear
	switch age {
	case 17:
		fmt.Println("bisa buat ktp")
	case 21:
		fmt.Println("bisa beli miras")
	default:
		{
			fmt.Println("other")
		}
	}

	// switch case condition 2
	switch {
	case age < 17:
		fmt.Println("blm bisa buat ktp")
	case age >= 17 && age < 21:
		fmt.Println("blm bisa beli miras")
	}

}
