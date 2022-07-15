package main

import "log"

type myStruct struct {
	firstName string
}

func main() {
	var myVar myStruct
	myVar.firstName = "John"

	myVar2 := myStruct{
		firstName: "Mukul",
	}

	log.Println("myvar is set to ", myVar.firstName)
	log.Println("myvar2 is set to ", myVar2.firstName)

	log.Println("The result is =>  ", myVar.printFirstName())
}

func (m *myStruct) printFirstName() string {
	return m.firstName
}
