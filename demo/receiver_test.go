package main

import "testing"

type SampleStruct struct {
	StringProperty string
}

func (a *SampleStruct) changeItP() {
	// Do nothing
}

func (a SampleStruct) changeItV() {
	// Do nothing
}

func BenchmarkChangePointerReceiver(b *testing.B) {
	s := SampleStruct{}
	for i := 0; i < b.N; i++ {
		s.changeItP()
	}
}

func BenchmarkChangeItValueReceiver(b *testing.B) {
	s := SampleStruct{}
	for i := 0; i < b.N; i++ {
		s.changeItV()
	}
}
