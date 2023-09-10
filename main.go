package main

import (
	"fmt"
	"go_backend_todolist/api/route"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World!!")

	// Checking that an environment variable is present or not.
	mysqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")

	if !ok {
		log.Fatalln("Missing MySQL connection string.")
	}

	dsn := mysqlConnStr
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	r := gin.Default()

	// CRUD: Create, Read, Update, Delete
	// POST /v1/items (create a new item)
	// GET /v1/items (list items) /v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// (PUT | PATH) v1/items/:id (update item by id)
	// DELETE /v1/items/:id (delete item by id)

	timeout := time.Duration(3) * time.Second

	v1 := r.Group("/v1")

	route.Settup(timeout, db, v1)

	r.Run(":3000")
}
