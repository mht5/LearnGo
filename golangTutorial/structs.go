package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Address struct {
	city  string
	state string
}

type Person struct {
	name string
	age  int
	Address
}

type name struct {
	firstName string
	lastName  string
}

type image struct {
	data map[int]int
}

func main() {

	//creating struct specifying field names
	/* emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	//creating struct without specifying field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Printf("type: %T\tvalue: %v\n", emp1, emp1)
	fmt.Printf("type: %T\tvalue: %v\n", emp2, emp2)

	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Printf("type: %T\tvalue: %v\n", emp3, emp3)

	var emp4 Employee //zero valued struct
	fmt.Println("First Name:", emp4.firstName)
	fmt.Println("Last Name:", emp4.lastName)
	fmt.Println("Age:", emp4.age)
	fmt.Println("Salary:", emp4.salary)

	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("First Name:", emp5.firstName)
	fmt.Println("Last Name:", emp5.lastName)
	fmt.Println("Age:", emp5.age)
	fmt.Println("Salary:", emp5.salary)

	emp6 := Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)
	emp6.salary = 6500
	fmt.Printf("New Salary: $%d\n", emp6.salary)

	emp7 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("First Name:", emp7.firstName)
	fmt.Println("Age:", emp7.age) */

	/* p1 := Person{
		string: "naveen",
		int:    50,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int) */

	/* p := Person{
	       name: "Naveen",
	       age:  50,
	       address: Address{
	           city:  "Chicago",
	           state: "Illinois",
	       },
	   }

	   fmt.Println("Name:", p.name)
	   fmt.Println("Age:", p.age)
	   fmt.Println("City:", p.address.city)
	   fmt.Println("State:", p.address.state) */

	// promoted fields can be accessed as if they are directly declared in the Person struct itself.
	/* p := Person{
		name: "Naveen",
		age:  50,
		Address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)
	fmt.Println("State:", p.state) */

	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name4 := name{
		firstName: "Steve",
	}

	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	// Struct variables are not comparable if they contain fields that are not comparable
	/* image1 := image{
		data: map[int]int{
			0: 155,
		}}
	image2 := image{
		data: map[int]int{
			0: 155,
		}}
	if image1 == image2 {
		fmt.Println("image1 and image2 are equal")
	} */
}
