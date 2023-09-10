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
