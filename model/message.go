package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Message struct
type Message struct {
	basic.Base
	basic.DataOwner
	MessageAPI
	MessageType *MessageType `json:"message_type,omitempty" gorm:"foreignkey:ID"` // Message Type
	Person      *Person      `json:"person,omitempty" gorm:"foreignkey:ID"`       // Person
	Employee    *Employee    `json:"employee,omitempty" gorm:"foreignkey:ID"`     // Employee
	UserAccount *UserAccount `json:"user_account,omitempty" gorm:"foreignkey:ID"` // User Account
}

// MessageAPI Message API
type MessageAPI struct {
	MessageTypeID     *uuid.UUID `json:"message_type_id,omitempty" gorm:"type:varchar(36)"`
	MessageStatus     *int       `json:"message_status,omitempty" gorm:"type:smallint"`
	Priority          *int       `json:"priority,omitempty" gorm:"type:smallint" description:"smaller numbers indicate higher priority"`
	FromEmail         *string    `json:"from_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`
	FromDisplay       *string    `json:"from_display,omitempty" gorm:"type:varchar(256)" validate:"omitempty,alphanumunicodespace"`
	FromPersonID      *uuid.UUID `json:"from_person_id,omitempty" gorm:"type:varchar(36)"`
	FromEmployeeID    *uuid.UUID `json:"from_employee_id,omitempty" gorm:"type:varchar(36)"`
	FromUserAccountID *uuid.UUID `json:"from_user_account_id,omitempty" gorm:"type:varchar(36)"`
	ReplyToEmail      *string    `json:"reply_to_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`
	ReplyToDisplay    *string    `json:"reply_to_display,omitempty" gorm:"type:varchar(256)" validate:"omitempty,alphanumunicodespace"`
	Subject           *string    `json:"subject,omitempty" gorm:"type:varchar(256)"`
	Body              *string    `json:"body,omitempty" gorm:"type:text"`
	BodyFileName      *string    `json:"body_file_name,omitempty" gorm:"type:varchar(256)"`
}
