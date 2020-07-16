package main

import (
	"fmt"
)

func main() {
	/* var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	personSalary1 := make(map[string]int)
	personSalary1["steve"] = 12000
	personSalary1["jamie"] = 15000
	personSalary1["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary1) */

	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)

	/* employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])
	// zero value is returned for an element that is not present
	fmt.Println("Salary of joe is", personSalary["joe"])

	newEmp := "joe"
	value, ok := personSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	for k, v := range personSalary {
		fmt.Printf("key: %v\t value: %v\n", k, v)
	}

	fmt.Println("map before deletion", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)
	fmt.Println("length is", len(personSalary)) */

	// Similar to slices, maps are reference types.
	// When a map is assigned to a new variable, they both point to the same internal data structure.
	// Hence changes made in one will reflect in the other.
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)

}
