package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViolationReason Violation Reason
type ViolationReason struct {
	basic.Base
	basic.DataOwner
	ViolationReasonAPI
	ProductType *ProductType `json:"product_type,omitempty" gorm:"foreignKey:ProductTypeID;references:ID"` // product type
}

// ViolationReasonAPI Violation Reason API
type ViolationReasonAPI struct {
	ViolationReasonCode *string    `json:"violation_reason_code,omitempty" gorm:"type:varchar(36);not null"`         // Violation Reason Code
	ViolationReasonName *string    `json:"violation_reason_name,omitempty" gorm:"type:varchar(256);not null"`        // Violation Reason Name
	ProductTypeID       *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // product type id
	Description         *string    `json:"description,omitempty" gorm:"type:text"`                                   // description
}
