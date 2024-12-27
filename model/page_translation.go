package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PageTranslation Page Translation
type PageTranslation struct {
	basic.Base
	basic.DataOwner
	PageTranslationAPI
	PageID *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36);uniqueIndex:page_translation_unique;not null" swaggertype:"string" format:"uuid"` // Page id
}

// PageTranslationAPI Page Translation API
type PageTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:page_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PageName     *string `json:"page_name,omitempty" gorm:"type:varchar(256)" example:"Contact Us"`                                        // Page Name
	Slug         *string `json:"slug,omitempty" gorm:"type:varchar(256)" example:"contact-us"`                                             // Slug
	HTMLTitle    *string `json:"html_title,omitempty" gorm:"type:varchar(256)"`                                                            // Html Title
	Description  *string `json:"description,omitempty" gorm:"type:varchar(4000)"`                                                          // Description
}
