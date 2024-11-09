package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/roh4nyh/swaggo/controllers"
	"github.com/roh4nyh/swaggo/middleware"
)

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	userProfileRoutes := incomingRoutes.Group("/profile")
	userProfileRoutes.Use(middleware.Authenticate(), middleware.AuthenticateUser())

	// user crud operations
	userProfileRoutes.GET("/", controller.GetUser())
	userProfileRoutes.PUT("/", controller.UpdateUser())
	userProfileRoutes.DELETE("/", controller.DeleteUser())
}
