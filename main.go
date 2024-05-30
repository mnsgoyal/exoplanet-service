package main

import (
	"exoplanet-service/handler"
	"exoplanet-service/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	exoplanetRepo := repository.NewInMemoryExoplanetRepository()
	exoplanetHandler := handler.NewExoplanetHandler(exoplanetRepo)
	r.POST("/exoplanets", exoplanetHandler.AddExoplanet)
	r.GET("/exoplanets", exoplanetHandler.ListExoplanets)
	r.GET("/exoplanets/:id", exoplanetHandler.GetExoplanetByID)
	r.PUT("/exoplanets/:id", exoplanetHandler.UpdateExoplanet)
	r.DELETE("/exoplanets/:id", exoplanetHandler.DeleteExoplanet)
	r.GET("/exoplanets/fuel-estimation/:id", exoplanetHandler.FuelEstimation)

	r.Run(":8080")
}
