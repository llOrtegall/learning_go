package main

import "fmt"

func main() {
	var arreglo [5]int

	fmt.Println("arreglo completo", arreglo)
	arreglo[2] = 100
	fmt.Println("arreglo completo", arreglo)
	fmt.Println("arreglo", len(arreglo))

	arreglo2 := [5]int{6, 1, 8, 9, 0}
	fmt.Println("arreglo2", arreglo2)

	arreglo3 := [...]int{5, 3, 6, 8, 10, 23}
	fmt.Println("arreglo3", arreglo3)
	fmt.Println("arreglo3", len(arreglo3))

	arreglo4 := [...]int{3, 1, 5, 2}
	fmt.Println("arreglo4", arreglo4)
}
