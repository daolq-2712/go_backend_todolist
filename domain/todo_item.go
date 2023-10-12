package domain

import (
	"context"
	"go_backend_todolist/common"
	"time"
)

type TodoItem struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      *ItemStatus `json:"status"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAd   *time.Time  `json:"updated_ad,omitempty"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int         `json:"id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
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
