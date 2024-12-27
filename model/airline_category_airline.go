package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirlineCategoryAirline table one to many / many to many airlineAPI to Airline Category
type AirlineCategoryAirline struct {
	basic.Base
	basic.DataOwner
	AirlineCategoryAirlineAPI
	Airline         *Airline         `json:"airline,omitempty"`
	AirlineCategory *AirlineCategory `json:"airline_category,omitempty"`
}

// AirlineCategoryAirlineAPI struct writable by request body
type AirlineCategoryAirlineAPI struct {
	AirlineCategoryID *uuid.UUID `json:"airline_category_id,omitempty" gorm:"type:varchar(36);not null;"`
	AirlineID         *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36);not null"`
}
