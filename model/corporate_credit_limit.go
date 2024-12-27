package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateCreditLimit Model
type CorporateCreditLimit struct {
	basic.Base
	basic.DataOwner
	CorporateCreditLimitAPI
	Corporate  *Corporate  `json:"corporate,omitempty"`   // Corporate
	CostCenter *CostCenter `json:"cost_center,omitempty"` // Cost Center
	Currency   *Currency   `json:"currency,omitempty"`    // Currency
	Project    *Project    `json:"project,omitempty"`     // Project
}

// CorporateCreditLimitAPI Model
type CorporateCreditLimitAPI struct {
	CorporateID                   *uuid.UUID       `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Corporate ID
	CostCenterID                  *uuid.UUID       `json:"cost_center_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`        // Cost Center ID
	CurrencyID                    *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`           // Currency ID
	CreditLimit                   *float64         `json:"credit_limit,omitempty" gorm:"type:decimal(19,4)"`                                           // Credit Limit
	IsPrimary                     *bool            `json:"is_primary,omitempty"`                                                                       // Is Primary
	IsTolerance                   *bool            `json:"is_tolerance,omitempty"`                                                                     // Is Tolerance
	ToleranceLevel                *int             `json:"tolerance_level,omitempty"`                                                                  // Tolerance Level
	UseCreditLimitFromCorporateID *uuid.UUID       `json:"use_credit_limit_from_corporate_id,omitempty"`                                               // Use Credit Limit From Corporate ID
	IsCreditLimitAllocated        *bool            `json:"is_credit_limit_allocated,omitempty"`                                                        // Is Credit Limit Allocated
	AskCreditLimitApproval        *bool            `json:"ask_credit_limit_approval,omitempty"`                                                        // Ask Credit Limit Approval
	ApprovalByEitherOne           *bool            `json:"approval_by_either_one,omitempty"`                                                           // Approval By Either One
	CreditUsed                    *float64         `json:"credit_used,omitempty" gorm:"type:decimal(19,4)"`                                            // Credit Used
	IsProject                     *bool            `json:"is_project,omitempty"`                                                                       // Is Project
	ProjectID                     *uuid.UUID       `json:"project_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`            // Project ID
	StartDate                     *strfmt.DateTime `json:"start_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`       // Start Date
	EndDate                       *strfmt.DateTime `json:"end_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`         // End Date
	IsTravelDate                  *bool            `json:"is_travel_date,omitempty"`                                                                   // Is Travel Date
	CreditStatus                  *int             `json:"credit_status,omitempty"`                                                                    // Credit Status
	CreditComment                 *string          `json:"credit_comment,omitempty" gorm:"varchar(4000)"`                                              // Credit Comment
}

// CorporateCreditLimitAPIAddition Model
type CorporateCreditLimitAddition struct {
	CorporateCreditLimit
	Status                          *Status    `json:"status,omitempty" gorm:"-"`
	CreditStatus                    *int       `json:"credit_status,omitempty"`
	TotalCreditUsed                 *float64   `json:"total_credit_used,omitempty"`
	CreditUsedRequested             *float64   `json:"credit_used_requested,omitempty"`
	OverCredit                      *float64   `json:"over_credit,omitempty"`
	CorporateCreditLimitToleranceID *uuid.UUID `json:"corporate_credit_limit_tolerance_id,omitempty"`
}
