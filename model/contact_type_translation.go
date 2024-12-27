package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ContactTypeTranslation ContactType Translation
type ContactTypeTranslation struct {
	basic.Base
	basic.DataOwner
	ContactTypeTranslationAPI
	ContactTypeID *uuid.UUID `json:"contact_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:contact_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Contact Type id
}

// ContactTypeTranslationAPI ContactType Translation API
type ContactTypeTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:contact_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ContactTypeName *string `json:"contact_type_name,omitempty" gorm:"type:varchar(256);"`                                                            // Contact Type Name
}
