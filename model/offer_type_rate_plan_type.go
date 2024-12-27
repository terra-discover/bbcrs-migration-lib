package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OfferTypeRatePlanType Offer Type Rate Plan Type
type OfferTypeRatePlanType struct {
	basic.Base
	basic.DataOwner
	OfferTypeRatePlanTypeAPI
}

// OfferTypeRatePlanTypeAPI Offer Type Rate Plan Type Api
type OfferTypeRatePlanTypeAPI struct {
	OfferTypeID    *uuid.UUID `json:"offer_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`     // Offer Type Id
	RatePlanTypeID *uuid.UUID `json:"rate_plan_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Rate Plan Type Id
}
