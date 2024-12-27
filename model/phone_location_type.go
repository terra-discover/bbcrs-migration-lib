package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// PhoneLocationType Phone Location Type
type PhoneLocationType struct {
	basic.Base
	basic.DataOwner
	PhoneLocationTypeAPI
	PhoneLocationTypeTranslation *PhoneLocationTypeTranslation `json:"phone_location_type_translation,omitempty"` // Phone Location Type Translation
}

// PhoneLocationTypeAPI Phone Location Type API
type PhoneLocationTypeAPI struct {
	PhoneLocationTypeCode *int    `json:"phone_location_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null" example:"1"`        // Phone Location Type Code
	PhoneLocationTypeName *string `json:"phone_location_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Home"` // Phone Location Type Name
}

// Seed data
func (s PhoneLocationType) Seed() *[]PhoneLocationType {
	data := []PhoneLocationType{}
	items := []string{
		"Brand reservations office",
		"Central reservations office",
		"Property reservation office",
		"Property direct",
		"Sales office",
		"Home",
		"Office",
		"Other",
		"Managing company",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, PhoneLocationType{
			PhoneLocationTypeAPI: PhoneLocationTypeAPI{
				PhoneLocationTypeCode: &code,
				PhoneLocationTypeName: &name,
			},
		})
	}

	return &data
}
