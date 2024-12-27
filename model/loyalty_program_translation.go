package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramTranslation Loyalty Program Translation
type LoyaltyProgramTranslation struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramTranslationAPI
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);uniqueIndex:loyalty_program_translation_unique;not null" swaggertype:"string" format:"uuid"` // Loyalty Program id
}

// LoyaltyProgramTranslationAPI Loyalty Program Translation API
type LoyaltyProgramTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:loyalty_program_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	LoyaltyProgramName *string `json:"loyalty_program_name,omitempty" gorm:"type:varchar(256)"`                                                             // Loyalty Program Name
	Comment            *string `json:"comment,omitempty" gorm:"type:varchar(4000)"`                                                                         // Comment
	Description        *string `json:"description,omitempty" gorm:"type:varchar(4000)"`                                                                     // Description
}
