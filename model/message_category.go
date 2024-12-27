package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// MessageCategory Message Category
type MessageCategory struct {
	basic.Base
	basic.DataOwner
	MessageCategoryAPI
}

// MessageCategoryAPI Message Category API
type MessageCategoryAPI struct {
	MessageCategoryCode *int    `json:"message_category_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null"`     // Message Category Code
	MessageCategoryName *string `json:"message_category_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Message Category Name
	Priority            *int    `json:"priority,omitempty" gorm:"type:smallint" description:"smaller numbers indicate higher priority"`           // Priority
}

// MessageCategories migrator
type MessageCategories []MessageCategory

// Seed data
func (messageCategories *MessageCategories) Seed() *MessageCategories {
	name := []string{"Newsletter", "Survey", "Pre-Arrival", "Post-Departure", "Invoice Per Transaction", "Invoice Batch"}
	id := []string{"06461623-39f8-4b0b-a128-726fa4fa5811", "20c7e10e-d0be-46cd-8411-77d00cb3962a", "81cfd4c5-f09d-4c2e-904d-d9685cd2dcaf", "62de0f2f-7c0e-4a31-b147-e7627b48f731", "30714390-e872-4e96-b927-d39de8c09ee3", "470df3ca-6fe3-421c-99e5-4ec2524ae15e"}
	for i := 0; i < len(name); i++ {
		*messageCategories = append(*messageCategories, MessageCategory{
			Base: basic.Base{
				ID: lib.UUIDPtr(uuid.MustParse(id[i])),
			},
			MessageCategoryAPI: MessageCategoryAPI{
				MessageCategoryCode: lib.Intptr(i + 1),
				MessageCategoryName: lib.Strptr(name[i]),
				Priority:            nil,
			},
		})
	}
	return messageCategories
}
