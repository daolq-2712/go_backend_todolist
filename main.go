package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAd   *time.Time `json:"updated_ad,omitempty"`
}

func main() {
	fmt.Println("Hello World!!")

	item := TodoItem{
		Id:          8,
		Title:       "This is title 8",
		Description: "This is description 8",
	}

	// Endcode
	jsonData, err := json.Marshal(item)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	// Decode
	jsonStr := "{\"id\":8,\"title\":\"This is title 8\",\"description\":\"This is description 8\",\"status\":\"\",\"created_at\":null,\"updated_ad\":null}"

	var item2 TodoItem
	if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(item2)
}
