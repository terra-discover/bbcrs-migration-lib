package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerLanguage Integration Partner Language
type IntegrationPartnerLanguage struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerLanguageAPI
}

// IntegrationPartnerLanguageAPI Integration Partner Language API
type IntegrationPartnerLanguageAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;"`    // Integration Partner ID
	LanguageID           *uuid.UUID `json:"language_id,omitempty" gorm:"type:varchar(36);not null;"`               // Language ID
	LanguageCode         *string    `json:"language_code,omitempty" gorm:"type:varchar(8);not null" example:"bhs"` // Language Code
	LanguageName         *string    `json:"language_name,omitempty" gorm:"type:varchar(256)" example:"bahasa"`     // Language Name
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                    // Comment
}
