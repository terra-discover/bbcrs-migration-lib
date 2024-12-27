package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingFrequentDestination struct {
	basic.Base
	basic.DataOwner
	FlightCachingFrequentDestinationAPI
	OriginCity      *City `json:"origin_city" gorm:"foreignKey:OriginCityID;references:ID"`
	DestinationCity *City `json:"destination_city" gorm:"foreignKey:DestinationCityID;references:ID"`
}

type FlightCachingFrequentDestinationAPI struct {
	OriginCityID      *uuid.UUID `json:"origin_city_id" gorm:"type:varchar(36);not null"`
	DestinationCityID *uuid.UUID `json:"destination_city_id" gorm:"type:varchar(36);not null"`
}
