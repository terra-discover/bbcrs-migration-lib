package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateCreditLimitApprover struct
type CorporateCreditLimitApprover struct {
	basic.Base
	basic.DataOwner
	CorporateCreditLimitApproverAPI
	Employee             *Employee             `json:"employee,omitempty"`
	CorporateCreditLimit *CorporateCreditLimit `json:"corporate_credit_limit,omitempty"`
}

// CorporateCreditLimitApproverAPI for payload
type CorporateCreditLimitApproverAPI struct {
	EmployeeID             *uuid.UUID `json:"employee_id" gorm:"type:varchar(36);not null;index:idx_corporate_credit_limit_approver__employee_id,unique,where:deleted_at is null"`               // Corporate ID
	CorporateCreditLimitID *uuid.UUID `json:"corporate_credit_limit_id" gorm:"type:varchar(36);not null;index:idx_corporate_credit_limit_approver__employee_id,unique,where:deleted_at is null"` // Corporate Credit Limit ID
}
