package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(getNumber())
}

type SafeNumber struct {
	val int
	m   sync.Mutex
}

func (i *SafeNumber) Read() int {
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}

func (i *SafeNumber) Write(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

func getNumber() int {
	i := &SafeNumber{}
	go func() {
		i.Write(5)
	}()
	return i.Read()
}
