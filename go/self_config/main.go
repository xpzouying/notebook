package main

import "fmt"

// Person needs to be config
type Person struct {
	name string
	age  int
}

// NewPerson1 is the common way to make a new person
func NewPerson1(name string, age int) *Person {
	return &Person{name, age}
}

type personCfg struct {
	name string
	age  int
}

// NewPerson2 use config struct to make a new person
func NewPerson2(cfg personCfg) *Person {
	return &Person{cfg.name, cfg.age}
}

// option define the interface to config *Person
type option func(*Person)

// NewPerson is the new way learning from rob pike
func NewPerson(opts ...option) *Person {

	p := new(Person)

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func name(name string) option {
	return func(p *Person) {
		p.name = name
	}
}

func age(age int) option {
	return func(p *Person) {
		p.age = age
	}
}

func main() {
	p1 := NewPerson(name("zouying"), age(32))
	p2 := NewPerson(name("adam"))

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
}
