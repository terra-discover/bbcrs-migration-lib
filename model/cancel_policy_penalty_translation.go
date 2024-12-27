package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CancelPolicyPenaltyTranslation Cancel Policy Penalty Translation
type CancelPolicyPenaltyTranslation struct {
	basic.Base
	basic.DataOwner
	CancelPolicyPenaltyTranslationAPI
	CancelPolicyPenaltyID *uuid.UUID `json:"cancel_policy_penalty_id,omitempty" gorm:"type:varchar(36);uniqueIndex:cancel_policy_penalty_translation_unique;not null" swaggertype:"string" format:"uuid"` // Cancel Policy Penalty id
}

// CancelPolicyPenaltyTranslationAPI Cancel Policy Penalty Translation API
type CancelPolicyPenaltyTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:cancel_policy_penalty_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                                    // Description
}
