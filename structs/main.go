package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main () {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	var marie person
	marie.firstName = "Marie"
	marie.lastName = "Curie"

	fmt.Println(marie)
	fmt.Printf("%+v", alex)
}