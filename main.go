package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println("Hello, World! B")
}
