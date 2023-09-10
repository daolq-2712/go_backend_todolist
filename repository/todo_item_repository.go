package repository

import (
	"context"
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
