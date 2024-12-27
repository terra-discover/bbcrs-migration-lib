package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateDestinationRestrictionCity Corporate Destination Restriction City
type CorporateDestinationRestrictionCity struct {
	basic.Base
	basic.DataOwner
	CorporateDestinationRestrictionID *uuid.UUID `json:"corporate_destination_restriction_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_destination_restriction_city,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	CorporateDestinationRestrictionCityAPI
	Country *Country `json:"country,omitempty"`
	City    *City    `json:"city,omitempty"`
}

// CorporateDestinationRestrictionCityAPI Corporate Destination Restriction City Api
type CorporateDestinationRestrictionCityAPI struct {
	CountryID *uuid.UUID `json:"country_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_destination_restriction_city,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	CityID    *uuid.UUID `json:"city_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_destination_restriction_city,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
}
