package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyLevelTranslation Loyalty Level Translation
type LoyaltyLevelTranslation struct {
	basic.Base
	basic.DataOwner
	LoyaltyLevelTranslationAPI
	LoyaltyLevelID *uuid.UUID `json:"loyalty_level_id,omitempty" gorm:"type:varchar(36);uniqueIndex:loyalty_level_translation_unique;not null" swaggertype:"string" format:"uuid"` // Loyalty Level id
}

// LoyaltyLevelTranslationAPI Loyalty Level Translation API
type LoyaltyLevelTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:loyalty_level_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	LoyaltyLevelName *string `json:"loyalty_level_name,omitempty" gorm:"type:varchar(256)"`                                                             // Loyalty Level Name
}
