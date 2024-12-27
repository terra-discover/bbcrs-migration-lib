package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PropertyCategoryTranslation Property Category Translation
type PropertyCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	PropertyCategoryTranslationAPI
	PropertyCategoryID *uuid.UUID `json:"property_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:property_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Property Category id
}

// PropertyCategoryTranslationAPI Property Category Translation API
type PropertyCategoryTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:property_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PropertyCategoryName *string `json:"property_category_name,omitempty" gorm:"type:varchar(256)"`                                                             // Property Category Name
}
