package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentTheme Model
type AgentTheme struct {
	basic.Base
	basic.DataOwner
	AgentThemeAPI
	Theme *Theme `json:"theme" gorm:"foreignKey:ThemeID;references:ID"`
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentThemeAPI API
type AgentThemeAPI struct {
	AgentID   *uuid.UUID       `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	ThemeID   *uuid.UUID       `json:"theme_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	StartDate *strfmt.DateTime `json:"start_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	EndDate   *strfmt.DateTime `json:"end_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	IsDefault *bool            `json:"is_default,omitempty"`
}
