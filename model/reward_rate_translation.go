package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardRateTranslation Reward Type Translation
type RewardRateTranslation struct {
	basic.Base
	basic.DataOwner
	RewardRateTranslationAPI
	RewardRateID *uuid.UUID `json:"reward_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:reward_rate_translation_unique;not null" swaggertype:"string" format:"uuid"` // Reward Type id
}

// RewardRateTranslationAPI Reward Type Translation API
type RewardRateTranslationAPI struct {
	LanguageID  *uuid.UUID `json:"language_code,omitempty" gorm:"type:varchar(36);uniqueIndex:reward_rate_translation_unique;not null" example:"en"` // Language Id
	Description *string    `json:"description,omitempty" gorm:"type:varchar(4000)"`                                                                  // Description
}
