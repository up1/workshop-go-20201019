package json_test

import (
	"demo/json"
	"testing"
)

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.StructToJson01()
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.StructToJson02()
	}
}
