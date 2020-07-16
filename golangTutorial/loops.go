package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		} else if i > 5 {
			break
		}
		fmt.Printf("%v ", i)
	}
	fmt.Println("\nloop finished.")

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Println()
	}

outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}

	}

	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present
		fmt.Printf("%d ", i)
		i += 2
	}

	for i, j := 10, 1; i < 20 && j <= 10; i, j = i+1, j+1 {
		fmt.Printf("\n%v * %v = %v", i, j, i*j)
	}
}
