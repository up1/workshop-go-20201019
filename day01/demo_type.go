package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [5]int{1, 2, 3}
	s := []int{4, 5, 6}
	fmt.Printf("%T, %T\n", arr, s)
	fmt.Printf("%s, %s\n", reflect.ValueOf(arr).Kind(), reflect.ValueOf(s).Kind())
}
