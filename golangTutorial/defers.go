package main

import (
	"fmt"
	"sync"
)

var num = 1

/* type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func finished() {
	fmt.Println("Finished finding largest")
}

// Defer statement is used to execute a function call just before the surrounding function where the defer statement is present returns.
// like finnaly? no, defer function is executed after the value is returned
func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
} */

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func increNum() {
	num += 1
	fmt.Println("new value:", num)
}

func getNum() int {
	defer increNum()
	fmt.Println("old value:", num)
	return num
}

func main() {
	/* nums := []int{78, 109, 2, 563, 300}
	largest(nums)

	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")

	// The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual function call is done.
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("\nvalue of a before deferred function call", a) */

	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")

	fmt.Println(getNum())
}
