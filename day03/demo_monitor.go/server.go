package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Response(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world."))
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", Response)
	http.ListenAndServe(":8080", nil)
}
