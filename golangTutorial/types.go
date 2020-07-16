package main

import "fmt"

func main() {
	// bool
	/* a, b := true, false
	fmt.Println(a, b)
	fmt.Println(a && b)
	fmt.Println(a || b) */

	// int
	/* var m int = 10
	n := 20
	fmt.Printf("%T %v\n", m, m)
	fmt.Printf("%T %v\n", n, n) */

	// float
	/* a, b := 12.34, 23.45
	fmt.Printf("%T %T\n", a, b)
	fmt.Println(a + b)
	fmt.Println(a - b) */

	// complex
	/* a := complex(1, 2)
	b := 3 + 4i
	sum := a + b
	// (ac - bd) + (ad + bc)i
	mul := a * b
	fmt.Printf("%T %v\n", sum, sum)
	fmt.Printf("%T %v\n", mul, mul) */

	// string
	/* item, price := "PS4", "3000"
	detail := "price for " + item + " is " + price
	fmt.Println(detail) */

	// need to convert type explicitly
	a := 10
	b := 10.5
	sum := a + int(b)
	fmt.Println(sum)
	var c float64 = float64(a)
	fmt.Println(c)
}
