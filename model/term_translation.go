package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TermTranslation Term Translation
type TermTranslation struct {
	basic.Base
	basic.DataOwner
	TermTranslationAPI
	TermID *uuid.UUID `json:"term_id,omitempty" gorm:"type:varchar(36);uniqueIndex:term_translation_unique;not null" swaggertype:"string" format:"uuid"` // Term id
}

// TermTranslationAPI Term Translation API
type TermTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:term_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	TermName     *string `json:"term_name,omitempty" gorm:"type:varchar(256)"`                                                             // Term Name
	Slug         *string `json:"slug,omitempty" gorm:"type:varchar(256)"`                                                                  // Slug
	HTMLTitle    *string `json:"html_title,omitempty" gorm:"type:varchar(256)"`                                                            // HTML Title
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                   // Description
}
