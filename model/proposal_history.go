package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Proposal History
type ProposalHistory struct {
	basic.Base
	basic.DataOwner
	ProposalHistoryAPI
}

// ProposalHistoryAPI Proposal Hostory API
type ProposalHistoryAPI struct {
	ProposalID           *uuid.UUID       `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProposalCode         *string          `json:"proposal_code,omitempty" gorm:"type:varchar(36);not null"`
	ProposalStatus       *int             `json:"proposal_status,omitempty" gorm:"type:smallint"`
	UserAccountID        *uuid.UUID       `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	CreditStatus         *int             `json:"credit_status,omitempty" gorm:"type:smallint"`
	CreditComment        *string          `json:"credit_comment,omitempty" gorm:"type:text"`
	TravelRequestStatus  *int             `json:"travel_request_status,omitempty" gorm:"type:smallint"`
	TravelRequestComment *string          `json:"travel_request_comment,omitempty" gorm:"type:text"`
	ModificationStatus   *int             `json:"modification_status,omitempty" gorm:"type:smallint"`
	ModificationComment  *string          `json:"modification_comment,omitempty" gorm:"type:text"`
	CancellationStatus   *int             `json:"cancellation_status,omitempty" gorm:"type:smallint"`
	CancellationComment  *string          `json:"cancellation_comment,omitempty" gorm:"type:text"`
	Description          *string          `json:"description,omitempty" gorm:"type:text"`
	Timestamp            *strfmt.DateTime `json:"timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	LocalTimestamp       *strfmt.DateTime `json:"local_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
}
