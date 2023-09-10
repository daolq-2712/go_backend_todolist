package route

import (
	"go_backend_todolist/api/controller"
	"go_backend_todolist/repository"
	"go_backend_todolist/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTodoItemRoute(db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {

	repository := repository.NewTodoItemRepository(db)
	controller := controller.TodoItemController{TodoItemUsecase: usecase.NewTodoItemUsecase(
		repository,
		timeout,
	)}

	items := group.Group("/items")
	{
		items.POST("", controller.Create)
		items.GET("", controller.Fetch)
		items.GET("/:id")
		items.PATCH("/:id")
		items.DELETE("/:id")
	}
}
