package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("This is an int", x)
	case string:
		fmt.Println("This is a string", x)
	case contact:
		fmt.Println("This is contact x", x)
		fmt.Println("\tGreeting : ", x.(contact).greeting)
		fmt.Println("\tName : ", x.(contact).name)
		var y = x.(contact)
		fmt.Println("This is contact y's greeting ", y.greeting)
		//y := returnContact(x.(contact))

	default:
		fmt.Println("Unknown")
	}

}

func main() {

	switchOnType(1)
	switchOnType("1")
	x := contact{
		greeting: "Hello",
		name:     "Ben",
	}
	switchOnType(x)
	//x.checkIfContact()

}
