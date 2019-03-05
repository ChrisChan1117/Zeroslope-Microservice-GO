package router

import (
	"../config"
	"../controllers"
	"../middlewares"

	// Use the swagger docs
	_ "../docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetupRouter creaes a router using middleware and controllers
func SetupRouter(cfg config.Configuration) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redirect the root to swagger
	router.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/swagger/index.html"
		router.HandleContext(c)
	})

	// Health endpoint
	health := new(controllers.HealthController)
	healthRoutes := router.Group("/health")
	{
		healthRoutes.GET("/", health.Status)
	}

	// Authentication endpoint
	auth := new(controllers.AuthController)
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", auth.Login)
	}

	// All routes below this are authenticated
	router.Use(middlewares.AuthMiddleware(cfg))

	// Sample endpoints
	sample := new(controllers.SampleController)
	sampleRoutes := router.Group("/samples")
	{
		sampleRoutes.GET("/", sample.List)
		sampleRoutes.GET("/:id", sample.Read)
		sampleRoutes.POST("/", sample.Create)
		sampleRoutes.PUT("/", sample.Update)
		sampleRoutes.DELETE("/:id", sample.List)
	}

	return router
}
