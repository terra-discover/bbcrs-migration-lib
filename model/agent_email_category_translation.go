package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentEmailCategoryTranslation Model
type AgentEmailCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	EmailCategory *EmailCategory `json:"email_category,omitempty" gorm:"foreignKey:EmailCategoryID"`
	AgentEmailCategoryTranslationAPI
	EmailTemplate *EmailTemplate `json:"email_template"`
}

// AgentEmailCategoryTranslationAPI API
type AgentEmailCategoryTranslationAPI struct {
	EmailCategoryID      *uuid.UUID `json:"email_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:agent_email_category_translation_unique;not null" swaggertype:"string" format:"uuid"`
	EmailTemplateID      *uuid.UUID `json:"email_template_id,omitempty" gorm:"type:varchar(36);uniqueIndex:agent_email_category_translation_unique;not null" swaggertype:"string" format:"uuid"`
	LanguageCode         *string    `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:agent_email_category_translation_unique;not null" example:"en"`
	EmailCategoryName    *string    `json:"email_category_name,omitempty" gorm:"type:varchar(256);"`
	EmailTemplateName    *string    `json:"email_template_name,omitempty" gorm:"type:varchar(256);"`
	MessageCategoryName  *string    `json:"message_category_name,omitempty" gorm:"type:varchar(256);"`
	EmailTemplateSubject *string    `json:"email_template_subject,omitempty" gorm:"type:varchar(256);"`
	EmailTemplateBody    *string    `json:"email_template_body,omitempty" gorm:"type:varchar(256);"`
}

// AgentEmailCategoryTranslationData Agent Email Category Translation Data
type AgentEmailCategoryTranslationData struct {
	AgentEmailCategoryTranslation
	MessageCategory   *MessageCategory `json:"message_category,omitempty" gorm:"foreignKey:MessageCategoryID"`
	MessageCategoryID *uuid.UUID       `json:"-"`
	MessageType       *MessageType     `json:"message_type,omitempty" gorm:"foreignKey:MessageTypeID"`
	MessageTypeID     *uuid.UUID       `json:"-"`
}

// EmailTemplateTranslationData Email Template Translation Data
type EmailTemplateTranslationData struct {
	EmailTemplateTranslation
	EmailCategory     *EmailCategory   `json:"email_category,omitempty" gorm:"foreignKey:EmailCategoryID"`
	EmailCategoryID   *uuid.UUID       `json:"-"`
	MessageCategory   *MessageCategory `json:"message_category,omitempty" gorm:"foreignKey:MessageCategoryID"`
	MessageCategoryID *uuid.UUID       `json:"-"`
	MessageType       *MessageType     `json:"message_type,omitempty" gorm:"foreignKey:MessageTypeID"`
	MessageTypeID     *uuid.UUID       `json:"-"`
}
