package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ChargeTypeTranslation Charge Type Translation
type ChargeTypeTranslation struct {
	basic.Base
	basic.DataOwner
	ChargeTypeTranslationAPI
	ChargeTypeID *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:charge_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Charge Type id
}

// ChargeTypeTranslationAPI Charge Type Translation API
type ChargeTypeTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:charge_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ChargeTypeName *string `json:"charge_type_name,omitempty" gorm:"type:varchar(256)" example:"Daily"`                                             // Charge Type Name
}
