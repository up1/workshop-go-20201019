package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	log.Printf("listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, golang"
	if name, ok := isGopher(r.URL.Path[1:]); ok {
		msg = "Hello, " + name
	}
	_, err := fmt.Fprintln(w, msg)
	if err != nil {
		log.Printf("could not print message: %v", err)
	}
}

func isGopher(email string) (string, bool) {
	re := regexp.MustCompile("^([[:alpha:]]+)@golang.org$")
	match := re.FindStringSubmatch(email)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}
