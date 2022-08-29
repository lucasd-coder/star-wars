package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/star-wars/internal/mock"
	"github.com/stretchr/testify/assert"
	mockTestify "github.com/stretchr/testify/mock"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func main() {
	r := SetUpRouter()
	r.Run(":8080")
}

func TestFindAllPlanetSuccessfully(t *testing.T) {
	createPlanetService := new(mock.MockCreatePlanetService)
	findPlanetByIdService := new(mock.MockFindPlanetByIdService)
	findPlanetByNameService := new(mock.MockFindPlanetByNameService)
	findAllPlanetService := new(mock.MockFindAllPlanetService)
	deletePlanetByIdService := new(mock.MockDeletePlanetByIdService)

	testController := PlanetController{
		createPlanetService, findPlanetByIdService,
		findPlanetByNameService, findAllPlanetService, deletePlanetByIdService,
	}

	planetsResponse := mock.FactoryPlanetsResponse()

	findAllPlanetService.On("Execute").Return(planetsResponse, nil)

	r := SetUpRouter()
	r.GET("/", testController.FindAll)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFailPlanetAlreadyExist(t *testing.T) {
	createPlanetService := new(mock.MockCreatePlanetService)
	findPlanetByIdService := new(mock.MockFindPlanetByIdService)
	findPlanetByNameService := new(mock.MockFindPlanetByNameService)
	findAllPlanetService := new(mock.MockFindAllPlanetService)
	deletePlanetByIdService := new(mock.MockDeletePlanetByIdService)

	testController := PlanetController{
		createPlanetService, findPlanetByIdService,
		findPlanetByNameService, findAllPlanetService, deletePlanetByIdService,
	}

	createPlanetService.On("Execute", mockTestify.Anything).Return(nil, errors.New("Planet already exist!"))

	r := SetUpRouter()
	r.POST("/", testController.CreatePlanet)
	req, _ := http.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreatePlanetSuccessfully(t *testing.T) {
	createPlanetService := new(mock.MockCreatePlanetService)
	findPlanetByIdService := new(mock.MockFindPlanetByIdService)
	findPlanetByNameService := new(mock.MockFindPlanetByNameService)
	findAllPlanetService := new(mock.MockFindAllPlanetService)
	deletePlanetByIdService := new(mock.MockDeletePlanetByIdService)

	testController := PlanetController{
		createPlanetService, findPlanetByIdService,
		findPlanetByNameService, findAllPlanetService, deletePlanetByIdService,
	}

	dto := mock.FactoryPlanetDTO()

	createPlanetService.On("Execute", mockTestify.Anything).Return(nil)

	jsonValue, _ := json.Marshal(dto)
	r := SetUpRouter()
	r.POST("/", testController.CreatePlanet)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFailRequiredFields(t *testing.T) {
	createPlanetService := new(mock.MockCreatePlanetService)
	findPlanetByIdService := new(mock.MockFindPlanetByIdService)
	findPlanetByNameService := new(mock.MockFindPlanetByNameService)
	findAllPlanetService := new(mock.MockFindAllPlanetService)
	deletePlanetByIdService := new(mock.MockDeletePlanetByIdService)

	testController := PlanetController{
		createPlanetService, findPlanetByIdService,
		findPlanetByNameService, findAllPlanetService, deletePlanetByIdService,
	}

	createPlanetService.On("Execute", mockTestify.Anything).Return(nil)

	r := SetUpRouter()
	r.POST("/", testController.CreatePlanet)
	req, _ := http.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFindPlanetByIdSuccessfully(t *testing.T) {
	createPlanetService := new(mock.MockCreatePlanetService)
	findPlanetByIdService := new(mock.MockFindPlanetByIdService)
	findPlanetByNameService := new(mock.MockFindPlanetByNameService)
	findAllPlanetService := new(mock.MockFindAllPlanetService)
	deletePlanetByIdService := new(mock.MockDeletePlanetByIdService)

	testController := PlanetController{
		createPlanetService, findPlanetByIdService,
		findPlanetByNameService, findAllPlanetService, deletePlanetByIdService,
	}

	var id string = "630b7bcb419f837457644cbc"

	planet := mock.FactoryPlanetResponse()

	findPlanetByIdService.On("Execute", id).Return(planet, nil)

	r := SetUpRouter()
	r.GET("/key/:id", testController.FindById)

	req, _ := http.NewRequest("GET", "/key/"+id, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFindPlanetByNameSuccessfully(t *testing.T) {
	createPlanetService := new(mock.MockCreatePlanetService)
	findPlanetByIdService := new(mock.MockFindPlanetByIdService)
	findPlanetByNameService := new(mock.MockFindPlanetByNameService)
	findAllPlanetService := new(mock.MockFindAllPlanetService)
	deletePlanetByIdService := new(mock.MockDeletePlanetByIdService)

	testController := PlanetController{
		createPlanetService, findPlanetByIdService,
		findPlanetByNameService, findAllPlanetService, deletePlanetByIdService,
	}

	var name string = "lucas"

	planet := mock.FactoryPlanetResponse()

	findPlanetByNameService.On("Execute", name).Return(planet, nil)

	r := SetUpRouter()
	r.GET("/:name", testController.FindByName)

	req, _ := http.NewRequest("GET", "/"+name, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFailDeleteSuccessfully(t *testing.T) {
	createPlanetService := new(mock.MockCreatePlanetService)
	findPlanetByIdService := new(mock.MockFindPlanetByIdService)
	findPlanetByNameService := new(mock.MockFindPlanetByNameService)
	findAllPlanetService := new(mock.MockFindAllPlanetService)
	deletePlanetByIdService := new(mock.MockDeletePlanetByIdService)

	testController := PlanetController{
		createPlanetService, findPlanetByIdService,
		findPlanetByNameService, findAllPlanetService, deletePlanetByIdService,
	}

	var id string = "630b7bcb419f837457644cbc"

	deletePlanetByIdService.On("Execute", id).Return(nil)

	r := SetUpRouter()
	r.DELETE("/key/:id", testController.Delete)

	req, _ := http.NewRequest("DELETE", "/key/"+id, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
}
