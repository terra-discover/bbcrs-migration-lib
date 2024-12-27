package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CostCenterTranslation Cost Center Translation
type CostCenterTranslation struct {
	basic.Base
	basic.DataOwner
	CostCenterTranslationAPI
	CostCenterID *uuid.UUID `json:"cost_center_id,omitempty" gorm:"type:varchar(36);uniqueIndex:cost_center_translation_unique;not null" swaggertype:"string" format:"uuid"` // Cost Center id
}

// CostCenterTranslationAPI Cost Center Translation API
type CostCenterTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:cost_center_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CostCenterName *string `json:"cost_center_name,omitempty" gorm:"type:varchar(256)"`                                                             // Cost Center Name
}
