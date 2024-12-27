package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// UserType User Type
type UserType struct {
	basic.Base
	basic.DataOwner
	UserTypeAPI
	UserTypeTranslation *UserTypeTranslation `json:"user_type_translation,omitempty"` // User Type Translation
}

// UserTypeAPI User Type API
type UserTypeAPI struct {
	UserTypeCode          *string `json:"user_type_code,omitempty" gorm:"type:varchar(36);not null;"`  // User Type Code
	UserTypeName          *string `json:"user_type_name,omitempty" gorm:"type:varchar(256);not null;"` // User Type Name
	IsSystem              *bool   `json:"is_system,omitempty"`                                         // Is System
	AssignAllCapabilities *bool   `json:"assign_all_capabilities,omitempty"`                           // Assign All Capabilities
	AssignAllMenuLinks    *bool   `json:"assign_all_menu_links,omitempty"`                             // Assign All Menu Links
	Comment               *string `json:"comment,omitempty"`                                           // Comment
	Description           *string `json:"description,omitempty"`                                       // Description
}

// Seed data
func (s UserType) Seed() *[]UserType {
	type rawData struct {
		UserTypeCode          string
		UserTypeName          string
		IsSystem              bool
		AssignAllCapabilities bool
		AssignAllMenuLinks    bool
	}

	listRawData := []rawData{
		{
			UserTypeCode:          "ADM",
			UserTypeName:          "Administrator",
			IsSystem:              true,
			AssignAllCapabilities: true,
			AssignAllMenuLinks:    true,
		},
	}

	var result []UserType

	for _, item := range listRawData {
		result = append(result, UserType{
			UserTypeAPI: UserTypeAPI{
				UserTypeCode:          &item.UserTypeCode,
				UserTypeName:          &item.UserTypeName,
				IsSystem:              &item.IsSystem,
				AssignAllCapabilities: &item.AssignAllCapabilities,
				AssignAllMenuLinks:    &item.AssignAllMenuLinks,
			},
		})
	}

	return &result
}
