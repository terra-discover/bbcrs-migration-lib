package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MessageRecipient Message Recipient
type MessageRecipient struct {
	basic.Base
	basic.DataOwner
	MessageRecipientAPI
	Message     *Message     `json:"message,omitempty" gorm:"foreignkey:ID"`      // Message
	Person      *Person      `json:"person,omitempty" gorm:"foreignkey:ID"`       // Person
	Employee    *Employee    `json:"employee,omitempty" gorm:"foreignkey:ID"`     // Employee
	UserAccount *UserAccount `json:"user_account,omitempty" gorm:"foreignkey:ID"` // User Account
}

// MessageRecipientAPI Message Recipient API
type MessageRecipientAPI struct {
	MessageID       *uuid.UUID `json:"message_id,omitempty" gorm:"type:varchar(36);not null"`                                             // Message ID
	AttentionType   *string    `json:"attention_type,omitempty" gorm:"type:varchar(3)" description:"Type of attention, e.g: To, Cc, Bcc"` // Attention Type
	ToEmail         *string    `json:"to_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`                            // To Email
	ToDisplay       *string    `json:"to_display,omitempty" gorm:"type:varchar(256)"`                                                     // To Display
	ToPersonID      *uuid.UUID `json:"to_person_id,omitempty" gorm:"type:varchar(36)"`                                                    // To Person ID
	ToEmployeeID    *uuid.UUID `json:"to_employee_id,omitempty" gorm:"type:varchar(36)"`                                                  // To Employee ID
	ToUserAccountID *uuid.UUID `json:"to_user_account_id,omitempty" gorm:"type:varchar(36)"`                                              // To User Account ID
}
