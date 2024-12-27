package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPurposeTranslation model
type TravelPurposeTranslation struct {
	basic.Base
	basic.DataOwner
	TravelPurposeID *uuid.UUID `json:"travel_purpose_id,omitempty" gorm:"type:varchar(36);uniqueIndex:travel_purpose_translation_unique;not null" swaggertype:"string" format:"uuid"` // travel purpose id
	TravelPurposeTranslationAPI
}

// TravelPurposeTranslationAPI model
type TravelPurposeTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:travel_purpose_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	TravelPurposeName *string `json:"travel_purpose_name,omitempty" gorm:"type:varchar(256)"`
}
