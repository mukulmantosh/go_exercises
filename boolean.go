package main

import "log"

func main() {

	var isExist bool = true

	if isExist == true {
		log.Println("Yes it's true")
	} else {
		log.Println("No it's not true")
	}

	num := 100

	if num > 80 && isExist == true {
		log.Println("YES !!!!!!")
	} else {
		log.Println("NO !!!!!!")
	}
}
