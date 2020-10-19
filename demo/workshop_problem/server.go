package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type add struct {
	Sum int
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	var html bytes.Buffer
	operand1, operand2 := r.FormValue("operand1"), r.FormValue("operand2")
	one, err, w := validate(operand1, w)
	two, err, w := validate(operand2, w)
	m := struct{ a, b int }{one, two}
	structSum := add{Sum: m.a + m.b}

	t, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	err = t.Execute(&html, structSum)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html.String()))
}

func validate(operand1 string, w http.ResponseWriter) (int, error, http.ResponseWriter) {
	one, err := strconv.Atoi(operand1)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	return one, err, w
}
func main() {
	http.HandleFunc("/add", handleAdd)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
