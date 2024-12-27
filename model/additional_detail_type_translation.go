package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdditionalDetailTypeTranslation Additional Detail Type Translation
type AdditionalDetailTypeTranslation struct {
	basic.Base
	basic.DataOwner
	AdditionalDetailTypeTranslationAPI
	AdditionalDetailTypeID *uuid.UUID `json:"additional_detail_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:additional_detail_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Additional Detail Type id
}

// AdditionalDetailTypeTranslationAPI Additional Detail Type Translation API
type AdditionalDetailTypeTranslationAPI struct {
	LanguageCode             *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:additional_detail_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AdditionalDetailTypeName *string `json:"additional_detail_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Additional Detail Type Name
}
