package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateViolationReason Corporate Violation Reason
type CorporateViolationReason struct {
	basic.Base
	basic.DataOwner
	CorporateViolationReasonAPI
	Corporate       *Corporate       `json:"corporate,omitempty" gorm:"foreignKey:CorporateID;references:ID"`              // corporate
	ViolationReason *ViolationReason `json:"violation_reason,omitempty" gorm:"foreignKey:ViolationReasonID;references:ID"` // violation reason

}

// CorporateViolationReasonAPI Corporate Violation Reason API
type CorporateViolationReasonAPI struct {
	CorporateID       *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`        // corporate id
	ViolationReasonID *uuid.UUID `json:"violation_reason_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // violation reason id
}

type CorporateViolationReasonData struct {
	basic.Base
	basic.DataOwner
	CorporateViolationReasonAPI
	Corporate       *Corporate       `json:"corporate,omitempty" gorm:"foreignKey:CorporateID;references:ID"`              // corporate
	ViolationReason *ViolationReason `json:"violation_reason,omitempty" gorm:"foreignKey:ViolationReasonID;references:ID"` // violation reason
	ProductTypeID   *uuid.UUID       `json:"product_type_id,omitempty" `
	ProductTypeCode *string          `json:"product_type_code,omitempty" `
	ProductTypeName *string          `json:"product_type_name,omitempty" `
}
