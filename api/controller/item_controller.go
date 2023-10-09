package controller

import (
	"go_backend_todolist/common"
	"go_backend_todolist/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoItemController struct {
	TodoItemUsecase domain.TodoItemUseCase
}

func (controller *TodoItemController) Create(ctx *gin.Context) {
	var data domain.TodoItemCreation

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		return
	}

	if err := controller.TodoItemUsecase.Create(ctx, &data); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrDB(err))
		return
	}

	ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
}

func (controller *TodoItemController) Fetch(ctx *gin.Context) {

	var paging common.Paging

	paging.Process()

	if err := ctx.ShouldBind(&paging); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		return
	}

	result, paging, err := controller.TodoItemUsecase.Fetch(ctx, paging)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrDB(err))
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
}
