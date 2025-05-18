package main

import "fmt"

func ValoresMultiples(nombre1, nombre2, apellido string) (map[string]string, map[string]string) {
	map1 := make(map[string]string)
	map2 := make(map[string]string)

	map1[nombre1] = apellido
	map2[nombre2] = apellido

	return map1, map2
}

func main() {
	nombre1 := "ivan"
	nombre2 := "ortega"
	apellido := "garzon"

	mapaFirtsName, mapaSecondName := ValoresMultiples(nombre1, nombre2, apellido)

	fmt.Println("mapara con el primer nombre", mapaFirtsName)
	fmt.Println("mapara con el segundo nombre", mapaSecondName)

	_, nombre := ValoresMultiples(nombre1, nombre2, apellido)
	fmt.Println(nombre)

}
