package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// WidgetTranslation Widget  Translation
type WidgetTranslation struct {
	basic.Base
	basic.DataOwner
	WidgetTranslationAPI
	WidgetID *uuid.UUID `json:"widget_id,omitempty" gorm:"type:varchar(36);uniqueIndex:widget_translation_unique;not null" swaggertype:"string" format:"uuid"` // Widget  id
}

// WidgetTranslationAPI Widget  Translation API
type WidgetTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:widget_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	WidgetName   *string `json:"widget_name,omitempty" gorm:"type:varchar(256)"`                                                             // Widget  Name
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                     // description
}
