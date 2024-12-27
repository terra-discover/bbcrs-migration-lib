package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TooltipTranslation Tooltip Translation
type TooltipTranslation struct {
	basic.Base
	basic.DataOwner
	TooltipTranslationAPI
	TooltipID *uuid.UUID `json:"tooltip_id,omitempty" gorm:"type:varchar(36);uniqueIndex:tooltip_translation_unique;not null" swaggertype:"string" format:"uuid"` // Tooltip id
}

// TooltipTranslationAPI Tooltip Translation API
type TooltipTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:tooltip_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	TooltipName  *string `json:"tooltip_name,omitempty" gorm:"type:varchar(64)"`
	Body         *string `json:"body,omitempty" gorm:"type:text" example:"body"`
	BodyFileName *string `json:"body_file_name,omitempty" gorm:"type:varchar(256)"`
	Description  *string `json:"description,omitempty" gorm:"type:text" example:"description"`
}
