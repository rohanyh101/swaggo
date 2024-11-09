package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/roh4nyh/swaggo/docs"
	"github.com/roh4nyh/swaggo/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	PORT := "8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/health", HealthCheck)

	routes.AuthRoutes(apiV1)
	routes.UserRoutes(apiV1)
	routes.AdminRoutes(apiV1)

	log.Printf("server running on port %s", PORT)
	r.Run(fmt.Sprintf(":%s", PORT))
}

// Health check route
// @Summary     Health check
// @Description Health check endpoint to see if the server is running
// @Tags        Health
// @Accept      json
// @Produce     json
// @Success     200 {object} map[string]string
// @Router      /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "server is up and running..."})
}
