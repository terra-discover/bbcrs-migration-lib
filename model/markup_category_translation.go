package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupCategoryTranslation Markup Category Translation
type MarkupCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	MarkupCategoryTranslationAPI
	MarkupCategoryID *uuid.UUID `json:"markup_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:markup_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Markup Category id
}

// MarkupCategoryTranslationAPI Mark Up Category Translation API
type MarkupCategoryTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:markup_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MarkupCategoryName *string `json:"markup_category_name,omitempty" gorm:"type:varchar(256)" example:"Agent Menu"`                                        // Menu Link Name
	URL                *string `json:"url,omitempty" gorm:"type:varchar(256)" example:"example.com"`                                                        // Menu URL
	Description        *string `json:"description,omitempty" gorm:"type:text"`                                                                              // Description
}
