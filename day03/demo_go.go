package main

import (
	"fmt"
	"sync"
)

func hello(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello ", name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello("1", &wg)
	wg.Add(1)
	go hello("2", &wg)
	wg.Add(1)
	go hello("3", &wg)

	wg.Wait()
	fmt.Println("Done")
}
