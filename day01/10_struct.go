package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	name string
	age  int
	c    struct{}
}

func (p *person) sing() {
	p.age = 100
	fmt.Println("In sing ", p)
}

func main() {
	p := person{name: "somkiat"}
	p.sing()
	fmt.Println("In main ", p)

	type first struct {
		f float64 // 8 bytes
		b bool    // 1 byte
		i int32   // 4 bytes
	}
	a := first{}
	fmt.Println(unsafe.Sizeof(a))
}
