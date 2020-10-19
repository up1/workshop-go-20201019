package main

import "testing"

func BenchmarkSolution01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		structToJson01()
	}
}

func BenchmarkSolution02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		structToJson02()
	}
}
