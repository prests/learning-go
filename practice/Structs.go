package main

import (
	"fmt"
)

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
}

func main() {
	type cat struct {
		claws bool
	}

	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)

	littleCat := cat{true}
	fmt.Println(littleCat)
}
