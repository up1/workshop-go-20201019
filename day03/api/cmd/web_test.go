package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHandleAdd(b *testing.B) {
	b.ReportAllocs()
	r := request(b, "/?operand1=70&operand2=30")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		handleAdd(rw, r)
	}
}

func request(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}
