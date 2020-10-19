package main

import "fmt"

var errorInput = fmt.Errorf("Input invalid")

func main() {
	result, err := add(10, 2)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(result)
}

func add(a int, b int) (int, error) {
	if a > b {
		return 0, fmt.Errorf("Input invalid")
	}
	return a + b, nil
}
