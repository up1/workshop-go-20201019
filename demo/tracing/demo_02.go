package main

import (
	"log"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	_ = trace.Start(os.Stdout)
	defer trace.Stop()

	table := make(chan int)
	go player(table, "ping")
	go player(table, "pong")

	table <- 0
	time.Sleep(time.Second)
	ball := <-table
	close(table)
	log.Printf("played %d turns", ball)

	// runtime.GC()
	// _ = pprof.WriteHeapProfile(os.Stdout)
}

func player(table chan int, name string) {
	for ball := range table {
		log.Printf("%d\t%s", ball, name)
		table <- ball + 1
	}
}
