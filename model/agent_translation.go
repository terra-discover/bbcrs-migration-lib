package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentTranslation Agent Translation
type AgentTranslation struct {
	basic.Base
	basic.DataOwner
	AgentTranslationAPI
	AgentID *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);uniqueIndex:agent_translation_unique;not null" swaggertype:"string" format:"uuid"`
}

// AgentTranslationAPI Agent Translation API
type AgentTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:agent_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AgentName    *string `json:"agent_name,omitempty" gorm:"type:varchar(256)"`
	Comment      *string `json:"comment,omitempty" gorm:"type:text"`
	Description  *string `json:"description,omitempty" gorm:"type:text"`
}
