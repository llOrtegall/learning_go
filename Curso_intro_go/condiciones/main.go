package main

import "fmt"

func main() {
	nombres := "jhon vans"
	edad := 27

	if nombres == "jhon vans" {
		fmt.Println("Hola jhon vans")
	} else {
		fmt.Println("Hola desconocido")
	}

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

}
