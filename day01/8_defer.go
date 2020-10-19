package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i, ".Called defer")
	}
	fmt.Println("End main")
}
