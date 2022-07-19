package main

import "log"

func main() {
	sayHi("Hello World with GoLang")
}

func sayHi(s string) {
	log.Println(s)
}
