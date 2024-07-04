package models

type Rating struct {
	BrandID string `json:"brand_id"`
	Variety string `json:"variety"`
	Rating  string `json:"rating"`
}
