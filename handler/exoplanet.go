package handler

import (
	"errors"
	"exoplanet-service/model"
	"exoplanet-service/repository"
	"exoplanet-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExoplanetHandler struct {
	Repo repository.ExoplanetRepository
}

func NewExoplanetHandler(repo repository.ExoplanetRepository) *ExoplanetHandler {
	return &ExoplanetHandler{Repo: repo}
}

func (h *ExoplanetHandler) AddExoplanet(c *gin.Context) {
	var exoplanet model.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dataValidation(exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Repo.AddExoplanet(&exoplanet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exoplanet)
}

func (h *ExoplanetHandler) ListExoplanets(c *gin.Context) {
	exoplanets := h.Repo.ListExoplanets()
	c.JSON(http.StatusOK, exoplanets)
}

func (h *ExoplanetHandler) GetExoplanetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	exoplanet, err := h.Repo.GetExoplanetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	c.JSON(http.StatusOK, exoplanet)
}

func (h *ExoplanetHandler) UpdateExoplanet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var exoplanet model.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dataValidation(exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Repo.UpdateExoplanet(id, &exoplanet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exoplanet)
}

func (h *ExoplanetHandler) DeleteExoplanet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.Repo.DeleteExoplanet(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ExoplanetHandler) FuelEstimation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	exoplanet, err := h.Repo.GetExoplanetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	crewCapacity, err := strconv.Atoi(c.Query("crewCapacity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crew capacity"})
		return
	}

	fuelCost := service.CalculateFuelCost(exoplanet, crewCapacity)
	c.JSON(http.StatusOK, gin.H{"fuelCost": fuelCost})
}

func dataValidation(exoplanet model.Exoplanet) (err error) {
	switch exoplanet.Type {
	case "Terrestrial":
		if !(exoplanet.Mass > 0.1 && exoplanet.Mass < 10) {
			err = errors.New("invalid value of mass")
		}
		return
	case "GasGiant":
		return
	default:
		err = errors.New("invalid value of exoplanet type")
		return
	}
}
