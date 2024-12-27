package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ModifyPolicyTranslation Modify Policy Translation
type ModifyPolicyTranslation struct {
	basic.Base
	basic.DataOwner
	ModifyPolicyTranslationAPI
	ModifyPolicyID *uuid.UUID `json:"modify_policy_id,omitempty" gorm:"type:varchar(36);uniqueIndex:modify_policy_translation_unique;not null" swaggertype:"string" format:"uuid"` // Modify Policy id
}

// ModifyPolicyTranslationAPI Modify Policy Translation API
type ModifyPolicyTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:modify_policy_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ModifyPolicyName *string `json:"modify_policy_name,omitempty" gorm:"type:varchar(256)"`                                                             // Modify Policy Name
	Description      *string `json:"description,omitempty" gorm:"type:text"`                                                                            // Description
}
