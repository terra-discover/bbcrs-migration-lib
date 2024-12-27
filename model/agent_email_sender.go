package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentEmailSender Model
type AgentEmailSender struct {
	basic.Base
	basic.DataOwner
	AgentEmailSenderAPI
	Person      *Person      `json:"person" gorm:"foreignKey:FromPersonID;references:ID"`
	Employee    *Employee    `json:"employee" gorm:"foreignKey:FromEmployeeID;references:ID"`
	UserAccount *UserAccount `json:"user_account" gorm:"foreignKey:FromUserAccountID;references:ID"`
	Agent       *Agent       `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	MessageType *MessageType `json:"message_type" gorm:"foreignKey:MessageTypeID;references:ID"`
}

// AgentEmailSenderAPI API
type AgentEmailSenderAPI struct {
	AgentID           *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	MessageTypeID     *uuid.UUID `json:"message_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid" validate:"required"`
	FromEmail         *string    `json:"from_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`
	FromDisplay       *string    `json:"from_display,omitempty" gorm:"type:varchar(256)" validate:"omitempty,alphanumunicodespace"`
	FromPersonID      *uuid.UUID `json:"from_person_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	FromEmployeeID    *uuid.UUID `json:"from_employee_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	FromUserAccountID *uuid.UUID `json:"from_user_account_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
}
