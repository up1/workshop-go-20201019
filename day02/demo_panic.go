package main

import (
	"fmt"
	"reflect"
)

func handle() {
	if err := recover(); err != nil {
		t1 := fmt.Sprintf("%T", err)
		t2 := fmt.Sprintf("%s", reflect.TypeOf(err))
		fmt.Println(t1, t2)
	}

}

func main() {
	defer handle()
	// panic("Hey")
	var a []int
	a[1] = 5 // Panic => index out of range
	fmt.Println("Finish")
}
