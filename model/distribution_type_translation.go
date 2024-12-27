package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DistributionTypeTranslation Distribution Type Translation
type DistributionTypeTranslation struct {
	basic.Base
	basic.DataOwner
	DistributionTypeTranslationAPI
	DistributionTypeID *uuid.UUID `json:"distribution_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:distribution_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Distribution Type id
}

// DistributionTypeTranslationAPI Distribution Type Translation API
type DistributionTypeTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:distribution_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DistributionTypeName *string `json:"distribution_type_name,omitempty" gorm:"type:varchar(256)" example:"Fax"`                                               // Distribution Type Name
}
