package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"go_backend_todolist/common"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatuses = [3]string{"Doing", "Done", "Deleted"}

func (item *ItemStatus) string() string {
	return allItemStatuses[*item]
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatuses {
		if allItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("invalid status string")
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	v, err := parseStr2ItemStatus(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*item = v

	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		defaultStatus := ItemStatus(0)
		return defaultStatus.string(), nil
	}

	return item.string(), nil
}

func (item *ItemStatus) MarshallJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.string())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {

	str := strings.ReplaceAll(string(data), "\"", "")

	v, err := parseStr2ItemStatus(str)

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", str)
	}

	*item = v

	return nil
}

type TodoItem struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      *ItemStatus `json:"status"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAd   *time.Time  `json:"updated_ad,omitempty"`
}

type TodoItemCreation struct {
	Id          int         `json:"id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string {
	return "todo_items"
}

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

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", createItem(db))
			items.GET("")
			items.GET("/:id")
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}

	r.Run(":3000")
}

func createItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data TodoItemCreation

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
