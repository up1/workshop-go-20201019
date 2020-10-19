package main

import "fmt"

func main() {
	// Array
	a := [3]string{"n1", "n2", "n3"}
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], ",")
	}

	for _, v := range a {
		fmt.Println(v, ",")
	}
}
