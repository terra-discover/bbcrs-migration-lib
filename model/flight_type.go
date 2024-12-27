package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// FlightType Flight Type
type FlightType struct {
	basic.Base
	basic.DataOwner
	FlightTypeAPI
	FlightTypeTranslation *FlightTypeTranslation `json:"flight_type_translation,omitempty"`
}

// FlightTypeAPI Flight Type API
type FlightTypeAPI struct {
	FlightTypeCode *string `json:"flight_type_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" example:"FC-01"`   // Flight Type Code
	FlightTypeName *string `json:"flight_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Forces"` // Flight Type Name
}
