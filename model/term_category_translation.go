package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TermCategoryTranslation Term Category Translation
type TermCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	TermCategoryTranslationAPI
	TermCategoryID *uuid.UUID `json:"term_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:term_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Term Category id
}

// TermCategoryTranslationAPI Term Category Translation API
type TermCategoryTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:term_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	TermCategoryName *string `json:"term_category_name,omitempty" gorm:"type:varchar(256)"`                                                             // Term Category Name
	Description      *string `json:"description,omitempty" gorm:"type:text"`                                                                            // Description
}
