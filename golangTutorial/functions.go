package main

import "fmt"

func main() {
	fmt.Println(multiply(3, 4))
	area, perimeter := rectProps(12.34, 56.78)
	fmt.Println(area, perimeter)

	_, perimeter = rectProps(12, 34)
	fmt.Println(area)
}

func multiply(a, b int) int {
	return a * b
}

func rectProps(width, length float32) (area, perimeter float32) {
	// if the return value is not set, return the default value of the type
	area = width * length
	perimeter = (width + length) * 2
	return
}
