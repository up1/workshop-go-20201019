package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Array
	a := [3]string{"n1", "n2", "n3"}
	// Slice
	s := a[:2]
	s[0] = "somkiat"
	fmt.Println(a, s)
	fmt.Printf("%T, %T\n", a, s)
	fmt.Printf("%s, %s", reflect.ValueOf(a).Kind(), reflect.ValueOf(s).Kind())
}
