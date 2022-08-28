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
	createPlanetService     interfaces.CreatePlanetService
	findPlanetByIdService   interfaces.FindPlanetByIdService
	findPlanetByNameService interfaces.FindPlanetByNameService
}

func NewPlanetController(createPlanetService *service.CreatePlanetService,
	findPlanetByIdService *service.FindPlanetByIdService,
	findPlanetByNameService *service.FindPlanetByNameService,
) *PlanetController {
	return &PlanetController{
		createPlanetService,
		findPlanetByIdService,
		findPlanetByNameService,
	}
}

func (planet *PlanetController) InitRoutes(group *gin.RouterGroup) {
	planets := group.Group("/planets")
	{
		planets.POST("", planet.CreatePlanet)
		planets.GET("key/:id", planet.FindById)
		planets.GET("/:name", planet.FindByName)
	}
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

func (planet *PlanetController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := planet.findPlanetByIdService.Execute(id)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(200, resp)
}

func (planet *PlanetController) FindByName(ctx *gin.Context) {
	name := ctx.Param("name")

	resp, err := planet.findPlanetByNameService.Execute(name)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(200, resp)
}
