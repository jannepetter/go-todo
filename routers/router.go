package routers

import (
	"goServer/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	handlers.SetTodoRoutes(router.Group(""))
	return router
}
