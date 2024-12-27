package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// FareType Fare Type
type FareType struct {
	basic.Base
	basic.DataOwner
	FareTypeAPI
	FareTypeTranslation *FareTypeTranslation `json:"fare_type_translation,omitempty"` // Fare Type Translation
}

// FareTypeAPI Fare Type API
type FareTypeAPI struct {
	FareTypeCode *string `json:"fare_type_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" example:"Lite"`  // Fare Type Code
	FareTypeName *string `json:"fare_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Lite"` // Fare Type Name
}

// Seed data
func (s FareType) Seed() *[]FareType {
	data := []FareType{}
	items := []string{
		"Lite",
		"Value",
		"Standard",
		"Flexi",
		"First",
	}

	for i := range items {
		var name string = items[i]
		data = append(data, FareType{
			FareTypeAPI: FareTypeAPI{
				FareTypeCode: &name,
				FareTypeName: &name,
			},
		})
	}

	return &data
}
