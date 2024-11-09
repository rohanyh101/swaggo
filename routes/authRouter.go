package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roh4nyh/swaggo/controllers"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("auth/signup", controllers.UserSignUp())
	incomingRoutes.POST("auth/login", controllers.UserLogIn())
}
