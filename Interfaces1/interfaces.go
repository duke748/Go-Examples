package main

import "fmt"

type Superheroes interface {
	SpecialPowers() string
	Weakness() string
}

type Batman struct {
	//alias, weakness string
}

func (b Batman) SpecialPowers() string {
	fmt.Println("batman func was passed", b)
	return "Money"
}

func (b Batman) Weakness() string {
	fmt.Println("Weakness func was passed", b)
	return "Stock Market Crash"
}

type Superman struct {
	//alias, weakness string
}

func (s Superman) SpecialPowers() string {
	return "Flight, Superhuman Strength, Laser Eyes, Speed"
}

func (s Superman) Weakness() string {
	return "Krptonite"
}

func main() {
	//heroes := []Superheroes{Batman{"Bruce Wayne"}, Superman{}}
	heroes := []Superheroes{Batman{}, Superman{}}
	for i, hero := range heroes {
		fmt.Println("hero=", heroes[i].SpecialPowers(), "weakness=", hero.Weakness(), "special Powers=", hero.SpecialPowers())
	}
	//returnPowers(heroes)
	//bat := Batman{}
	//fmt.Println(bat.SpecialPowers())

}

func returnPowers(args ...interface{}) {
	fmt.Println("returnPowers called with")
	for _, item := range args {
		fmt.Printf("%v\n", item)
	}
}
