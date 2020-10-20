package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Post struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func StructToJson01() {
	p1 := Post{ID: 100}
	_, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error ", err)
	}
}

func StructToJson02() {
	p1 := Post{ID: 100}
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(&p1)
}
