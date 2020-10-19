package main

import "fmt"

var c int = 5

func main() {
	name := "somkiat"
	names := `
		line 1
		line 2
		line 3
	`
	fmt.Println(string(name[0]))
	fmt.Println(names[0])
}

func var01() {
	a, b := 1, 2
	b, a = a, b
	i, _ := 10, 20
	fmt.Println(a, b)
	fmt.Println(i)
}

//go run 1_hello.go
