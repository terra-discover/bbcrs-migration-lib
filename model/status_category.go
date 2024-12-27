package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// StatusCategory struct
type StatusCategory struct {
	basic.Base
	basic.DataOwner
	StatusCategoryCode *int    `json:"status_category_code,omitempty" gorm:"type:smallint;not null"`     // Status Category Code
	StatusCategoryName *string `json:"status_category_name,omitempty" gorm:"type:varchar(256);not null"` // Status Category Name
}

// Seed data
func (s StatusCategory) Seed() *[]StatusCategory {
	data := []StatusCategory{}
	items := []string{"Ticket Status"}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, StatusCategory{
			StatusCategoryCode: &code,
			StatusCategoryName: &name,
		})
	}

	return &data
}
