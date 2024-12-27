package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CommissionTypeTranslation Commission Type Translation
type CommissionTypeTranslation struct {
	basic.Base
	basic.DataOwner
	CommissionTypeTranslationAPI
	CommissionTypeID *uuid.UUID `json:"commission_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:commission_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Commission Type id
}

// CommissionTypeTranslationAPI Commission Type Translation API
type CommissionTypeTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:commission_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CommissionTypeName *string `json:"commission_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Commission Type Name
	Description        *string `json:"description,omitempty" gorm:"type:text"`                                                                              // Description
}
