package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// WidgetTypeTranslation Widget Type Translation
type WidgetTypeTranslation struct {
	basic.Base
	basic.DataOwner
	WidgetTypeTranslationAPI
	WidgetTypeID *uuid.UUID `json:"widget_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:widget_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Widget Type id
}

// WidgetTypeTranslationAPI Widget Type Translation API
type WidgetTypeTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:widget_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	WidgetTypeName *string `json:"widget_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Widget Type Name
	Description    *string `json:"description,omitempty" gorm:"type:text"`                                                                          // Description
}
