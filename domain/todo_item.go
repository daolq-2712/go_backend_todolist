package domain

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"go_backend_todolist/common"
	"strings"
	"time"
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

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
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

func (TodoItem) TableName() string {
	return "todo_items"
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemRepository interface {
	Create(ctx context.Context, todoItem *TodoItemCreation) error

	Fetch(ctx context.Context, paging common.Paging) ([]TodoItem, common.Paging, error)
}

type TodoItemUseCase interface {
	Create(ctx context.Context, todoItem *TodoItemCreation) error

	Fetch(ctx context.Context, paging common.Paging) ([]TodoItem, common.Paging, error)
}
