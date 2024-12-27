package model

import (
	"strings"

	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EmailCategory model
type EmailCategory struct {
	basic.Base
	basic.DataOwner
	EmailCategoryAPI
	MessageCategory *MessageCategory `json:"message_category,omitempty"` // Message Category
}

// EmailCategoryAPI model
type EmailCategoryAPI struct {
	EmailCategoryCode *string    `json:"email_category_code,omitempty" gorm:"type:varchar(36)"`               // Email Category Code
	EmailCategoryName *string    `json:"email_category_name,omitempty" gorm:"type:varchar(256)"`              // Email Category Name
	MessageCategoryID *uuid.UUID `json:"message_category_id,omitempty" gorm:"type:varchar(26)" format:"uuid"` // Message Category Id
	Description       *string    `json:"description,omitempty" gorm:"type:text"`                              // Description
}

// Seed EmailCategory Data
func (emailCategory *EmailCategory) Seed() *[]EmailCategory {
	// name := "Email Category A"
	// code := "email-category-a"
	// description := "Email Category A Description"
	// emailCategoryID, _ := uuid.Parse("3fa85f64-5717-4562-b3fc-2c963f66ab00")
	// seed := EmailCategory{
	// Base: basic.Base{
	// 		ID: &emailCategoryID,
	// 	},
	// 	EmailCategoryAPI: EmailCategoryAPI{
	// 		EmailCategoryCode: &code,
	// 		EmailCategoryName: &name,
	// 		Description:       &description,
	// 	},
	// }
	// _ = lib.Merge(seed, &emailCategory)
	// return emailCategory

	data := []EmailCategory{}

	items := []string{
		"Email Category A|email-category-a|Email Category A Description|3fa85f64-5717-4562-b3fc-2c963f66ab00|470df3ca-6fe3-421c-99e5-4ec2524ae15e",
		"Email Category B|email-category-b|Email Category B Description|b2d72e27-1f02-4f35-8fb7-89b468b6a4af|30714390-e872-4e96-b927-d39de8c09ee3",
	}

	for i := range items {
		content := strings.Split(items[i], "|")
		name := content[0]
		code := content[1]
		description := content[2]
		emailCategoryID, _ := uuid.Parse(content[3])
		messageCategoryID, _ := uuid.Parse(content[4])
		data = append(data, EmailCategory{
			Base: basic.Base{
				ID: &emailCategoryID,
			},
			EmailCategoryAPI: EmailCategoryAPI{
				EmailCategoryName: &name,
				EmailCategoryCode: &code,
				Description:       &description,
				MessageCategoryID: &messageCategoryID,
			},
		})
	}
	return &data
}
