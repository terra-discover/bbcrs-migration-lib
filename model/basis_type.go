package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// BasisType Basis Type
type BasisType struct {
	basic.Base
	basic.DataOwner
	BasisTypeAPI
	BasisTypeTranslation *BasisTypeTranslation `json:"basis_type_translation,omitempty"` // Basis Type Translation
}

// BasisTypeAPI Basis Type API
type BasisTypeAPI struct {
	BasisTypeCode *int    `json:"basis_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`         // Basis Type Code
	BasisTypeName *string `json:"basis_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Daily"` // Basis Type Name
}

// Seed data
func (s BasisType) Seed() *[]BasisType {
	data := []BasisType{}
	items := []string{
		"Daily",
		"Hourly",
		"Half day",
		"Additions per stay",
		"Per occurrence",
		"Per event",
		"Per person",
		"First use",
		"One time use",
		"Per minute",
		"Per function",
		"Per stay",
		"Complimentary",
		"Other",
		"Maximum charge",
		"Over-minute charge",
		"Weekly",
		"Per room per stay",
		"Per room per night",
		"Per person per stay",
		"Per person per night",
		"Minimum charge",
		"Per rental",
		"Per item",
		"Per room",
		"Per reservation/booking",
		"Per gallon",
		"Per dozen",
		"Per tray",
		"Per order",
		"Per unit",
		"One way",
		"Round trip",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, BasisType{
			BasisTypeAPI: BasisTypeAPI{
				BasisTypeCode: &code,
				BasisTypeName: &name,
			},
		})
	}

	return &data
}
