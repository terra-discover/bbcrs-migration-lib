package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DestinationGroupCountry model
type DestinationGroupCountry struct {
	basic.Base
	basic.DataOwner
	DestinationGroupCountryAPI
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" swaggertype:"string" format:"uuid"` // Destination Group ID
	Country            *Country   `json:"country,omitempty"`                                                 // Country
}

// DestinationGroupCountryAPI model
type DestinationGroupCountryAPI struct {
	CountryID *uuid.UUID `json:"country_id,omitempty" swaggertype:"string" format:"uuid"` // country id
}
