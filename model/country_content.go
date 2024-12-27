package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CountryContent Country Content
type CountryContent struct {
	basic.Base
	basic.DataOwner
	CountryContentAPI
	Country            *Country            `json:"country,omitempty" gorm:"foreignKey:CountryID;references:ID"`
	ContentDescription *ContentDescription `json:"content_description,omitempty" gorm:"foreignKey:ContentDescriptionID;references:ID"`
}

// CountryContentAPI Country Content API
type CountryContentAPI struct {
	CountryID            *uuid.UUID             `json:"country_id,omitempty" swaggertype:"string" format:"uuid" validate:"required"` // Country ID
	ContentDescriptionID *uuid.UUID             `json:"content_description_id,omitempty" swaggertype:"string" format:"uuid"`         // Content Description ID
	ContentDescription   *ContentDescriptionAPI `json:"content_description,omitempty" gorm:"-"`
}
