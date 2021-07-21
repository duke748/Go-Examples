//
package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}
	
	fmt.Println(m["James"])

	v, ok := m["James"]

	m["Q"] = 108

	fmt.Println(v, ok)
	if v, ok := m["Q"]; ok {
		fmt.Println("This is Ok ", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
