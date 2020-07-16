package main

import (
	"fmt"
)

func main() {
	/* var a [3]int
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	b := [3]int{1, 2}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4} // ... makes the compiler determine the length
	fmt.Println(c) */

	// not possible since [3]int and [4]int are distinct types
	// c = a

	// Arrays in Go are value types and not reference types.
	// This means that when they are assigned to a new variable,
	// a copy of the original array is assigned to the new variable.
	// If changes are made to the new variable, it will not be reflected in the original array.
	/* a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	fmt.Println("length of a is", len(a))
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v] = %v\n", i, a[i])
	}

	for i, v := range a {
		fmt.Printf("a[%v] = %v\n", i, v)
	} */

	/* a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printArray(a)
	var b [3][2]string
	b[0][0] = "java"
	b[0][1] = "golang"
	b[1][0] = "python"
	b[1][1] = "sql"
	b[2][0] = "c"
	b[2][1] = "c++"
	printArray(b) */

	/* a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:4]
	fmt.Println(b)
	c := []int{1, 2, 3}
	fmt.Println(c)

	// A slice does not own any data of its own.
	// It is just a representation of the underlying array.
	// Any modifications done to the slice will be reflected in the underlying array.
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr) */

	// When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array.
	/* numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	fmt.Println("nums2 after modification to slice nums1", nums2)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
	fmt.Println("nums1 after modification to slice nums2", nums1)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d, capacity %d\n", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                         //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), ", capacity is", cap(fruitslice))

	i := make([]int, 5, 10)
	fmt.Println(i)
	fmt.Printf("length of slice %d, capacity %d\n", len(i), cap(i)) */

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		fmt.Println("old length", len(names), "old capacity", cap(names))
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
		fmt.Println("new length", len(names), "and capacity", cap(names))
	}

	// It is also possible to append one slice to another using the ... operator.
	names = append(names, cars...)
	fmt.Println(names)

	// When a slice is passed to a function, even though it's passed by value,
	// the pointer variable will refer to the same underlying array.
	// Hence when a slice is passed to a function as parameter,
	// changes made inside the function are visible outside the function too.
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside

	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

// use the copy function func copy(dst, src []T) int to make a copy of that slice.
// this way we can use the new slice and the original array can be garbage collected.
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 1
	}
}

func printArray(a [3][2]string) {
	for _, v := range a {
		for _, v1 := range v {
			fmt.Printf("%v ", v1)
		}
		fmt.Println()
	}
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}
