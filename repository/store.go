package repository

import (
	"encoding/json"
	"exoplanet-service/common"
	"exoplanet-service/model"
	"fmt"
	"io/ioutil"
)

type ExoplanetInfo struct {
	Exoplanets map[int]model.Exoplanet
	NextID     int
}

func (r *InMemoryExoplanetRepository) storeExoplanetData() {
	// Marshal data into JSON
	var objExoplanetRepository ExoplanetInfo
	objExoplanetRepository.Exoplanets = r.exoplanets
	objExoplanetRepository.NextID = r.nextID
	jsonData, err := json.Marshal(objExoplanetRepository)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON data to a file
	err = ioutil.WriteFile(common.FilenameForData, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
	return
}

func readExoplanetData() ExoplanetInfo {
	var objExoplanetRepository ExoplanetInfo
	// Write JSON data to a file
	jsonData, err := ioutil.ReadFile(common.FilenameForData)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return objExoplanetRepository
	}

	err = json.Unmarshal(jsonData, &objExoplanetRepository)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return objExoplanetRepository
	}

	return objExoplanetRepository
}
