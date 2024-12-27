package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// ContactType model
type ContactType struct {
	basic.Base
	basic.DataOwner
	ContactTypeAPI
	ContactTypeTranslation *ContactTypeTranslation `json:"contact_type_translation,omitempty"`
}

// ContactTypeAPI model
type ContactTypeAPI struct {
	ContactTypeCode *int    `json:"contact_type_code,omitempty" gorm:"type:smallint;not null;index:unique_contact_type__contact_type_code,unique,where:deleted_at is null" example:"1"`
	ContactTypeName *string `json:"contact_type_name,omitempty" gorm:"type:varchar(256);"`
}

// Seed data
func (s ContactType) Seed() *[]ContactType {
	data := []ContactType{}
	items := []string{
		"Emergency contact",
		"Travel arranger",
		"Daytime contact",
		"Evening contact",
		"Contact",
		"Toll free number",
		"Guest use",
		"Pickup contact",
		"Electronic document reference",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, ContactType{
			ContactTypeAPI: ContactTypeAPI{
				ContactTypeCode: &code,
				ContactTypeName: &name,
			},
		})
	}

	return &data
}
