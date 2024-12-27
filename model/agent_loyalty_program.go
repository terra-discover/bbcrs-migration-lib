package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentLoyaltyProgram Agent Loyalty Program
type AgentLoyaltyProgram struct {
	basic.Base
	basic.DataOwner
	AgentLoyaltyProgramAPI
	Agent          *Agent          `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	LoyaltyProgram *LoyaltyProgram `json:"loyalty_program" gorm:"foreignKey:LoyaltyProgramID;references:ID"`
	ProfileStatus  *ProfileStatus  `json:"profile_status" gorm:"foreignKey:ProfileStatusID;references:ID"`
	ApprovedBy     *UserAccount    `json:"approved_by" gorm:"foreignKey:ApprovedByID;references:ID"`
	RegisteredBy   *UserAccount    `json:"registered_by" gorm:"foreignKey:RegisteredByID;references:ID"`
	RejectedBy     *UserAccount    `json:"rejected_by" gorm:"foreignKey:RejectedByID;references:ID"`
	TerminatedBy   *UserAccount    `json:"terminated_by" gorm:"foreignKey:TerminatedByID;references:ID"`
}

// AgentLoyaltyProgramAPI Agent Loyalty Program Api
type AgentLoyaltyProgramAPI struct {
	AgentID            *uuid.UUID       `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	LoyaltyProgramID   *uuid.UUID       `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ProfileStatusID    *uuid.UUID       `json:"profile_status_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	EffectiveDate      *strfmt.DateTime `json:"effective_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	ExpiredDate        *strfmt.DateTime `json:"expired_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
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
	IsPrimary          *bool            `json:"is_primary,omitempty"`
}
