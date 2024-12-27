package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PageDestination struct
type PageDestination struct {
	basic.Base
	basic.DataOwner
	PageDestinationAPI
	Page          *Page          `json:"page,omitempty"`           // Page
	City          *City          `json:"city,omitempty"`           // City
	Country       *Country       `json:"country,omitempty"`        // Country
	StateProvince *StateProvince `json:"state_province,omitempty"` // State Province
	Destination   *Destination   `json:"destination,omitempty"`    // Destination
}

// PageDestinationAPI struct
type PageDestinationAPI struct {
	PageID          *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`  // Page ID
	CityID          *uuid.UUID `json:"city_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`           // City ID
	CountryID       *uuid.UUID `json:"country_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`        // Country ID
	StateProvinceID *uuid.UUID `json:"state_province_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // State Province ID
	DestinationID   *uuid.UUID `json:"destination_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`    // Destination ID
}
