package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BenefitTranslation Benefit Translation
type BenefitTranslation struct {
	basic.Base
	basic.DataOwner
	BenefitTranslationAPI
	BenefitID *uuid.UUID `json:"benefit_id,omitempty" gorm:"type:varchar(36);uniqueIndex:benefit_translation_unique;not null" swaggertype:"string" format:"uuid"` // Benefit id
}

// BenefitTranslationAPI Benefit Translation API
type BenefitTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:benefit_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	BenefitName  *string `json:"benefit_name,omitempty" gorm:"type:varchar(256)" example:"free parking"`                                      // Benefit Name
	Description  *string `json:"description,omitempty" gorm:"type:text" example:"free parking for all passengers"`                            // Description
}
