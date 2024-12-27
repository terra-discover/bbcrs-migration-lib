package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentSpecialDateTranslation model
type AgentSpecialDateTranslation struct {
	basic.Base
	basic.DataOwner
	AgentSpecialDateTranslationAPI
	AgentSpecialDateID *uuid.UUID `json:"agent_special_date_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // Special Date ID
}

// AgentSpecialDateTranslationAPI struct
type AgentSpecialDateTranslationAPI struct {
	SpecialDateName *string `json:"special_date_name,omitempty" gorm:"type:varchar(256);uniqueIndex:,where:deleted_at is null;not null"` // Special Date Name
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);not null"`                                             // Language Code
}
