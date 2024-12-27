package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// OccupancyType Occupancy Type
type OccupancyType struct {
	basic.Base
	basic.DataOwner
	OccupancyTypeAPI
	OccupancyTypeTranslation *OccupancyTypeTranslation `json:"occupancy_type_translation,omitempty"`
}

// OccupancyTypeAPI Occupancy Type API
type OccupancyTypeAPI struct {
	OccupancyTypeCode *string `json:"occupancy_type_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`  // Occupancy Type Code
	OccupancyTypeName *string `json:"occupancy_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Occupancy Type Name
	Occupancy         *int    `json:"occupancy,omitempty" gorm:"type:smallint;not null;" example:"1"`
	IsDefault         *bool   `json:"is_default,omitempty"` // Is Default
}
