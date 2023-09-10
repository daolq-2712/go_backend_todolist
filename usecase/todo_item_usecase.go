package usecase

import (
	"context"
	"go_backend_todolist/domain"
	"time"
)

type todoItemUseCase struct {
	todoItemRespository domain.TodoItemRepository
	contextTimeout      time.Duration
}

func NewTodoItemUsecase(repository domain.TodoItemRepository, timeout time.Duration) domain.TodoItemUsecase {
	return &todoItemUseCase{
		todoItemRespository: repository,
		contextTimeout:      timeout,
	}
}

func (usecase *todoItemUseCase) Create(ctx context.Context, todoItem *domain.TodoItemCreation) error {
	c, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()
	return usecase.todoItemRespository.Create(c, todoItem)
}

func (usecase *todoItemUseCase) Fetch(ctx context.Context) ([]domain.TodoItem, error) {
	c, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	todoItems, err := usecase.todoItemRespository.Fetch(c)
	if err != nil {
		return nil, err
	}
	return todoItems, nil
}
