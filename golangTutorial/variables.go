package main

import "fmt"

func main() {
	// variable declarition
	/* i := 1.1
	fmt.Printf("type: %T, value: %v\n", i, i)
	i = 10
	fmt.Printf("type: %T, value: %v\n", i, i)
	i = 20
	fmt.Printf("type: %T, value: %v\n", i, i)

	width, height := 10, 10
	fmt.Printf("width: %v, height: %v\n", width, height) */

	// declare variables of different types
	/* var (
		name  = "PS4"
		price = 3000
	)
	fmt.Printf("the price for %v is %v\n", name, price) */

	// short hand declaration
	/* name1, price1 := "switch", 3000
	fmt.Printf("the price for %v is %v\n", name1, price1) */

	// can not assign 1 value to 2 variables
	// name2, price2 := "switch"

	// at least 1 unassigned variable should be in the left hand of ":=
	a, b := 20, 30
	fmt.Println(a, b)
	b, c := 40, 50
	fmt.Println(b, c)
	// b, c := 40, 50
}
