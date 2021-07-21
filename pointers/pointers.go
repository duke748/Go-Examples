package main

import "fmt"

func main() {
	a := 42
	fmt.Println("A Was :", a)

	// Get the address i.e 0x1040a122
	b := &a // b now is equal to that address in memory i.e. 0x1040a122
	fmt.Println("B = ", b)
	fmt.Printf("&a = %p", &a)

	// Increment the value at the address
	*b++
	fmt.Println("\nA Now :", a)
}
