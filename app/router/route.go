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

	publisher := router.Group("/publisher")
	publisherController := init.PublisherController
	{
		publisher.GET("/", publisherController.Index)
		publisher.GET("/:id", publisherController.Show)
		publisher.POST("/", publisherController.Create)
		publisher.PUT("/:id", publisherController.Update)
		publisher.DELETE("/:id", publisherController.Delete)
	}

	book := router.Group("/book")
	bookController := init.BookController
	{
		book.GET("/", bookController.Index)
		book.GET("/:id", bookController.Show)
		book.POST("/", bookController.Create)
		book.PUT("/:id", bookController.Update)
		book.DELETE("/:id", bookController.Delete)
	}

	return router
}
