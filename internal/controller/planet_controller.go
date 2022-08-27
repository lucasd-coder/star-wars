package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/star-wars/internal/interfaces"
	"github.com/lucasd-coder/star-wars/internal/models"
	"github.com/lucasd-coder/star-wars/internal/service"
	"github.com/lucasd-coder/star-wars/pkg/logger"
)

type PlanetController struct {
	createPlanetService interfaces.CreatePlanetService
}

func NewPlanetController(createPlanetService *service.CreatePlanetService) *PlanetController {
	return &PlanetController{
		createPlanetService,
	}
}

func (planet *PlanetController) InitRoutes(group *gin.RouterGroup) {
	group.POST("/planets", planet.CreatePlanet)
}

func (planet *PlanetController) CreatePlanet(ctx *gin.Context) {
	var body models.PlanetDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logger.Log.Error(err.Error())
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	err := planet.createPlanetService.Execute(&body)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.Status(200)
}
