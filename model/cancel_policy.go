package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CancelPolicy Cancel Policy
type CancelPolicy struct {
	basic.Base
	basic.DataOwner
	CancelPolicyAPI
	CancelPolicyTranslation *CancelPolicyTranslation `json:"cancel_policy_translation,omitempty"` // Cancel Policy Translation
	OffsetTimeUnit          *OffsetTimeUnit          `json:"offset_time_unit,omitempty"`          // Offset Time Unit
	BasisType               *BasisType               `json:"basis_type,omitempty"`                // Basis Type
}

// CancelPolicyAPI Cancel Policy API
type CancelPolicyAPI struct {
	CancelPolicyCode *string    `json:"cancel_policy_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`  // Cancel Policy Code
	CancelPolicyName *string    `json:"cancel_policy_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Cancel Policy Name
	OffsetTimeUnitID *uuid.UUID `json:"offset_time_unit_id,omitempty" swaggertype:"string" format:"uuid"`                                      // Offset Time Unit ID
	BasisTypeID      *uuid.UUID `json:"basis_type_id,omitempty" swaggertype:"string" format:"uuid"`                                            // Basis Type ID
	IsTaxInclusive   *bool      `json:"is_tax_inclusive,omitempty"`                                                                            // Is Tax Inclusive
	IsFeeInclusive   *bool      `json:"is_fee_inclusive,omitempty"`                                                                            // Is Fee Inclusive
	NotCancellable   *bool      `json:"not_cancellable,omitempty"`                                                                             // Not Cancellable
	NotRefundable    *bool      `json:"not_refundable,omitempty"`                                                                              // Not Refundable
	Description      *string    `json:"description,omitempty" gorm:"type:text"`                                                                // Description
}
