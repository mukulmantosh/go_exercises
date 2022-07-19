package main

import (
	"awesomeProject/helpers"
	"log"
)

func main() {

	var myVar helpers.SomeType
	myVar.TypeName = "SOME NAME"
	log.Println(myVar)
}
