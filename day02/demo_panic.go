package main

import (
	"fmt"
	"runtime/debug"
)

func handle() {
	if err := recover(); err != nil {
		if err == "Hey" {
			fmt.Println("Hey Hey")
		}
		debug.PrintStack()
	}
	fmt.Println("handle...")
}

func main() {
	defer handle()
	// panic("Hey")
	var a []int
	a[1] = 5
	fmt.Println("Finish")
}
