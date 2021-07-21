package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Always CAPITALISE the fields you wish to export
// below example will no export id
type Engineer struct {
	id       int
	Forename string `json:"fn"`
	Surname  string `json:"sn"`
	Age      int    `json:"-"` // Ignore this
}

func main() {

	group := Engineer{
		id:       1,
		Forename: "Ben",
		Surname:  "Wood",
		Age:      50,
	}
	fmt.Println(group)

	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	os.Stdout.Write(b) // Returns {"fn":"Ben","sn":"Wood"}

}
