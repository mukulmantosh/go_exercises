package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
}

func main() {
	user := User{
		FirstName: "Mukul",
		LastName:  "Mantosh",
		Age:       10,
		BirthDate: time.Now(),
	}

	fmt.Println(user.BirthDate)
}
