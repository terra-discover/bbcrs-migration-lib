package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeCategoryTranslation Processing Fee Category Translation
type ProcessingFeeCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeCategoryTranslationAPI
	ProcessingFeeCategoryID *uuid.UUID `json:"processing_fee_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:processing_fee_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Processing Fee Category id
}

// ProcessingFeeCategoryTranslationAPI Processing Fee Category Translation API
type ProcessingFeeCategoryTranslationAPI struct {
	LanguageCode              *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:processing_fee_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ProcessingFeeCategoryName *string `json:"processing_fee_category_name,omitempty" example:"Flight" gorm:"type:varchar(256)"`                                            // Processing Fee Category Name
	Description               *string `json:"description,omitempty" gorm:"type:text"`                                                                                      // Description
}
