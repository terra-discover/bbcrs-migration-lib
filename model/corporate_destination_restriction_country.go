package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateDestinationRestrictionCountry Corporate Destination Restriction Country
type CorporateDestinationRestrictionCountry struct {
	basic.Base
	basic.DataOwner
	CorporateDestinationRestrictionID *uuid.UUID `json:"corporate_destination_restriction_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_destination_restriction_country,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	CorporateDestinationRestrictionCountryAPI
	Country *Country `json:"country,omitempty"`
}

// CorporateDestinationRestrictionCountryAPI Corporate Destination Restriction Country Api
type CorporateDestinationRestrictionCountryAPI struct {
	CountryID *uuid.UUID `json:"country_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_destination_restriction_country,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
}
