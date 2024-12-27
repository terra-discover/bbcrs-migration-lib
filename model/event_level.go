package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// EventLevel Event Level
type EventLevel struct {
	basic.Base
	basic.DataOwner
	EventLevelAPI
	EventLevelTranslation *EventLevelTranslation `json:"event_level_translation,omitempty"` // Event Level Translation
}

// EventLevelAPI Event Level API
type EventLevelAPI struct {
	EventLevelCode *int    `json:"event_level_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`         // Event Level Code
	EventLevelName *string `json:"event_level_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Error"` // Event Level Name
	Level          *int    `json:"level,omitempty" example:"1"`                                                                                         // Level
}

// Seed data
func (EventLevel) Seed() *[]EventLevel {
	data := []EventLevel{}
	items := []string{
		"Error",
		"Warning",
		"Info",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, EventLevel{
			EventLevelAPI: EventLevelAPI{
				EventLevelCode: &code,
				EventLevelName: &name,
			},
		})
	}

	return &data
}
