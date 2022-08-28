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
	findAllPlanetService    interfaces.FindAllPlanetService
	deletePlanetByIdService interfaces.DeletePlanetByIdService
}

func NewPlanetController(createPlanetService *service.CreatePlanetService,
	findPlanetByIdService *service.FindPlanetByIdService,
	findPlanetByNameService *service.FindPlanetByNameService,
	findAllPlanetService *service.FindAllPlanetService,
	deletePlanetByIdService *service.DeletePlanetByIdService,
) *PlanetController {
	return &PlanetController{
		createPlanetService,
		findPlanetByIdService,
		findPlanetByNameService,
		findAllPlanetService,
		deletePlanetByIdService,
	}
}

func (planet *PlanetController) InitRoutes(group *gin.RouterGroup) {
	planets := group.Group("/planets")
	{
		planets.POST("", planet.CreatePlanet)
		planets.GET("key/:id", planet.FindById)
		planets.GET("/:name", planet.FindByName)
		planets.GET("", planet.FindAll)
		planets.DELETE("key/:id", planet.Delete)
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

	ctx.Status(201)
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

func (planet *PlanetController) FindAll(ctx *gin.Context) {
	resp, err := planet.findAllPlanetService.Execute()
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(200, resp)
}

func (planet *PlanetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := planet.deletePlanetByIdService.Execute(id)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.Status(202)
}
