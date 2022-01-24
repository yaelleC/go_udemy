package main

import "fmt"

type contactInfo struct {
	email string
	zipcode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main () {
	alex := person{
		firstName: "Alex", 
		lastName: "Anderson",
		contactInfo: contactInfo{
			email: "a@alex.com",
			zipcode: 12345,
		},
	}
	var marie person
	marie.firstName = "Marie"
	marie.lastName = "Curie"

	fmt.Println(marie)

	alex.updateName("Jim")
	alex.print()

	fmt.Println(*&alex)
	name := "bill"
 
	namePointer := &name
	
	fmt.Println(&namePointer)
	printPointer(namePointer)
   }
	
   func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
   }

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerPerson * person) updateName(newFirstName string) {
	(*pointerPerson).firstName = newFirstName
}