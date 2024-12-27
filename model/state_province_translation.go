package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// StateProvinceTranslation State Province Translation
type StateProvinceTranslation struct {
	basic.Base
	basic.DataOwner
	StateProvinceTranslationAPI
	StateProvinceID *uuid.UUID `json:"state_province_id,omitempty" gorm:"type:varchar(36);uniqueIndex:state_province_translation_unique;not null" swaggertype:"string" format:"uuid"` // State Province id
}

// StateProvinceTranslationAPI State Province Translation API
type StateProvinceTranslationAPI struct {
	StateProvinceID   *uuid.UUID `json:"state_province_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                             // State Province Category ID
	LanguageCode      *string    `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:state_province_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	StateProvinceName *string    `json:"state_province_name,omitempty" gorm:"type:varchar(256)" example:"Jogjakarta"`                                        // State Province Name
}
