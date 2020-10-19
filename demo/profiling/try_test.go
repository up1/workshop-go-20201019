package main

import "testing"

var name string
var ok bool

func BenchmarkIsGopher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name, ok = isGopher("demo@golang.org")
	}
}
