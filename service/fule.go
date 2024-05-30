package service

import (
	"exoplanet-service/model"
	"math"
)

func CalculateFuelCost(exoplanet *model.Exoplanet, crewCapacity int) float64 {
	var gravity float64
	switch exoplanet.Type {
	case "GasGiant":
		gravity = float64(0.5) * math.Pow(exoplanet.Radius, float64(2))
	case "Terrestrial":
		gravity = exoplanet.Mass * math.Pow(exoplanet.Radius, float64(2))
	default:
		gravity = 0.5
	}
	fuelCost := (float64(exoplanet.Distance) / math.Pow(gravity, float64(2))) * float64(crewCapacity)
	return float64(int(fuelCost*100)) / 100
}
