package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// EmailTemplateTranslation Email Template Translation
type EmailTemplateTranslation struct {
	basic.Base
	basic.DataOwner
	EmailTemplateTranslationAPI
	EmailTemplate   *EmailTemplate `json:"email_template,omitempty" foreignKey:"EmailTemplateID;references:ID"`
	EmailTemplateID *uuid.UUID     `json:"email_template_id,omitempty" gorm:"type:varchar(36);uniqueIndex:email_template_unique;not null" swaggertype:"string" format:"uuid"` // Email Template id
}

// EmailTemplateTranslationAPI Email Template Translation API
type EmailTemplateTranslationAPI struct {
	LanguageCode          *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:email_template_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	Subject               *string `json:"subject,omitempty" gorm:"type:varchar(256)"`                                                             // Subject
	EmailTemplateName     *string `json:"email_template_name,omitempty" gorm:"type:varchar(256);not null"`                                        // Email Template Name
	Description           *string `json:"description,omitempty" gorm:"type:text"`                                                                 // Description
	Body                  *string `json:"body,omitempty" gorm:"type:text"`                                                                        // Body
	BodyFileName          *string `json:"body_file_name,omitempty" gorm:"type:varchar(256)"`                                                      // Body File Name
	BodyEmailNotification *string `json:"body_email_notification,omitempty" gorm:"-" example:"<html><body><p>Paragraph</p></body></html>"`
}

// Seed EmailTemplateTranslation
func (s EmailTemplateTranslation) Seed() *[]EmailTemplateTranslation {
	data := []EmailTemplateTranslation{}
	items := []string{
		"34eff14b-cb13-4fc0-b08f-efa164d49aa9|en|Test Subject en|Email Template A|<p>Test Body en</p>|34eff14b-cb13-4fc0-b08f-efa164d49aa9.html|4fa75f64-5777-4562-b3fc-2c963f66ab00",
		"f537448b-d498-4fed-8090-e5cd62f363ff|id|Test Subject id|Email Template A|<p>Test Body id</p>|f537448b-d498-4fed-8090-e5cd62f363ff.html|4fa75f64-5777-4562-b3fc-2c963f66ab00",
	}

	for i := range items {
		values := strings.Split(items[i], "|")
		data = append(data, EmailTemplateTranslation{
			Base: basic.Base{
				ID: lib.StringToUUID(values[0]),
			},
			EmailTemplateID: lib.StringToUUID(values[6]),
			EmailTemplateTranslationAPI: EmailTemplateTranslationAPI{
				LanguageCode:      lib.Strptr(values[1]),
				Subject:           lib.Strptr(values[2]),
				EmailTemplateName: lib.Strptr(values[3]),
				Body:              lib.Strptr(values[4]),
				BodyFileName:      lib.Strptr(values[5]),
			},
		})
	}

	return &data
}
