package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:firstname`
	LastName  string `json:lastname`
	Email     string `json:email`
}

func main() {
	data := `	
[
		{
		  "firstname": "John",
		  "lastname": "Beatles",
		  "email": "johnbeatles@gmail.com"
		},
		{
		  "firstname": "Rocky",
		  "lastname": "Balboa",
		  "email": "rockybalboa@gmail.com"		
		}
]
`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(data), &unmarshalled)
	if err != nil {
		log.Println("Error Unmarshelling JSON ", err)
	}

	// write JSON from struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Mukul"
	m1.LastName = "Mantosh"
	m1.Email = "abc@gmail.com"

	var m2 Person
	m2.FirstName = "Mukul"
	m2.LastName = "Mantosh"
	m2.Email = "abc@gmail.com"

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	log.Print(unmarshalled)

	log.Println(mySlice)

	newJson, error := json.MarshalIndent(mySlice, "", "  ")

	if error != nil {
		log.Println("Error Marshelling !!!")
	}

	log.Println(string(newJson))
}
