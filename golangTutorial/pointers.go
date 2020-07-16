package main

import (
	"fmt"
)

func main() {
	/* b := 255
	a := &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a) */

	/* a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}

	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size) */

	/* b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b) */

	/* a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)

	d := hello()
	fmt.Printf("type: %T\tvalue: %v", d, *d) */

	// slice is recomendded for this kind of scenerio
	a := [3]int{89, 90, 91}
	modify(&a)
	fmt.Println(a)
}

func modify(arr *[3]int) {
	arr[0] = 90
}

func hello() *int {
	i := 5
	return &i
}

func change(val *int) {
	*val = 55
}
