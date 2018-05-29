package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meritlabs/achievement-engine/api/controllers"
	"github.com/meritlabs/achievement-engine/api/middleware"
)

func main() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.JSONErrorReporter())

	apiGroup := router.Group("/achivement-engine/api/v1")
	{
		goals := apiGroup.Group("/goals")
		{
			goals.GET("/", controllers.ListGoals)
		}
	}

	router.Run()
}
