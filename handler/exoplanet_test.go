package handler

import (
	"exoplanet-service/model"
	"exoplanet-service/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddExoplanet(t *testing.T) {
	repo := repository.TestNewInMemoryExoplanetRepository()
	handler := NewExoplanetHandler(repo)

	router := gin.Default()
	router.POST("/exoplanets", handler.AddExoplanet)

	exoplanet := `{
		"name": "New Exoplanet",
		"description": "A new exoplanet description",
		"distance": 123,
		"radius": 0.21,
		"type": "GasGiant"
	}`

	req, _ := http.NewRequest("POST", "/exoplanets", strings.NewReader(exoplanet))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestListExoplanets(t *testing.T) {
	repo := repository.TestNewInMemoryExoplanetRepository()
	handler := NewExoplanetHandler(repo)

	router := gin.Default()
	router.GET("/exoplanets", handler.ListExoplanets)

	req, _ := http.NewRequest("GET", "/exoplanets", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetExoplanetByID(t *testing.T) {
	repo := repository.TestNewInMemoryExoplanetRepository()
	handler := NewExoplanetHandler(repo)

	repo.AddExoplanet(&model.Exoplanet{
		Name:        "Existing Exoplanet",
		Description: "An existing exoplanet",
		Distance:    100,
		Radius:      0.2,
		Mass:        0.2,
		Type:        "Terrestrial",
	})

	router := gin.Default()
	router.GET("/exoplanets/:id", handler.GetExoplanetByID)

	req, _ := http.NewRequest("GET", "/exoplanets/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateExoplanet(t *testing.T) {
	repo := repository.TestNewInMemoryExoplanetRepository()
	handler := NewExoplanetHandler(repo)

	repo.AddExoplanet(&model.Exoplanet{
		Name:        "Old Exoplanet",
		Description: "An old exoplanet description",
		Distance:    200,
		Radius:      0.3,
		Type:        "GasGiant",
	})

	router := gin.Default()
	router.PUT("/exoplanets/:id", handler.UpdateExoplanet)

	updatedExoplanet := `{
		"name": "Updated Exoplanet",
		"description": "An updated exoplanet description",
		"distance": 300,
		"radius": 0.4,
		"type": "GasGiant"
	}`

	req, _ := http.NewRequest("PUT", "/exoplanets/1", strings.NewReader(updatedExoplanet))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// func TestDeleteExoplanet(t *testing.T) {
// 	repo := repository.TestNewInMemoryExoplanetRepository()
// 	handler := NewExoplanetHandler(repo)

// 	repo.AddExoplanet(&model.Exoplanet{
// 		Name:        "Existing Exoplanet",
// 		Description: "An existing exoplanet",
// 		Distance:    100,
// 		Radius:      0.2,
// 		Mass:        0.2,
// 		Type:        "Terrestrial",
// 	})

// 	router := gin.Default()

// 	req1, _ := http.NewRequest("GET", "/exoplanets/1", nil)
// 	w1 := httptest.NewRecorder()
// 	router.ServeHTTP(w1, req1)

// 	assert.Equal(t, http.StatusOK, w1.Code)

// 	router.DELETE("/exoplanets/:id", handler.DeleteExoplanet)

// 	req, _ := http.NewRequest("DELETE", "/exoplanets/1", nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)
// }

func TestDeleteExoplanet(t *testing.T) {
	repo := repository.TestNewInMemoryExoplanetRepository()
	handler := NewExoplanetHandler(repo)

	repo.AddExoplanet(&model.Exoplanet{
		Name:        "Existing Exoplanet",
		Description: "An existing exoplanet",
		Distance:    100,
		Radius:      0.2,
		Mass:        0.2,
		Type:        "Terrestrial",
	})

	router := gin.Default()
	router.DELETE("/exoplanets/:id", handler.DeleteExoplanet)

	req, _ := http.NewRequest("DELETE", "/exoplanets/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestFuelEstimation(t *testing.T) {
	repo := repository.TestNewInMemoryExoplanetRepository()
	handler := NewExoplanetHandler(repo)

	repo.AddExoplanet(&model.Exoplanet{
		Name:        "Exoplanet",
		Description: "Exoplanet description",
		Distance:    150,
		Radius:      0.5,
		Type:        "GasGiant",
	})

	router := gin.Default()
	router.GET("/exoplanets/:id/fuel", handler.FuelEstimation)

	req, _ := http.NewRequest("GET", "/exoplanets/1/fuel?crewCapacity=5", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
