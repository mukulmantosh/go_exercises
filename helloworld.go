package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var whatToSay string = "Mukul"
	fmt.Println(whatToSay)

	var i int64 = 2
	fmt.Println(i)

	sayMe := saySomething()
	fmt.Println(sayMe)
}

func saySomething() string {
	return "something"
}
