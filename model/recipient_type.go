package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RecipientType Recipient Type
type RecipientType struct {
	basic.Base
	basic.DataOwner
	RecipientTypeAPI
	RecipientTypeTranslation *RecipientTypeTranslation `json:"recipient_type_translation,omitempty"` // Recipient Type Translation
	BuiltInType              *BuiltInType              `json:"build_in_type,omitempty"`              // Built In Type
}

// RecipientTypeAPI Recipient Type API
type RecipientTypeAPI struct {
	RecipientTypeCode *int       `json:"recipient_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`                   // Recipient Type Code
	RecipientTypeName *string    `json:"recipient_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Agent/ Internal"` // Recipient Type Name
	BuiltInTypeID     *uuid.UUID `json:"built_in_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`                                   // Built In Type ID
}

// Seed data
func (s RecipientType) Seed() *[]RecipientType {
	data := []RecipientType{}
	items := []string{
		"Agent/ Internal",
		"Vendor",
		"Sub Agent",
		"Corporate Client",
		"Member",
		"Retail Customer",
		"Agent User",
		"Vendor User",
		"Sub Agent User",
		"Corporate Client User",
		"Agent Superuser",
		"Vendor Superuser",
		"Sub Agent Superuser",
		"Corporate Client Superuser",
		"Agent Employee",
		"Corporate Client Employee",
		"Corporate Booker",
		"Traveler",
		"Travel Consultant",
		"Corporate Client Consultant",
		"Corporate Personal Traveler",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		var builtInTypeID = uuid.New()
		data = append(data, RecipientType{
			RecipientTypeAPI: RecipientTypeAPI{
				RecipientTypeCode: &code,
				RecipientTypeName: &name,
				BuiltInTypeID:     &builtInTypeID,
			},
		})
	}

	return &data
}
