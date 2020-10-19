package main

import (
	"encoding/json"
	"fmt"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p Post) String() string {
	return fmt.Sprintf("ID=%v, User ID=%v", p.ID, p.UserID)
}

func main() {
	p1 := Post{ID: 100}
	// b, err := json.Marshal(p1)
	b, err := json.MarshalIndent(p1, "", "    ")
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Println(string(b))
	}

	var out Post
	err = json.Unmarshal(b, &out)
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Println(out)
	}
}
