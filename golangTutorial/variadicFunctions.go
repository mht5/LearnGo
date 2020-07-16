package main

import (
	"fmt"
)

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
	nums := []int{89, 90, 95}
	find(89, nums...)

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
	fmt.Println(len(welcome), cap(welcome))

}

func change(s ...string) {
	s[0] = "Go"
	// the original capacity is 2, append() will create a new array and a slice of that array with capacity 4
	// so the changes after append() will not be reflected on the original array
	s = append(s, "playground")
	s[1] = "rocks"
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
