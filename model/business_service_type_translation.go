package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessServiceTypeTranslation Business Service Type Translation
type BusinessServiceTypeTranslation struct {
	basic.Base
	basic.DataOwner
	BusinessServiceTypeTranslationAPI
	BusinessServiceTypeID *uuid.UUID `json:"business_service_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:business_service_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Business Service Type id
}

// BusinessServiceTypeTranslationAPI Business Service Type Translation API
type BusinessServiceTypeTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:business_service_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	BusinessServiceName *string `json:"business_service_name,omitempty" gorm:"type:varchar(256)" example:"Maintenance"`                                            // Business Service Name
}
