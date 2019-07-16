package jedi2

import (
	"fmt"
)

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("hello", p.first)
}

func saySomething(h human) {
	h.speak()
}

// Ex2 tbd
func Ex2() {
	fmt.Println("\nEx2")

	p := person{
		first: "Tyler",
	}

	p.speak()

	saySomething(&p)
}
