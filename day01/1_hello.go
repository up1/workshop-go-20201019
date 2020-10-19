package main

import "fmt"

func init() {
	fmt.Println("Init")
}

func main() {
	fmt.Println("Hello")
	call()
}

func say() {
	fmt.Println("Say")
}

//go run 1_hello.go
