package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// RecurrenceType Recurrence Type
type RecurrenceType struct {
	basic.Base
	basic.DataOwner
	RecurrenceTypeAPI
	RecurrenceTypeTranslation *RecurrenceTypeTranslation `json:"recurrence_type_translation,omitempty"` // Recurrence Type Translation
}

// RecurrenceTypeAPI Recurrence Type API
type RecurrenceTypeAPI struct {
	RecurrenceTypeCode *int    `json:"recurrence_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`         // Recurrence Type Code
	RecurrenceTypeName *string `json:"recurrence_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Daily"` // Recurrence Type Name
}

// Seed data
func (s RecurrenceType) Seed() *[]RecurrenceType {
	data := []RecurrenceType{}
	items := []string{
		"Daily",
		"Hourly",
		"Half day",
		"Additions per stay",
		"Per occurrence",
		"Per event",
		"Per person",
		"First use",
		"One time use",
		"Per minute",
		"Per function",
		"Per stay",
		"Complimentary",
		"Other",
		"Maximum charge",
		"Over-minute charge",
		"Weekly",
		"Per room per stay",
		"Per room per night",
		"Per person per stay",
		"Per person per night",
		"Minimum charge",
		"Per rental",
		"Per item",
		"Per room",
		"Per reservation/booking",
		"Per gallon",
		"Per dozen",
		"Per tray",
		"Per order",
		"Per unit",
		"One way",
		"Round trip",
		"Per ticket",
		"Monthly",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, RecurrenceType{
			RecurrenceTypeAPI: RecurrenceTypeAPI{
				RecurrenceTypeCode: &code,
				RecurrenceTypeName: &name,
			},
		})
	}

	return &data
}
