package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateEmailRecipient Corporate Email Recipient
type CorporateEmailRecipient struct {
	basic.Base
	basic.DataOwner
	CorporateEmailRecipientAPI
	Corporate   *Corporate   `json:"corporate,omitempty"`                                      // Corporate
	MessageType *MessageType `json:"message_type,omitempty"`                                   // Message Type
	Person      *Person      `json:"person,omitempty" gorm:"foreignKey:ToPersonID"`            // Person
	Employee    *Employee    `json:"employee,omitempty" gorm:"foreignKey:ToEmployeeID"`        // Employee
	UserAccount *UserAccount `json:"user_account,omitempty" gorm:"foreignKey:ToUserAccountID"` // User Account
}

// CorporateEmailRecipientAPI Corporate Email Recipient API
type CorporateEmailRecipientAPI struct {
	CorporateID     *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`    // corporate_id
	MessageTypeID   *uuid.UUID `json:"message_type_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // message_type_id
	AttentionType   *string    `json:"attention_type,omitempty" gorm:"type:varchar(3)"`                          // attention_type
	ToEmail         *string    `json:"to_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`   // to_email
	ToDisplay       *string    `json:"to_display,omitempty" gorm:"type:varchar(256)"`                            // to_display
	ToPersonID      *uuid.UUID `json:"to_person_id,omitempty" gorm:"type:varchar(36)"`                           // to_person_id
	ToEmployeeID    *uuid.UUID `json:"to_employee_id,omitempty" gorm:"type:varchar(36)"`                         // to_employee_id
	ToUserAccountID *uuid.UUID `json:"to_user_account_id,omitempty" gorm:"type:varchar(36)"`                     // to_user_account_id
}
