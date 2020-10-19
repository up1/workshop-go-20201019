package main

import "fmt"

func main() {
	var number = make(map[string]int)
	number["one"] = 1
	number["two"] = 2
	for k, v := range number {
		fmt.Println(k, v)
	}

	if v, found := number["three"]; !found {
		fmt.Println("Not found")
	} else {
		fmt.Println(v)
	}

}
