package demo_test

import (
	. "demo"
	"testing"
)

func TestHello(t *testing.T) {
	result := SayHi()
	if result != "Hello sayHi" {
		t.Errorf("Fail with =>%s", result)
	}
}

func BenchmarkDemo01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHi()
	}
}

func BenchmarkDemo02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHi2()
	}
}
