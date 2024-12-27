package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentLanguage Model
type AgentLanguage struct {
	basic.Base
	basic.DataOwner
	AgentLanguageAPI
	Language *Language  `json:"language,omitempty"`                                                                                                              // language detail
	Agent    *Agent     `json:"agent,omitempty" swaggerignore:"true"`                                                                                            // agent detail
	AgentID  *uuid.UUID `json:"-" gorm:"type:varchar(36);index:idx_agent_languauge__language_id,unique,where:deleted_at is null;not null;" swaggerignore:"true"` // agent id
}

// AgentLanguageAPI API
type AgentLanguageAPI struct {
	LanguageID *uuid.UUID `json:"language_id,omitempty" gorm:"type:varchar(36);index:idx_agent_languauge__language_id,unique,where:deleted_at is null;not null;" format:"uuid"` // language idw
	IsPrimary  *bool      `json:"is_primary,omitempty"`                                                                                                                         // is primary
}

// AgentLanguageDataAPI API
type AgentLanguageDataAPI struct {
	PrimaryLanguage struct {
		ID *uuid.UUID `json:"id,omitempty" swaggertype:"string" format:"uuid" validate:"required"` // Language ID
	} `json:"primary_language,omitempty"` // primary language
	AdditionalLanguages []struct {
		ID *uuid.UUID `json:"id,omitempty" swaggertype:"string" format:"uuid" validate:"required"` // Language ID
	} `json:"additional_languages,omitempty"` // additional language
}

// AgentLanguageData Model
type AgentLanguageData struct {
	PrimaryLanguage     *Language   `json:"primary_language"`     // primary language
	AdditionalLanguages []*Language `json:"additional_languages"` // additional languages
}
