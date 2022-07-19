package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Dog",
		Breed: "German Shephard",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "Grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)

}

func PrintInfo(a Animal) {
	fmt.Println("The animal says ", a.Says(), " it has number of legs ", a.NumberOfLegs())
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "HoHoHooo"
}

func (d *Gorilla) NumberOfLegs() int {
	return 4
}
