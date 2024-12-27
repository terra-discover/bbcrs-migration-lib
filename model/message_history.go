package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MessageHistory struct
type MessageHistory struct {
	basic.Base
	basic.DataOwner
	MessageHistoryAPI
	Message         *Message         `json:"message,omitempty"`
	UserAccount     *UserAccount     `json:"user_account,omitempty"`
	OperatingSystem *OperatingSystem `json:"operating_system,omitempty"`
	//Device          *Device
}

// MessageHistoryAPI struct
type MessageHistoryAPI struct {
	MessageID         *uuid.UUID       `json:"message_id,omitempty" gorm:"type:varchar(36);not null"`                                     // Message ID
	MessageStatus     *int             `json:"message_status,omitempty" gorm:"type:smallint"`                                             // MessageStatus
	UserAccountID     *uuid.UUID       `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null"`                                // UserAccount ID
	IPAddress         *string          `json:"ip_address,omitempty" gorm:"type:varchar(256)"`                                             // IP Address
	UserAgent         *string          `json:"user_agent,omitempty" gorm:"type:text"`                                                     // User Agent
	DeviceID          *uuid.UUID       `json:"device_id,omitempty" gorm:"type:varchar(36)"`                                               // Device ID
	OperatingSystemID *uuid.UUID       `json:"operating_system_id,omitempty" gorm:"type:varchar(36)"`                                     // OperatingSystem ID
	Comment           *string          `json:"comment,omitempty" gorm:"type:text"`                                                        // Comment
	Timestamp         *strfmt.DateTime `json:"timestamp,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`       // Timestamp
	LocalTimestamp    *strfmt.DateTime `json:"local_timestamp,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // Local Timestamp
}
