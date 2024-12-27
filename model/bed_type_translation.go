package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BedTypeTranslation Bed Type Translation
type BedTypeTranslation struct {
	basic.Base
	basic.DataOwner
	BedTypeTranslationAPI
	BedTypeID *uuid.UUID `json:"bed_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:bed_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Bed Type id
}

// BedTypeTranslationAPI Bed Type Translation API
type BedTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:bed_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	BedTypeName  *string `json:"bed_type_name,omitempty" gorm:"type:varchar(256)" example:"queen bed"`                                         // Bed Type Name
}
