package main

import (
	"fmt"
	"learnstructs/computer"
)

func main() {
	spec := computer.Spec{
		Maker: "apple",
		Price: 50000,
	}
	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
}
