package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CapabilityRecipient Capability Recipient
type CapabilityRecipient struct {
	basic.Base
	basic.DataOwner
	CapabilityRecipientAPI
}

// CapabilityRecipientAPI Capability Recipient API
type CapabilityRecipientAPI struct {
	CapabilityID    *uuid.UUID `json:"capability_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Capability ID
	AttentionType   *string    `json:"attention_type,omitempty" gorm:"type:varchar(3)"`                                             // Attention Type
	ToEmail         *string    `json:"to_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`                      // To Email
	ToDisplay       *string    `json:"to_display,omitempty" gorm:"type:varchar(256)"`                                               // To Display
	ToPersonID      *uuid.UUID `json:"to_person_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`           // To Person ID
	ToEmployeeID    *uuid.UUID `json:"to_employee_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`         // To Employee ID
	ToUserAccountID *uuid.UUID `json:"to_user_account_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`     // To User Account ID
}
