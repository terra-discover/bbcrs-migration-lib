package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EventLog Event Log
type EventLog struct {
	basic.Base
	basic.DataOwner
	EventID              *uuid.UUID `json:"event_id,omitempty" gorm:"type:varchar(36);not null"`      // Event Id
	SessionID            *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)"`             // Session Id
	ModuleID             *uuid.UUID `json:"module_id,omitempty" gorm:"type:varchar(36)"`              // Module Id
	PageID               *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36)"`                // Page Id
	UserAccountID        *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36)"`        // User Account Id
	IPAddress            *string    `json:"ip_address,omitempty" gorm:"type:varchar(256)"`            // Ip Address
	ProposalID           *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36)"`            // Proposal Id
	ReservationID        *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36)"`         // Reservation Id
	MessageID            *uuid.UUID `json:"message_id,omitempty" gorm:"type:varchar(36)"`             // Message Id
	CurrencyConversionID *uuid.UUID `json:"currency_conversion_id,omitempty" gorm:"type:varchar(36)"` // Currency Conversion Id
	OldValue             *string    `json:"old_value,omitempty" gorm:"type:text"`                     // Old Value
	NewValue             *string    `json:"new_value,omitempty" gorm:"type:text"`                     // New Value
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                   // Description
	Event                *Event     `json:"event,omitempty"`                                          // Event
	Person               *Person    `json:"person,omitempty" gorm:"foreignKey:ID"`                    // Person
}
