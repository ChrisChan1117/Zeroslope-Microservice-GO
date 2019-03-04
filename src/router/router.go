package router

import (
	"../controllers"
	"../middlewares"
	"../utilities"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	// Import the docs
	_ "../docs"
)

// SetupRouter creaes a router using middleware and controllers
func SetupRouter(cfg utilities.Configuration) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Expose swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redirect the root to swagger
	router.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/swagger"
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
	sampleRoutes := router.Group("/sample")
	{
		sampleRoutes.GET("/", sample.List)
		sampleRoutes.GET("/:id", sample.Read)
		sampleRoutes.POST("/", sample.Create)
		sampleRoutes.PUT("/", sample.Update)
		sampleRoutes.DELETE("/:id", sample.List)
	}

	return router
}
