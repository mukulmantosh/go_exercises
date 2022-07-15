package main

import (
	"log"
	"sort"
)

type Usr struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)
	newMap := make(map[string]int)
	userMap := make(map[string]Usr)
	unknownMap := make(map[string]interface{}) // map which accepts anything

	me := Usr{
		FirstName: "Mukul",
		LastName:  "Mantosh",
	}

	myMap["dog"] = "jerry"
	newMap["phone_number"] = 123456
	userMap["me"] = me

	log.Println(myMap["dog"])
	log.Println(newMap["phone_number"])
	log.Println(userMap["me"].FirstName)

	var myNewVar float32 = 11.2
	log.Println(myNewVar)

	unknownMap["firstName"] = "Mukul"
	unknownMap["lastName"] = "Mantosh"
	unknownMap["phoneNumber"] = 123456

	log.Println(unknownMap)

	/* slices */
	var mySlice []string

	mySlice = append(mySlice, "Mukul")
	mySlice = append(mySlice, "Mantosh")

	sort.Strings(mySlice)
	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	log.Println(numbers)

	log.Println(numbers[0:2])

}
