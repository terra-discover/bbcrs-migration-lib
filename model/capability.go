package model

import (
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Capability Capability
type Capability struct {
	basic.Base
	basic.DataOwner
	CapabilityAPI
	CapabilityTranslation *CapabilityTranslation `json:"capability_translation,omitempty"` // Capability Translation
}

// CapabilityAPI Capability API
type CapabilityAPI struct {
	CapabilityCode *string `json:"capability_code,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Capability Code
	CapabilityName *string `json:"capability_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Capability Name
	Description    *string `json:"description,omitempty" gorm:"type:text"`                                                             // Description
}

// Seed data
func (s Capability) Seed() *[]Capability {
	data := []Capability{}
	items := []string{
		"view|View",
		"create|Create",
		"edit|Edit",
		"delete|Delete",
		"bulk_update|Mass Update",
		"export|Export",
		"publish|Publish",
		"activate|Activate/ Deactivate",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		var code string = contents[0]
		var name string = contents[1]
		data = append(data, Capability{
			CapabilityAPI: CapabilityAPI{
				CapabilityCode: &code,
				CapabilityName: &name,
			},
		})
	}

	return &data
}
