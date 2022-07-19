package main

import (
	"fmt"
	"log"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	animals := []string{"dog", "fish", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	data := make(map[string]string)
	data["dog"] = "German Shephard"
	data["cat"] = "Meow"
	log.Println(data)

	for key, value := range data {
		log.Println(key, value)
	}

	var firstLine = "This is a sample text word document"

	for i, j := range firstLine {
		log.Println(i, j) // This is the ASCII Information.
	}

	// range over maps, list, custom objects, string

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@gmail.com", 30})
	users = append(users, User{"Mary", "Smith", "john@gmail.com", 45})

	for i, j := range users {
		log.Println(i, j)
	}

}
