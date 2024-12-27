package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MenuTranslation Menu Translation
type MenuTranslation struct {
	basic.Base
	basic.DataOwner
	MenuTranslationAPI
	MenuID *uuid.UUID `json:"menu_id,omitempty" gorm:"type:varchar(36);uniqueIndex:menu_translation_unique;not null" swaggertype:"string" format:"uuid"` // Menu id
}

// MenuTranslationAPI Menu Translation API
type MenuTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:menu_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MenuName     *string `json:"menu_name,omitempty" gorm:"type:varchar(256)" example:"Agent Menu"`                                        // Menu Name
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                   // Description
}
