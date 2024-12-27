package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// DestinationSystemType Destination System Type
type DestinationSystemType struct {
	basic.Base
	basic.DataOwner
	DestinationSystemTypeAPI
	DestinationSystemTypeTranslation *DestinationSystemTypeTranslation `json:"destination_system_type_translation,omitempty"` // Destination System Type Translation
}

// DestinationSystemTypeAPI Destination System Type API
type DestinationSystemTypeAPI struct {
	DestinationSystemTypeCode *int    `json:"destination_system_type_code,omitempty" gorm:"type:int;index:,unique,where:deleted_at is null;not null" example:"1"`               // Destination System Type Code
	DestinationSystemTypeName *string `json:"destination_system_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Sample"` // Destination System Type Name
}
