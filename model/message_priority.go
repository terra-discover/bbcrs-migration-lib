package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MessagePriority struct
type MessagePriority struct {
	basic.Base
	basic.DataOwner
	MessagePriorityAPI
}

// MessagePriorityAPI struct
type MessagePriorityAPI struct {
	MessagePriorityCode *int    `json:"message_priority_code,omitempty" gorm:"type:smallint;not null"`
	MessagePriorityName *string `json:"message_priority_name,omitempty" gorm:"type:varchar(256);not null"`
	Priority            *int    `json:"priority,omitempty" gorm:"type:smallint;default:0"`
}

// MessagePriorities struct
type MessagePriorities []MessagePriority

// Seed for seeder
func (m *MessagePriorities) Seed() *MessagePriorities {
	name := []string{"Notification", "Alert", "CRM"}
	for i := 0; i < len(name); i++ {
		*m = append(*m, MessagePriority{
			MessagePriorityAPI: MessagePriorityAPI{
				MessagePriorityCode: lib.Intptr(i + 1),
				MessagePriorityName: lib.Strptr(name[i]),
			},
		})
	}
	return m
}
