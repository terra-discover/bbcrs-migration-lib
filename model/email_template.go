package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// EmailTemplate model
type EmailTemplate struct {
	basic.Base
	basic.DataOwner
	EmailTemplateAPI
	MessageType   *MessageType   `json:"message_type,omitempty"`
	EmailCategory *EmailCategory `json:"email_category,omitempty"`
}

// EmailTemplateAPI struct
type EmailTemplateAPI struct {
	EmailTemplateCode     *string    `json:"email_template_code,omitempty" gorm:"type:varchar(36);not null"`           // Email Template Code
	EmailTemplateName     *string    `json:"email_template_name,omitempty" gorm:"type:varchar(256);not null"`          // Email Template Name
	MessageTypeID         *uuid.UUID `json:"message_type_id,omitempty" gorm:"type:varchar(36)"`                        // Message Type Id
	EmailCategoryID       *uuid.UUID `json:"email_category_id,omitempty" gorm:"type:varchar(36)"`                      // Email Category Id
	IsTemplate            *bool      `json:"is_template,omitempty"`                                                    // Is Template
	FromEmail             *string    `json:"from_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"` // From Email
	FromDisplay           *string    `json:"from_display,omitempty" gorm:"type:varchar(256)"`                          // From Display
	FromPersonID          *uuid.UUID `json:"from_person_id,omitempty" gorm:"type:varchar(36)"`                         // From Person Id
	FromEmployeeID        *uuid.UUID `json:"from_employee_id,omitempty" gorm:"type:varchar(36)"`                       // From Employee Id
	FromUserAccountID     *uuid.UUID `json:"from_user_account_id,omitempty" gorm:"type:varchar(36)"`                   // From User Account Id
	Subject               *string    `json:"subject,omitempty" gorm:"type:varchar(256)"`                               // Subject
	Body                  *string    `json:"body,omitempty" gorm:"type:text"`                                          // Body
	BodyFileName          *string    `json:"body_file_name,omitempty" gorm:"type:varchar(256)"`                        // Body File Name
	ScheduleID            *uuid.UUID `json:"schedule_id,omitempty" gorm:"type:varchar(36)"`                            // Schedule Id
	HasNews               *bool      `json:"has_news,omitempty"`                                                       // Has News
	NewsID                *uuid.UUID `json:"news_id,omitempty" gorm:"type:varchar(36)"`                                // News Id
	HasSurvey             *bool      `json:"has_survey,omitempty"`                                                     // Has Survey
	AssignAllProductTypes *bool      `json:"assign_all_product_types,omitempty"`                                       // Assign All Product Types
	AssignAllOffices      *bool      `json:"assign_all_offices,omitempty"`                                             // Assign All Offices
	Description           *string    `json:"description,omitempty" gorm:"type:text"`                                   // Description
}

// EmailTemplateData Email Template Data
type EmailTemplateData struct {
	EmailTemplate
	MessageCategory   *MessageCategory `json:"message_category,omitempty" gorm:"foreignKey:MessageCategoryID"`
	MessageCategoryID *uuid.UUID       `json:"-"`
}

// Seed EmailTemplate Data
func (emailTemplate *EmailTemplate) Seed() *EmailTemplate {
	emailTemplateCode := "email-template-a"
	emailTemplateName := "Email Template A"
	messageTypeID := lib.UUIDPtr(uuid.New())
	emailCategoryID, _ := uuid.Parse("3fa85f64-5717-4562-b3fc-2c963f66ab00")
	isTemplate := true
	fromEmail := "From Email A"
	fromDisplay := "From Display A"
	fromPersonID := lib.UUIDPtr(uuid.New())
	fromEmployeeID := lib.UUIDPtr(uuid.New())
	fromUserAccountID := lib.UUIDPtr(uuid.New())
	subject := "Subject A"
	body := "Body A"
	bodyFileName := "4fa75f64-5777-4562-b3fc-2c963f66ab00.html"
	scheduleID := lib.UUIDPtr(uuid.New())
	hasNews := true
	newsID := lib.UUIDPtr(uuid.New())
	hasSurvey := true
	assignAllProductTypes := true
	assignAllOffices := true
	description := "Email Template A Description"
	emailTemplateID, _ := uuid.Parse("4fa75f64-5777-4562-b3fc-2c963f66ab00")
	seed := EmailTemplate{
		Base: basic.Base{
			ID: &emailTemplateID,
		},
		EmailTemplateAPI: EmailTemplateAPI{
			EmailTemplateCode:     &emailTemplateCode,
			EmailTemplateName:     &emailTemplateName,
			MessageTypeID:         messageTypeID,
			EmailCategoryID:       &emailCategoryID,
			IsTemplate:            &isTemplate,
			FromEmail:             &fromEmail,
			FromDisplay:           &fromDisplay,
			FromPersonID:          fromPersonID,
			FromEmployeeID:        fromEmployeeID,
			FromUserAccountID:     fromUserAccountID,
			Subject:               &subject,
			Body:                  &body,
			BodyFileName:          &bodyFileName,
			ScheduleID:            scheduleID,
			HasNews:               &hasNews,
			NewsID:                newsID,
			HasSurvey:             &hasSurvey,
			AssignAllProductTypes: &assignAllProductTypes,
			AssignAllOffices:      &assignAllOffices,
			Description:           &description,
		},
	}
	_ = lib.Merge(seed, &emailTemplate)
	return emailTemplate
}
