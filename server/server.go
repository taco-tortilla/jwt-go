package server

import (
	"github.com/gin-gonic/gin"
	"github.com/taco-tortilla/jwt-go/controllers"
	"github.com/taco-tortilla/jwt-go/middleware"
)

func Init() {
	router := router()
	router.Run()
}

func router() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	return router
}
