package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name  string
	Age   int
	Alive bool
}

// func (p *Person) changeLifeStatus() {
// 	p.Alive = false
// }

// func (s *Event) addField(key, value string) {
// 	v, _ := s.Event.(map[string]interface{})
// 	v[key] = value
// 	s.data = v
// }

func main() {

	p1 := Person{"Ben", 50, true}
	p2 := Person{"Bob", 52, true}

	fmt.Println(p1)
	fmt.Println(p2)

	var now time.Time = time.Now()
	now1 := time.Now()
	//now = time.Now()
	fmt.Println(now)
	fmt.Println(now1)

}
