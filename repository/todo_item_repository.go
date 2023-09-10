package repository

import (
	"context"
	"go_backend_todolist/common"
	"go_backend_todolist/domain"

	"gorm.io/gorm"
)

type todoItemRepository struct {
	database *gorm.DB
}

func NewTodoItemRepository(db *gorm.DB) domain.TodoItemRepository {
	return &todoItemRepository{db}
}

func (repository *todoItemRepository) Create(ctx context.Context, todoItem *domain.TodoItemCreation) error {
	if err := repository.database.Create(&todoItem).Error; err != nil {
		return err
	}
	return nil
}

func (repository *todoItemRepository) Fetch(ctx context.Context, paging common.Paging) ([]domain.TodoItem, common.Paging, error) {
	db := repository.database.Where("status <> ?", "Deleted")

	if err := db.Table(domain.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, paging, err
	}

	var todoItems []domain.TodoItem

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&todoItems).Error; err != nil {
		return nil, paging, err
	}
	return todoItems, paging, nil
}
