package main

import (
	"fmt"
	"slices"
)

func main() {
	var arregloString []string

	fmt.Println(arregloString)

	arregloString = make([]string, 6)
	arregloString[0] = "Hola"
	arregloString[1] = "Mundo"
	arregloString[5] = "Golang"
	fmt.Println(arregloString)

	fmt.Println(len(arregloString))

	arregloString = append(arregloString, "!")
	fmt.Println(arregloString)
	fmt.Println(len(arregloString))

	secondArray := []string{"test", "append", "campus"}

	if slices.Equal(arregloString, secondArray) {
		fmt.Println("Los arreglos son iguales")
	} else {
		fmt.Println("Los arreglos no son iguales")
	}
}
