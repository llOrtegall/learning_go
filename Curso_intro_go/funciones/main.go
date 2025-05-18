package main

import "fmt"

func sum(numero1, numero2 int) int {
	resultado := numero1 + numero2
	return resultado
}

func sum2(numero1, numero2, numero3 int) int {
	resultado := numero1 + numero2 + numero3
	return resultado
}

func main() {
	var numero1, numero2, numero3 int

	fmt.Println("Ingresa el primer número: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingresa el segundo número: ")
	fmt.Scanln(&numero2)
	fmt.Println("Ingresa el tercer número: ")
	fmt.Scanln(&numero3)

	resultado := sum(numero1, numero2)
	fmt.Println("la suma de dos número es: ", resultado)

	resultado2 := sum2(numero1, numero2, numero3)
	fmt.Println("la suma de tres número es: ", resultado2)

}
