package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CancelPolicyTranslation Cancel Policy Translation
type CancelPolicyTranslation struct {
	basic.Base
	basic.DataOwner
	CancelPolicyTranslationAPI
	CancelPolicyID *uuid.UUID `json:"cancel_policy_id,omitempty" gorm:"type:varchar(36);uniqueIndex:cancel_policy_translation_unique;not null" swaggertype:"string" format:"uuid"` // Cancel Policy id
}

// CancelPolicyTranslationAPI Cancel Policy Translation API
type CancelPolicyTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:cancel_policy_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CancelPolicyName *string `json:"cancel_policy_name,omitempty" gorm:"type:varchar(256)"`                                                             // Cancel Policy Name
	Description      *string `json:"description,omitempty" gorm:"type:text"`                                                                            // Description
}
