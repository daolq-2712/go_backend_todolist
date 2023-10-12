package usecase

import (
	"context"
	"go_backend_todolist/common"
	"go_backend_todolist/domain"
	"time"
)

type todoItemUseCase struct {
	todoItemRepository domain.TodoItemRepository
	contextTimeout     time.Duration
}

func NewTodoItemUseCase(repository domain.TodoItemRepository, timeout time.Duration) domain.TodoItemUseCase {
	return &todoItemUseCase{
		todoItemRepository: repository,
		contextTimeout:     timeout,
	}
}

func (useCase *todoItemUseCase) Create(ctx context.Context, todoItem *domain.TodoItemCreation) error {
	c, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()
	return useCase.todoItemRepository.Create(c, todoItem)
}

func (useCase *todoItemUseCase) Fetch(ctx context.Context, paging common.Paging) ([]domain.TodoItem, common.Paging, error) {
	c, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()

	todoItems, paging, err := useCase.todoItemRepository.Fetch(c, paging)
	if err != nil {
		return nil, paging, err
	}
	return todoItems, paging, nil
}
