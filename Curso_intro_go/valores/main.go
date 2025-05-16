package main

import "fmt"

func main() {
	var numero int = 10
	numero2 := 20

	fmt.Println(numero, numero2)

	var DoubleNumber = 20.6
	fmt.Println(DoubleNumber)

	var result = numero + int(DoubleNumber)
	fmt.Println(result)

	var names = "jhon vans"
	lastName := "doe"
	fullName := names + " " + lastName
	fmt.Println(fullName)
}
