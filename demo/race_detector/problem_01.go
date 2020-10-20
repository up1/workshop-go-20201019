package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Seeding is important for choosing a random counter
	rand.Seed(time.Now().Unix())

	// 3 counters
	a := []int{0, 0, 0}

	// 2 tasks will be launched. The main goroutine will
	// wait for them to complete.
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		i := rand.Intn(3)
		a[i]++
		wg.Done()
	}()

	go func() {
		j := rand.Intn(3)
		a[j]++
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(a)
}
