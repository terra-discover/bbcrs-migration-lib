package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// AgentCorporate Agent Corporate
type AgentCorporate struct {
	basic.Base
	basic.DataOwner
	AgentCorporateAPI
	Agent     *Agent     `json:"agent,omitempty"`
	Corporate *Corporate `json:"corporate,omitempty"`
	Division  *Division  `json:"division,omitempty"`
	// ProfileStatus *ProfileStatus `json:"profile_status,omitempty" gorm:"foreignKey:ProfileStatusCode;references:profile_status_code"`
}

// AgentCorporateAPI Agent Corporate Api
type AgentCorporateAPI struct {
	AgentID            *uuid.UUID       `json:"agent_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`                                            // Agent Id
	CorporateID        *uuid.UUID       `json:"corporate_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Corporate Id
	ProfileStatus      *int             `json:"profile_status,omitempty" gorm:"type:int"`                                                                                          // Profile Status
	EffectiveDate      *strfmt.DateTime `json:"effective_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`                                          // Effective Date
	ExpireDate         *strfmt.DateTime `json:"expire_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`                                             // Expire Date
	RegisteredByID     *uuid.UUID       `json:"registered_by_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                             // Registered By Id
	RegisteredAt       *strfmt.DateTime `json:"registered_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`                                           // Registered At
	RegistrationReason *string          `json:"registration_reason,omitempty" gorm:"type:varchar(4000)"`                                                                           // Registration Reason
	ApprovedByID       *uuid.UUID       `json:"approved_by_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                               // Approved By Id
	ApprovedAt         *strfmt.DateTime `json:"approved_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`                                             // Approved At
	ApprovalReason     *string          `json:"approval_reason,omitempty" gorm:"type:varchar(4000)"`                                                                               // Approval Reason
	RejectedByID       *uuid.UUID       `json:"rejected_by_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                               // Rejected By Id
	RejectedAt         *strfmt.DateTime `json:"rejected_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`                                             // Rejected At
	RejectionReason    *string          `json:"approved_reason,omitempty" gorm:"type:varchar(4000)"`                                                                               // Rejection Reason
	TerminatedByID     *uuid.UUID       `json:"terminated_by_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                             // Terminated By Id
	TerminatedAt       *strfmt.DateTime `json:"terminated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`                                           // Terminated At
	TerminationReason  *string          `json:"termination_reason,omitempty" gorm:"type:varchar(4000)"`                                                                            // Termination Reason
	DivisionID         *uuid.UUID       `json:"division_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                  // Division Id
	Comment            *string          `json:"comment,omitempty" gorm:"type:text"`                                                                                                // Comment
	Description        *string          `json:"description,omitempty" gorm:"type:text"`                                                                                            // Description
}

// Seed agent corporate
func (a *AgentCorporate) Seed(agentID uuid.UUID) *[]AgentCorporate {
	data := []AgentCorporate{}
	agentCorporate := AgentCorporate{}
	agentCorporate.AgentID = &agentID
	agentCorporate.CorporateID = &agentID
	agentCorporate.ProfileStatus = lib.Intptr(1)

	data = append(data, agentCorporate)

	return &data
}

type AgentCorporateListData struct {
	Corporate
	ParentCorporateCode    *string  `json:"parent_corporate_code,omitempty"`
	CorporateRating        *float64 `json:"corporate_rating,omitempty"`
	TeamConsultantNames    *string  `json:"team_consultant_names,omitempty"`
	ConsultantNames        *string  `json:"consultant_names,omitempty"`
	CorporateGeneralStatus *string  `json:"corporate_general_status,omitempty"`
	ProfileStatusCode      *int     `json:"profile_status_code"` // For knowing agent corporate status is active (1), draft (2), or blocked (3)
}
