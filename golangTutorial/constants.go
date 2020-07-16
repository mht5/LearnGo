package main

import "fmt"

func main() {
	// can not re-assign a value to constants
	const a = 10
	// a = 20
	fmt.Println(a)

	// constant with no type
	// untyped constants have a default type associated with them and they supply it if and only if a line of code demands it
	const b = "Hello, "
	// constant with type
	const c string = "World!"
	fmt.Println(b + c)

	var name = "Sam"
	type myString string
	var newName myString = "Jack"
	fmt.Println(name, newName)
	// mixing types during assignment is not allowed
	// newName = name
	newName = myString(name)
	fmt.Println(name, newName)

	// must declare x as a const, because const has no type, so it could be assigned to any compatible type
	// var x = 5
	const x = 5
	var intVar int = x
	var int32Var int32 = x
	var float64Var float64 = x
	var complex64Var complex64 = x
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// Numeric constants are free to be mixed and matched in expressions
	// and a type is needed only when they are assigned to variables or used in any place in code which demands a type
	var i = 5.9 / 8
	fmt.Printf("a's type %T value %v", i, i)
}
