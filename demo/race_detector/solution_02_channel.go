package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNumber())
}
func getNumber() int {
	var i int
	done := make(chan struct{})

	go func() {
		i = 5
		done <- struct{}{}
	}()

	<-done
	return i
}
