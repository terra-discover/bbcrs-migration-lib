package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UniqueIDTypeTranslation Unique ID Type Translation
type UniqueIDTypeTranslation struct {
	basic.Base
	basic.DataOwner
	UniqueIDTypeTranslationAPI
	UniqueIDTypeID *uuid.UUID `json:"unique_id_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:unique_id_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Unique ID Type id
}

// UniqueIDTypeTranslationAPI Unique ID Type Translation API
type UniqueIDTypeTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:unique_id_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	UniqueIDTypeName *string `json:"unique_id_type_name,omitempty" gorm:"type:varchar(256)" example:"Customer"`                                          // Unique ID Type Name
}
