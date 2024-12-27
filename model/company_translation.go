package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CompanyTranslation Company Translation
type CompanyTranslation struct {
	basic.Base
	basic.DataOwner
	CompanyTranslationAPI
	CompanyID *uuid.UUID `json:"company_id,omitempty" gorm:"type:varchar(36);uniqueIndex:company_translation_unique;not null" swaggertype:"string" format:"uuid"` // Company id
}

// CompanyTranslationAPI Company Translation API
type CompanyTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:company_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CompanyName  *string `json:"company_name,omitempty" gorm:"type:varchar(256)"`                                                             // Company Name
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                      // Description
}
