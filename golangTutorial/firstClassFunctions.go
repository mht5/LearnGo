package main

import (
	"fmt"
)

/* type add func(a int, b int) int

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func simpleReturn() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
} */

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	/* a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a)

	func() {
		fmt.Println("hello world first class function")
	}()

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers") */

	/* var a add = func(a, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)

	simple(a)

	f := simpleReturn()
	fmt.Println(f(60, 7)) */

	// Closures are anonymous functions which access the variables defined outside the body of the function.
	/* a := 5
	func(b int) {
		fmt.Printf("a = %v, b = %v\n", a, b)
	}(6) */

	// Every closure is bound to its own surrounding variable.
	// a and b are bound to their own value of t.
	/* a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!")) */

	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
}
