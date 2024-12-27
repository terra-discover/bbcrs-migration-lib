package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreCacheSelected struct {
	basic.Base
	basic.DataOwner
	SelectedSessionID *uuid.UUID                `json:"selected_session_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	FlightID          *uuid.UUID                `json:"flight_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	SelectedRoutes    []SabreCacheSelectedRoute `json:"selected_route,omitempty" gorm:"foreignKey:SelectedFlightID;references:ID"`
	SelectedFare      SabreCacheSelectedFare    `json:"selected_fare,omitempty" gorm:"foreignKey:SelectedFlightID;references:ID"`
}

type SabreCacheSelectedRoute struct {
	SelectedSessionID *uuid.UUID `json:"selected_session_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	SelectedFlightID  *uuid.UUID `json:"selected_flight_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	SelectedRouteID   *uuid.UUID `json:"selected_route_id,omitempty" gorm:"type:varchar(36);" format:"uuid"` //flight_caching_route_id
	SabreCacheRoutes
}

type SabreCacheSelectedFare struct {
	SelectedSessionID *uuid.UUID `json:"selected_session_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	SelectedFlightID  *uuid.UUID `json:"selected_flight_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	SelectedFareID    *uuid.UUID `json:"selected_fare_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	SabreCacheFares
}
