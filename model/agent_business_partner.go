package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentBusinessPartner Model
type AgentBusinessPartner struct {
	basic.Base
	basic.DataOwner
	AgentBusinessPartnerAPI
	Agent         *Agent         `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	ProfileStatus *ProfileStatus `json:"profile_status" gorm:"foreignKey:ProfileStatusID;references:ID"`
	ApprovedBy    *UserAccount   `json:"approved_by" gorm:"foreignKey:ApprovedByID;references:ID"`
	RegisteredBy  *UserAccount   `json:"registered_by" gorm:"foreignKey:RegisteredByID;references:ID"`
	RejectedBy    *UserAccount   `json:"rejected_by" gorm:"foreignKey:RejectedByID;references:ID"`
	TerminatedBy  *UserAccount   `json:"terminated_by" gorm:"foreignKey:TerminatedByID;references:ID"`
}

// AgentBusinessPartnerAPI API
type AgentBusinessPartnerAPI struct {
	AgentID            *uuid.UUID       `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	BusinessPartnerID  *uuid.UUID       `json:"business_partner_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	ProfileStatusID    *uuid.UUID       `json:"profile_status_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	EffectiveDate      *strfmt.DateTime `json:"effective_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	ExpireDate         *strfmt.DateTime `json:"expire_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	RegisteredByID     *uuid.UUID       `json:"registered_by_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	RegisteredAt       *strfmt.DateTime `json:"registered_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	RegistrationReason *string          `json:"registration_reason,omitempty" gorm:"type:text"`
	ApprovedByID       *uuid.UUID       `json:"approved_by_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	ApprovedAt         *strfmt.DateTime `json:"approved_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	ApprovalReason     *string          `json:"approval_reason,omitempty" gorm:"type:text"`
	RejectedByID       *uuid.UUID       `json:"rejected_by_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	RejectedAt         *strfmt.DateTime `json:"rejected_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	RejectionReason    *string          `json:"rejection_reason,omitempty" gorm:"type:text"`
	TerminatedByID     *uuid.UUID       `json:"terminated_by_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	TerminatedAt       *strfmt.DateTime `json:"terminated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	TerminationReason  *string          `json:"termination_reason,omitempty" gorm:"type:text"`
	Comment            *string          `json:"comment,omitempty" gorm:"type:text"`
	Description        *string          `json:"description,omitempty" gorm:"type:text"`
}
