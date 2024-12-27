package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DivisionTranslation Division Translation
type DivisionTranslation struct {
	basic.Base
	basic.DataOwner
	DivisionTranslationAPI
	DivisionID *uuid.UUID `json:"division_id,omitempty" gorm:"type:varchar(36);uniqueIndex:division_translation_unique;not null" swaggertype:"string" format:"uuid"` // Division id
}

// DivisionTranslationAPI Division Translation API
type DivisionTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:division_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DivisionName *string `json:"division_name,omitempty" gorm:"type:varchar(256)"`                                                             // Division Name
}
