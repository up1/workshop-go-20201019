package main

import (
	"net/http"
)

func Response(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world."))
}

func main() {
	http.HandleFunc("/", Response)
	http.ListenAndServe(":8080", nil)
}
