package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secret struct {
	person
	license bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says,  "Good morning, James."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		fname: "miss",
		lname: "hey",
	}

	p1.speak()

	sal := secret{
		person: person{
			fname: "james",
			lname: "bond",
		},
		license: true,
	}

	sal.speak()

	saySomething(sal)
}
