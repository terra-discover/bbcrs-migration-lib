package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardTypeTranslation Reward Type Translation
type RewardTypeTranslation struct {
	basic.Base
	basic.DataOwner
	RewardTypeTranslationAPI
	RewardTypeID *uuid.UUID `json:"reward_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:reward_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Reward Type id
}

// RewardTypeTranslationAPI Reward Type Translation API
type RewardTypeTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:reward_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RewardTypeName *string `json:"reward_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Reward Type Name
}
