package main

import (
	"fmt"
	"runtime/debug"
	_ "time"
)

// If recover is called outside the deferred function, it will not stop a panicking sequence.
/*func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

// Recover works only when it is called from the same goroutine.
func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
} */

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		// print the stack trace after recovery
		debug.PrintStack()
	}
}

func a() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	// all deferred actions will be done before printing the panic message
	/* defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	a()
	fmt.Println("normally returned from main") */

	a()
	fmt.Println("normally returned from main")
}
