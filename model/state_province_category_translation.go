package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// StateProvinceCategoryTranslation State Province Category Translation
type StateProvinceCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	StateProvinceCategoryTranslationAPI
	StateProvinceCategoryID *uuid.UUID `json:"state_province_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:state_province_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // State Province Category id
}

// StateProvinceCategoryTranslationAPI State Province Category Translation API
type StateProvinceCategoryTranslationAPI struct {
	LanguageCode              *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:state_province_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	StateProvinceCategoryName *string `json:"state_province_category_name,omitempty" gorm:"type:varchar(256)" example:"Jogjakarta"`                                        // State Province Category Name
}
