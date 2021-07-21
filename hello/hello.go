package main

import "fmt"

var z int = 5
var rawstring string = `You
Stink
of 
Shit`

func main() {
	fmt.Println("z was ", z)
	z = 99
	fmt.Println("z now ", z)
	fmt.Printf("hello, world\n")
	fmt.Println("Poo Skids")
	for i := 0; i < 5; i++ {
		//poo()
	}

	n, _ := fmt.Println("balls", 42, true, false)
	fmt.Println(n)

	x := 42
	y := 100 + x
	fmt.Println("x=", x, " y=", y, " sum of both=", x+y)

	fmt.Println(rawstring)

	const s = `bum
	holes`
	fmt.Println(s)

	fmt.Printf("X= %x", x)
}

// Cheese
func cheese() {
	fmt.Println("Cheese ")
}

// Burger
func burger() {
	fmt.Println("Burger")
}
