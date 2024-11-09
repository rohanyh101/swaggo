package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/roh4nyh/swaggo/controllers"
	"github.com/roh4nyh/swaggo/middleware"
)

func AdminRoutes(incomingRoutes *gin.RouterGroup) {
	// admin routes
	incomingRoutes.GET("/users", middleware.Authenticate(), middleware.AuthenticateAdmin(), controller.GetUsers())
}
