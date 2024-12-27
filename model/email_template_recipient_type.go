package model

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EmailTemplateRecipientType model
type EmailTemplateRecipientType struct {
	basic.Base
	basic.DataOwner
	EmailTemplateRecipientTypeAPI
	AgentID       *uuid.UUID     `json:"agent_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	Agent         *Agent         `json:"agent,omitempty"`
	EmailTemplate *EmailTemplate `json:"email_template,omitempty"`
	RecipientType *RecipientType `json:"recipient_type,omitempty"`
	Corporate     *Corporate     `json:"corporate,omitempty"`
}

// EmailTemplateRecipientTypeAPI struct
type EmailTemplateRecipientTypeAPI struct {
	AttentionType   *string    `json:"attention_type,omitempty" gorm:"type:varchar(3)"`
	EmailTemplateID *uuid.UUID `json:"email_template_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	RecipientTypeID *uuid.UUID `json:"recipient_type_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	CorporateID     *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
}

// Seed EmailTemplateRecipientType Data
func (emailTemplateRecipientType *EmailTemplateRecipientType) Seed() *EmailTemplateRecipientType {
	attentionType := "to"
	emailTemplateID, _ := uuid.Parse("4fa75f64-5777-4562-b3fc-2c963f66ab00")
	recipientTypeID, _ := uuid.Parse("e02547cf-d7b2-4346-b704-067b9a4ed5c0")
	corporateID := lib.UUIDPtr(uuid.New())
	agentID, _ := uuid.Parse("8d5fb953-73af-4a07-b450-f8641b40e383")
	agentEmailTemplateID, _ := uuid.Parse("4ba85f64-5717-4562-b3fc-2c963f66ab00")
	seed := EmailTemplateRecipientType{
		AgentID: &agentID,
		Base: basic.Base{
			ID: &agentEmailTemplateID,
		},
		EmailTemplateRecipientTypeAPI: EmailTemplateRecipientTypeAPI{
			AttentionType:   &attentionType,
			EmailTemplateID: &emailTemplateID,
			RecipientTypeID: &recipientTypeID,
			CorporateID:     corporateID,
		},
	}
	_ = lib.Merge(seed, &emailTemplateRecipientType)
	return emailTemplateRecipientType
}
