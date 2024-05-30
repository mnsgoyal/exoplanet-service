package model

type Exoplanet struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Distance    int     `json:"distance" binding:"required,gt=10,lt=1000"`
	Radius      float64 `json:"radius" binding:"required,gt=0.1,lt=10"`
	Mass        float64 `json:"mass,omitempty"` // Only for Terrestrial
	Type        string  `json:"type" binding:"required"`
}
