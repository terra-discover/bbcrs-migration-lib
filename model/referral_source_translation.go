package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReferralSourceTranslation Referral Source Translation
type ReferralSourceTranslation struct {
	basic.Base
	basic.DataOwner
	ReferralSourceTranslationAPI
	ReferralSourceID *uuid.UUID `json:"referral_source_id,omitempty" gorm:"type:varchar(36);uniqueIndex:referral_source_translation_unique;not null" swaggertype:"string" format:"uuid"` // Referral Source id
}

// ReferralSourceTranslationAPI Referral Source Translation API
type ReferralSourceTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:referral_source_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ReferralSourceName *string `json:"referral_source_name,omitempty" gorm:"type:varchar(64)" example:"Referral Source Name"`                               // Referral Source Name
}
