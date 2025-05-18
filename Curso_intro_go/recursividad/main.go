package main

import "fmt"

// factorial example 7! = 7*6*5*4*3*2*1
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(7))

	var fibonacci func(numeroBase int) int

	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}

		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Println(fibonacci(7))
}
