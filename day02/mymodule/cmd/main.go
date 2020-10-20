package main

import (
	"demo"
	"demo/json"
	"os"
	"runtime/trace"
)

func main() {
	_ = trace.Start(os.Stdout)
	defer trace.Stop()
	demo.SayHi()
	json.StructToJson01()
	json.StructToJson02()
}
