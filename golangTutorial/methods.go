package main

import (
	"fmt"
	_ "math"
)

/* type Employee struct {
	name     string
	salary   int
	currency string
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
} */

type Employee struct {
	name string
	age  int
}

// Method with value receiver
func (e Employee) changeName(newName string) {
	e.name = newName
}

// changes made inside a method with a pointer receiver is visible to the caller whereas this is not the case in value receiver
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

// To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package
/* func (a int) add(b int) {
} */

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

/* func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// displaySalary() method has Employee as the receiver type
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// displaySalary() method converted to function with Employee as parameter
func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
} */

func main() {
	/* emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type
	displaySalary(emp1)

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area()) */

	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
	e.changeAge(52)
	fmt.Printf("\nEmployee age after change: %d\n", e.age)

	// ethods belonging to anonymous fields of a struct can be called as if they belong to the structure where the anonymous field is defined.
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}

	p.fullAddress() //accessing fullAddress method of address struct

	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("\nSum is", sum)
}
