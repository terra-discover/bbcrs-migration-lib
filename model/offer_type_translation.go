package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OfferTypeTranslation Offer Type Translation
type OfferTypeTranslation struct {
	basic.Base
	basic.DataOwner
	OfferTypeTranslationAPI
	OfferTypeID *uuid.UUID `json:"offer_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:offer_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Offer Type id
}

// OfferTypeTranslationAPI Offer Type Translation API
type OfferTypeTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:offer_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	OfferTypeName *string `json:"offer_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Offer Type Name
}
