package main

import (
	"os"
	"runtime/trace"
)

func main() {
	_ = trace.Start(os.Stdout)
	defer trace.Stop()
}
