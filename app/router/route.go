package router

import (
	"library-be/app/middleware"
	"library-be/bootstrap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(init *bootstrap.Initialization) *gin.Engine {
	router := gin.Default()

	// cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// middleware
	router.Use(gin.CustomRecovery(middleware.ErrorMiddleware))

	author := router.Group("/author")
	authorController := init.AuthorController
	{
		author.GET("/", authorController.Index)
		author.GET("/:id", authorController.Show)
		author.POST("/", authorController.Create)
		author.PUT("/:id", authorController.Update)
		author.DELETE("/:id", authorController.Delete)
	}

	return router
}
