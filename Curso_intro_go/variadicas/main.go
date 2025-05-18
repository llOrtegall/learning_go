package main

import "fmt"

func sum(numeros ...int) {
	fmt.Println("números: ", numeros)

	total := 0

	for _, num := range numeros {
		total += num
	}

	fmt.Println("La suma es: ", total)
}

func main() {
	sum(1, 2)
	sum(1, 3, 5)
	sum(54, 67, 31, 43, 78)

	numeros := []int{78, 45, 32, 12, 66, 89}
	sum(numeros...)
}
