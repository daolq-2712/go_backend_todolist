package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	r.Run(":3000")
}
