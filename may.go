package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Title   string `json:"title"`
}

func main() {

	item := Item{Title: "Car"}
	jitem, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(jitem))
}
