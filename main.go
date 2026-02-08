package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(add(2, 4))

	fmt.Println(add(2, 5))
}
