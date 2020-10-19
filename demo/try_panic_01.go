package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read data from file
	b, err := ioutil.ReadFile("try_panic.go")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
