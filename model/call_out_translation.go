package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CallOutTranslation Call Out Translation
type CallOutTranslation struct {
	basic.Base
	basic.DataOwner
	CallOutTranslationAPI
	CallOutID *uuid.UUID `json:"call_out_id,omitempty" gorm:"type:varchar(36);uniqueIndex:call_out_translation_unique;not null" swaggertype:"string" format:"uuid"` // Call Out id
}

// CallOutTranslationAPI Call Out Translation API
type CallOutTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:call_out_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CallOutName  *string `json:"call_out_name,omitempty" gorm:"type:varchar(256)"`                                                             // Call Out Name 	// Background Color
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                       // Description
}
