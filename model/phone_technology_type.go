package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// PhoneTechnologyType Phone Technology Type
type PhoneTechnologyType struct {
	basic.Base
	basic.DataOwner
	PhoneTechnologyTypeAPI
	PhoneTechnologyTypeTranslation *PhoneTechnologyTypeTranslation `json:"phone_technology_type_translation,omitempty"` // Phone Technology Type Translation
}

// PhoneTechnologyTypeAPI Phone Technology Type API
type PhoneTechnologyTypeAPI struct {
	PhoneTechnologyTypeCode *int    `json:"phone_technology_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null" example:"1"`         // Phone Technology Type Code
	PhoneTechnologyTypeName *string `json:"phone_technology_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Voice"` // Phone Technology Type Name
}

// Seed data
func (s PhoneTechnologyType) Seed() *[]PhoneTechnologyType {
	data := []PhoneTechnologyType{}
	items := []string{
		"Voice",
		"Data",
		"Fax",
		"Pager",
		"Mobile",
		"TTY",
		"Telex",
		"Voice over IP",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, PhoneTechnologyType{
			PhoneTechnologyTypeAPI: PhoneTechnologyTypeAPI{
				PhoneTechnologyTypeCode: &code,
				PhoneTechnologyTypeName: &name,
			},
		})
	}

	return &data
}
