package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// OfferType Offer Type
type OfferType struct {
	basic.Base
	basic.DataOwner
	OfferTypeAPI
	OfferTypeTranslation *OfferTypeTranslation `json:"offer_type_translation,omitempty"` // Offer Type Translation
}

// OfferTypeAPI Offer Type API
type OfferTypeAPI struct {
	OfferTypeCode *int    `json:"offer_type_code,omitempty" gorm:"type:int;index:,unique,where:deleted_at is null;not null"`          // Offer Type Code
	OfferTypeName *string `json:"offer_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Offer Type Name
}
