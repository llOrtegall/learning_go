package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("HOME environment variable not found")
	} else {
		fmt.Printf("HOME environment variables: %s\n", envVar)
	}

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
}
