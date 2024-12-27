package model

import (
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MessageVariable Message Variable
type MessageVariable struct {
	basic.Base
	basic.DataOwner
	MessageVariableAPI
	MessageVariableTranslation *MessageVariableTranslation `json:"message_variable_translation,omitempty"` // Message Variable Translation
}

// MessageVariableAPI Message Variable API
type MessageVariableAPI struct {
	MessageVariableCode *string `json:"message_variable_code,omitempty" gorm:"type:varchar(26);index:,unique,where:deleted_at is null;not null" example:"departure_date"`  // Message Variable Code
	MessageVariableName *string `json:"message_variable_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Departure Date"` // Message Variable Name
	DummyValue          *string `json:"dummy_value,omitempty" gorm:"type:text"`                                                                                            // Dummy Value
	Description         *string `json:"description,omitempty" gorm:"type:text"`                                                                                            // Description
}

// Seed data
func (s MessageVariable) Seed() *[]MessageVariable {
	data := []MessageVariable{}
	items := []string{
		"departure_date|Departure Date",
		"arrival_date|Arrival Date",
		"check_in_date|Check-In Date",
		"check_out_date|Check_Out Date",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		var code string = contents[0]
		var name string = contents[1]
		data = append(data, MessageVariable{
			MessageVariableAPI: MessageVariableAPI{
				MessageVariableCode: &code,
				MessageVariableName: &name,
			},
		})
	}

	return &data
}
