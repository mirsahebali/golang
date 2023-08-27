package main

import "fmt"

type Direction string

const (
	EAST  Direction = "EAST"
	WEST  Direction = "WEST"
	NORTH Direction = "NORTH"
	SOUTH Direction = "SOUTH"
)

type person struct {
	name      string
	age       int
	role      string
	direction Direction
}

func createPerson(name string) *person {
	p := person{name: name, role: "Admin"}
	p.age = 99
	return &p
}

func main() {
	fmt.Println(person{name: "Bob", age: 88, role: "Admin"})

	p := person{name: "Help", role: "Admin"}
	p.age = 99
	fmt.Println(&person{name: "yellow", age: 33})

	fmt.Println(&person{name: "hellwo", age: 33})
	fmt.Println(createPerson("Yello"))
	sp := &p

	sp.direction = NORTH
	fmt.Println(sp.direction)

	cat := struct {
		sound  string
		isGood bool
	}{
		sound:  "Meow",
		isGood: true,
	}
	fmt.Println(cat.isGood)
}
