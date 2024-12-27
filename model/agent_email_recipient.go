package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentEmailRecipient Model
type AgentEmailRecipient struct {
	basic.Base
	basic.DataOwner
	AgentEmailRecipientAPI
	Person      *Person      `json:"person" gorm:"foreignKey:ToPersonID;references:ID"`
	Employee    *Employee    `json:"employee" gorm:"foreignKey:ToEmployeeID;references:ID"`
	UserAccount *UserAccount `json:"user_account" gorm:"foreignKey:ToUserAccountID;references:ID"`
	Agent       *Agent       `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	MessageType *MessageType `json:"message_type,omitempty" gorm:"foreignKey:MessageTypeID;references:ID"`
}

// AgentEmailRecipientAPI API
type AgentEmailRecipientAPI struct {
	AgentID         *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	MessageTypeID   *uuid.UUID `json:"message_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	AttentionType   *string    `json:"attention_type,omitempty" gorm:"type:varchar(3)"`
	ToEmail         *string    `json:"to_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`
	ToDisplay       *string    `json:"to_display,omitempty" gorm:"type:varchar(256)" validate:"omitempty,alphanumunicodespace"`
	ToPersonID      *uuid.UUID `json:"to_person_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	ToEmployeeID    *uuid.UUID `json:"to_employee_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	ToUserAccountID *uuid.UUID `json:"to_user_account_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
}

// AgentEmailRecipientData model
type AgentEmailRecipientData struct {
	AgentEmailRecipient
	MessageTypeAPI
	NumberOfRecipients *int64  `json:"number_of_recipients,omitempty"`
	ToEmail            *string `json:"to_email,omitempty" validate:"omitempty,email"`
	// Email *[]Email `json:"to_email,omitempty" gorm:"foreignKey:MessageTypeID;references:ID"`
}

// AgentEmailRecipientRequestAPI API
type AgentEmailRecipientRequestAPI struct {
	MessageTypeID   *uuid.UUID `json:"message_type_id,omitempty"`
	ToEmail         *string    `json:"to_email,omitempty" validate:"omitempty,email"`
	ToDisplay       *string    `json:"to_display,omitempty" validate:"omitempty,alphanumunicodespace"`
	ToPersonID      *uuid.UUID `json:"to_person_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	ToUserAccountID *uuid.UUID `json:"to_user_account_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	ToEmployeeID    *uuid.UUID `json:"to_employee_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
}
