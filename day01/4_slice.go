package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:2]
	b[0] = 5555
	fmt.Printf("A=%t, B=%p", a, b)
}
