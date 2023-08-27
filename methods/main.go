package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) data() person {
	fmt.Println(p.name)
	fmt.Println(p.age)
	return p
}

func (p *person) mutate(name string, age int) *person {
	p.name = name
	p.age = age
	return p
}

func main() {
	a := person{age: 11, name: "yo nigga"}
	fmt.Println(a.data())
	fmt.Println(a.mutate("Yo nigger", 18))
	fmt.Println(a)
}
