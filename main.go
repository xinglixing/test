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

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Hello, World! B")

	fmt.Println("Hello, World! C")

	fmt.Println("Hello, World! D")

	fmt.Println(add(2, 3))

	fmt.Println(subtract(5, 2))

	fmt.Println(multiply(4, 5))

	fmt.Println(divide(10, 2))

	fmt.Println(power(2, 3))

	fmt.Println(factorial(5))
}
