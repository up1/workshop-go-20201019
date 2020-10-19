package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleStructAdd(t *testing.T) {

	r := request(t, "/?operand1=70&operand2=30")

	rw := httptest.NewRecorder()

	handleAdd(rw, r)
	if rw.Code == 500 {
		t.Fatal("Internal server Error: " + rw.Body.String())
	}
	if rw.Body.String() != "<h1>Result is 100</h1>" {
		t.Fatal("Expected " + rw.Body.String())
	}

}
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
