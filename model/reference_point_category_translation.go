package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReferencePointCategoryTranslation Reference Point Category Translation
type ReferencePointCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	ReferencePointCategoryTranslationAPI
	ReferencePointCategoryID *uuid.UUID `json:"reference_point_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:reference_point_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Reference Point Category id
}

// ReferencePointCategoryTranslationAPI Reference Point Category Translation API
type ReferencePointCategoryTranslationAPI struct {
	LanguageCode               *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:reference_point_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ReferencePointCategoryName *string `json:"reference_point_category_name,omitempty" gorm:"type:varchar(256)" example:"Airport"`                                           // Reference Point Category Name
}
