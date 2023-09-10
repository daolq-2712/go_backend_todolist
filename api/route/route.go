package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Settup(timeout time.Duration, database *gorm.DB, ginRouter *gin.RouterGroup) {
	publicRouter := ginRouter
	// All public apis
	NewTodoItemRoute(database, timeout, publicRouter)
}
