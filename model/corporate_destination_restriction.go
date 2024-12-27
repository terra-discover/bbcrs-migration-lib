package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateDestinationRestriction model
type CorporateDestinationRestriction struct {
	basic.Base
	basic.DataOwner
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Corporate ID
	Corporate   *Corporate `json:"-"`
	CorporateDestinationRestrictionAPI
	CorporateDestinationRestrictionCountry []CorporateDestinationRestrictionCountry `json:"corporate_destination_restriction_country,omitempty"`
	CorporateDestinationRestrictionCity    []CorporateDestinationRestrictionCity    `json:"corporate_destination_restriction_city,omitempty"`
	CorporateDestinationRestrictionAsset   *CorporateDestinationRestrictionAsset    `json:"corporate_destination_restriction_asset,omitempty"`
}

// CorporateDestinationRestrictionAPI model
type CorporateDestinationRestrictionAPI struct {
	IsAllowed *bool   `json:"is_allowed,omitempty" gorm:"type:bool;default:false" swaggertype:"boolean"`
	Condition *string `json:"condition,omitempty" gorm:"type:text" swaggertype:"string"`
}
