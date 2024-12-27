package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeCategoryTranslation Service Fee Category Translation
type ServiceFeeCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	ServiceFeeCategoryTranslationAPI
	ServiceFeeCategoryID *uuid.UUID `json:"service_fee_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:service_fee_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Service Fee Category id
}

// ServiceFeeCategoryTranslationAPI Service Fee Category Translation API
type ServiceFeeCategoryTranslationAPI struct {
	LanguageCode           *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:service_fee_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ServiceFeeCategoryName *string `json:"service_fee_category_name,omitempty" gorm:"type:varchar(256)"`                                                             // Service Fee Category Name
	Description            *string `json:"description,omitempty" gorm:"type:text"`                                                                                   // Description
}
