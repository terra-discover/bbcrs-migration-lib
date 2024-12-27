package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// TripType Trip Type
type TripType struct {
	basic.Base
	basic.DataOwner
	TripTypeAPI
	TripTypeTranslation *TripTypeTranslation `json:"trip_type_translation,omitempty"`
}

// TripTypeAPI Trip Type API
type TripTypeAPI struct {
	TripTypeCode *string `json:"trip_type_code,omitempty" gorm:"type:varchar(36);not null;index:unique_trip_type__trip_type_code,unique,where:deleted_at is null"`  // Trip Type Code
	TripTypeName *string `json:"trip_type_name,omitempty" gorm:"type:varchar(256);not null;index:unique_trip_type__trip_type_name,unique,where:deleted_at is null"` // Trip Type Name
	IsDefault    *bool   `json:"is_default,omitempty"`                                                                                                              // Is Default
}

// GetTripType by query filter
func (data *TripType) GetTripType(tx *gorm.DB, queryFilters string) *TripType {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Take(&data)
	return data
}

// Seed data
func (s TripType) Seed() *[]TripType {
	data := []TripType{}
	items := []string{
		"Oneway|One-way|0|",
		"Roundtrip|Round trip|1|",
		"Multicity|Multi-city|0|",
	}
	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		code := c[0]
		name := c[1]
		boolCase := c[2]
		isDefault := false
		if boolCase == "1" {
			isDefault = true
		}
		data = append(data, TripType{
			TripTypeAPI: TripTypeAPI{
				TripTypeCode: &code,
				TripTypeName: &name,
				IsDefault:    &isDefault,
			},
		})
	}
	return &data
}
