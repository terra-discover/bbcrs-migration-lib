package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentWidgetType Model
type AgentWidgetType struct {
	basic.Base
	basic.DataOwner
	AgentWidgetTypeAPI
	WidgetType *WidgetType `json:"widget_type" gorm:"foreignKey:WidgetTypeID;references:ID"`
	Agent      *Agent      `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentWidgetTypeAPI API
type AgentWidgetTypeAPI struct {
	AgentID      *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	WidgetTypeID *uuid.UUID `json:"widget_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
