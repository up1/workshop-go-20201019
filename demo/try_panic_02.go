package main

import (
	"fmt"
	"io/ioutil"
	"runtime/debug"
)

func panicHandler() {
	err := recover()
	if err == "some error" {
		fmt.Println("Try to recover from panic")
		debug.PrintStack()
	}
}

func main() {
	// Defer
	defer panicHandler()
	// Read data from file
	b, err := ioutil.ReadFile("try_panic.go")
	if err != nil {
		panic("some error")
	}

	fmt.Println(string(b))
}
