package model

import (
	"strings"

	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentEmailCategory Model
type AgentEmailCategory struct {
	basic.Base
	basic.DataOwner
	AgentEmailCategoryAPI
	AgentID       *uuid.UUID     `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	Agent         *Agent         `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	EmailCategory *EmailCategory `json:"email_category,omitempty"`
}

// AgentEmailCategoryAPI API
type AgentEmailCategoryAPI struct {
	EmailCategoryID *uuid.UUID `json:"email_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsDefault       *bool      `json:"is_default,omitempty"`
}

// AgentEmailCategoryDetail struct
type AgentEmailCategoryDetail struct {
	EmailCategory
	MessageCategoryAPI
	IsDefault       *bool   `json:"is_default,omitempty"`                           // Is Default
	MessageTypeCode *int    `json:"message_type_code,omitempty" example:"1"`        // Message Type Code
	MessageTypeName *string `json:"message_type_name,omitempty" example:"Feedback"` // Message Type Name

}

// AgentEmailCategoryRequestAPI struct
type AgentEmailCategoryRequestAPI struct {
	EmailCategoryName *string    `json:"email_category_name,omitempty"`
	EmailCategoryID   *uuid.UUID `json:"email_category_id,omitempty"`
	Description       *string    `json:"description,omitempty"`
}

// Seed AgentEmailCategory Data
func (agentEmailCategory *AgentEmailCategory) Seed() *[]AgentEmailCategory {
	// isDefault := true
	// agentID, _ := uuid.Parse("8d5fb953-73af-4a07-b450-f8641b40e383")
	// emailCategoryID, _ := uuid.Parse("3fa85f64-5717-4562-b3fc-2c963f66ab00")
	// agentEmailCategoryID, _ := uuid.Parse("4fa85f64-5717-4562-b3fc-2c963f77ab00")
	// seed := AgentEmailCategory{
	// 	AgentID: &agentID,
	// Base: basic.Base{
	// 		ID: &agentEmailCategoryID,
	// 	},
	// 	AgentEmailCategoryAPI: AgentEmailCategoryAPI{
	// 		EmailCategoryID: &emailCategoryID,
	// 		IsDefault:       &isDefault,
	// 	},
	// }
	// _ = lib.Merge(seed, &agentEmailCategory)
	// return agentEmailCategory

	data := []AgentEmailCategory{}

	items := []string{
		"8d5fb953-73af-4a07-b450-f8641b40e383|3fa85f64-5717-4562-b3fc-2c963f66ab00|4fa85f64-5717-4562-b3fc-2c963f77ab00",
		"8d5fb953-73af-4a07-b450-f8641b40e383|b2d72e27-1f02-4f35-8fb7-89b468b6a4af|f53f6aa9-89ff-4121-85c1-d5266b0ec353",
	}

	for i := range items {
		content := strings.Split(items[i], "|")
		isDefault := true
		agentID, _ := uuid.Parse(content[0])
		emailCategoryID, _ := uuid.Parse(content[1])
		agentEmailCategoryID, _ := uuid.Parse(content[2])
		data = append(data, AgentEmailCategory{
			AgentID: &agentID,
			Base: basic.Base{
				ID: &agentEmailCategoryID,
			},
			AgentEmailCategoryAPI: AgentEmailCategoryAPI{
				IsDefault:       &isDefault,
				EmailCategoryID: &emailCategoryID,
			},
		})
	}
	return &data
}
