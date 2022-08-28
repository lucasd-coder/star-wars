package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/internal/controller"
	"github.com/lucasd-coder/star-wars/internal/middlewares"
	"github.com/lucasd-coder/star-wars/internal/pkg/mongodb"
	"github.com/lucasd-coder/star-wars/pkg/cache"
	"github.com/lucasd-coder/star-wars/pkg/logger"
)

func Run(cfg *config.Config) {
	// Log config
	logger.SetUpLog(cfg)

	// Mongo Config
	mongodb.SetUpMongoDB(cfg)

	// Redis Config
	cache.SetUpRedis(cfg)

	// Clode MongoDB
	defer mongodb.CloseConnMongoDB()

	// Http server
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middlewares.JSONAppErrorReporter())

	// Routers
	handler := engine.Group("/" + cfg.Name)
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	createPlanetService := InitializeCreatePlanetService()
	findPlanetByIdService := InitializeFindPlanetByIdService()
	findPlanetByNameService := InitializeFindPlanetByNameService()
	planetController := controller.NewPlanetController(createPlanetService, findPlanetByIdService, findPlanetByNameService)

	planetController.InitRoutes(handler)

	err := engine.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
