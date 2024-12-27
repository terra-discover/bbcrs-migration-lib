package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentSite AgentSite
type AgentSite struct {
	basic.Base
	basic.DataOwner
	AgentSiteAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	Site  *Site  `json:"site" gorm:"foreignKey:SiteID;references:ID"`
}

// AgentSiteAPI AgentSite API
type AgentSiteAPI struct {
	AgentID            *uuid.UUID       `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	SiteID             *uuid.UUID       `json:"site_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	ProfileStatusID    *uuid.UUID       `json:"profile_status_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	EffectiveDate      *strfmt.DateTime `json:"effective_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	ExpireDate         *strfmt.DateTime `json:"expire_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	RegisteredByID     *uuid.UUID       `json:"registered_by_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	RegisteredAt       *strfmt.DateTime `json:"registered_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	RegistrationReason *string          `json:"registration_reason,omitempty" gorm:"type:text"`
	ApprovedByID       *uuid.UUID       `json:"approved_by_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	ApprovedAt         *strfmt.DateTime `json:"approved_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	ApprovalReason     *string          `json:"approval_reason,omitempty" gorm:"type:text"`
	RejectedByID       *uuid.UUID       `json:"rejected_by_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	RejectedAt         *strfmt.DateTime `json:"rejected_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	RejectionReason    *string          `json:"rejection_reason,omitempty" gorm:"type:text"`
	TerminatedByID     *uuid.UUID       `json:"terminated_by_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	TerminatedAt       *strfmt.DateTime `json:"terminated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	TerminationReason  *string          `json:"termination_reason,omitempty" gorm:"type:text"`
	Comment            *string          `json:"comment,omitempty" gorm:"type:text"`
	Description        *string          `json:"description,omitempty" gorm:"type:text"`
}
