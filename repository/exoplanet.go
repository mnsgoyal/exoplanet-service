package repository

import (
	"errors"
	"exoplanet-service/model"
	"sync"
)

type ExoplanetRepository interface {
	AddExoplanet(exoplanet *model.Exoplanet) error
	ListExoplanets() []model.Exoplanet
	GetExoplanetByID(id int) (*model.Exoplanet, error)
	UpdateExoplanet(id int, exoplanet *model.Exoplanet) error
	DeleteExoplanet(id int) error
	storeExoplanetData()
}

type InMemoryExoplanetRepository struct {
	mu         sync.RWMutex
	exoplanets map[int]model.Exoplanet
	nextID     int
}

func NewInMemoryExoplanetRepository() *InMemoryExoplanetRepository {
	existingExoplanetData := readExoplanetData()
	if existingExoplanetData.Exoplanets != nil {
		return &InMemoryExoplanetRepository{
			exoplanets: existingExoplanetData.Exoplanets,
			nextID:     existingExoplanetData.NextID,
		}
	}

	return &InMemoryExoplanetRepository{
		exoplanets: map[int]model.Exoplanet{},
		nextID:     1,
	}
}

func TestNewInMemoryExoplanetRepository() *InMemoryExoplanetRepository {
	return &InMemoryExoplanetRepository{
		exoplanets: map[int]model.Exoplanet{},
		nextID:     1,
	}
}

func (r *InMemoryExoplanetRepository) AddExoplanet(exoplanet *model.Exoplanet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	defer r.storeExoplanetData()

	exoplanet.ID = r.nextID
	r.nextID++
	r.exoplanets[exoplanet.ID] = *exoplanet

	return nil
}

func (r *InMemoryExoplanetRepository) ListExoplanets() []model.Exoplanet {
	r.mu.RLock()
	defer r.mu.RUnlock()

	exoplanetsList := []model.Exoplanet{}

	for _, exoplanetInfo := range r.exoplanets {
		exoplanetsList = append(exoplanetsList, exoplanetInfo)
	}

	return exoplanetsList
}

func (r *InMemoryExoplanetRepository) GetExoplanetByID(id int) (*model.Exoplanet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if exoplanetInfo, okay := r.exoplanets[id]; okay {
		return &exoplanetInfo, nil
	}
	return nil, errors.New("exoplanet not found")
}

func (r *InMemoryExoplanetRepository) UpdateExoplanet(id int, exoplanet *model.Exoplanet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	defer r.storeExoplanetData()

	if _, okay := r.exoplanets[id]; okay {
		r.exoplanets[id] = *exoplanet
		return nil

	}
	return errors.New("exoplanet not found")
}

func (r *InMemoryExoplanetRepository) DeleteExoplanet(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	defer r.storeExoplanetData()

	if _, okay := r.exoplanets[id]; okay {
		delete(r.exoplanets, id)
	}
	return errors.New("exoplanet not found")
}
