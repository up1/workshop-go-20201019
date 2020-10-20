package main

import "fmt"

type animal interface {
	eat()
}
type human int
type cat int

// human implements animal
func (h human) eat() { fmt.Println("Human eat") }
func (c cat) eat()   { fmt.Println("Cat eat") }

func eat(a animal) {
	a.eat()
}

func eat2(a interface{}) {
	a.(animal).eat() // Cast type
}

func main() {
	var h human
	var c cat
	eat(h)
	eat(c)
}
