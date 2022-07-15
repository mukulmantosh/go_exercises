package main

import "log"

func main() {
	var data int = 12

	switch data {
	case 1:
		log.Println("Inside case 1")
	case 2:
		log.Println("Inside case 2")
	case 3:
		log.Println("Inside case 3")
	default:
		log.Println("Nothing Happens !!!!")
	}
}
