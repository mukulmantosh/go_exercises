package main

import "log"

func main() {

	var myString string
	myString = "Hello Mukul !"

	log.Println("Before Change -> ", myString)

	changePointer(&myString)

	log.Println("After Change -> ", myString)
}

func changePointer(s *string) {
	newData := "Hello Mukul Mantosh !"
	*s = newData
}
