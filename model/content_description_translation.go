package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ContentDescriptionTranslation Content Description Translation
type ContentDescriptionTranslation struct {
	basic.Base
	basic.DataOwner
	ContentDescriptionTranslationAPI
	ContentDescriptionID *uuid.UUID `json:"content_description_id,omitempty" gorm:"type:varchar(36);uniqueIndex:content_description_translation_unique;not null" swaggertype:"string" format:"uuid"` // Content Description id
}

// ContentDescriptionTranslationAPI Content Description Translation API
type ContentDescriptionTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:content_description_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	Title        *string `json:"title,omitempty" gorm:"type:varchar(256)"`                                                                                // Title
	URL          *string `json:"url,omitempty" gorm:"type:varchar(256)"`                                                                                  // Url
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                                  // Description
}
