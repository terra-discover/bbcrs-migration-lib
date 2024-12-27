package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// ReferencePointCategory Reference Point Category
type ReferencePointCategory struct {
	basic.Base
	basic.DataOwner
	ReferencePointCategoryAPI
	ReferencePointCategoryTranslation *ReferencePointCategoryTranslation `json:"reference_point_category_translation,omitempty"` // Reference Point Category Translation
}

// ReferencePointCategoryAPI Reference Point Category API
type ReferencePointCategoryAPI struct {
	ReferencePointCategoryCode *int    `json:"reference_point_category_code,omitempty" gorm:"type:int;index:,unique,where:deleted_at is null;not null" example:"1"`                // Reference Point Category Code
	ReferencePointCategoryName *string `json:"reference_point_category_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Airport"` // Reference Point Category Name
}

// Seed data
func (s ReferencePointCategory) Seed() *[]ReferencePointCategory {
	data := []ReferencePointCategory{}
	items := []string{
		"Bar",
		"Bay",
		"Beach",
		"Boat dock",
		"Bus station",
		"Church",
		"City center",
		"Corporation",
		"Educational institution",
		"Ferry station",
		"Financial district",
		"Financial institution",
		"Lake",
		"Landmark",
		"Library",
		"Marina",
		"Market",
		"Medical facility",
		"Metro/subway station ",
		"Monument",
		"Museum",
		"Park",
		"Racetrack",
		"Restaurant",
		"River",
		"School",
		"Shopping center",
		"Sports facility",
		"Synagogue",
		"Town center",
		"Train station",
		"University",
		"Zoo",
		"Local area",
		"Point of interest",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, ReferencePointCategory{
			ReferencePointCategoryAPI: ReferencePointCategoryAPI{
				ReferencePointCategoryCode: &code,
				ReferencePointCategoryName: &name,
			},
		})
	}

	return &data
}
