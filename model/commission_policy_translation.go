package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CommissionPolicyTranslation Commission Policy Translation
type CommissionPolicyTranslation struct {
	basic.Base
	basic.DataOwner
	CommissionPolicyTranslationAPI
	CommissionPolicyID *uuid.UUID `json:"commission_policy_id,omitempty" gorm:"type:varchar(36);uniqueIndex:commission_policy_translation_unique;not null" swaggertype:"string" format:"uuid"` // Commission Policy id
}

// CommissionPolicyTranslationAPI Commission Policy Translation API
type CommissionPolicyTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:commission_policy_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CommissionPolicyName *string `json:"commission_policy_name,omitempty" gorm:"type:varchar(256)"`                                                             // Commission Policy Name
	Description          *string `json:"description,omitempty" gorm:"type:text"`                                                                                // Description
}
