package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BasisTypeTranslation Basis Type Translation
type BasisTypeTranslation struct {
	basic.Base
	basic.DataOwner
	BasisTypeTranslationAPI
	BasisTypeID *uuid.UUID `json:"basis_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:basis_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Basis Type id
}

// BasisTypeTranslationAPI Basis Type Translation API
type BasisTypeTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:basis_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	BasisTypeName *string `json:"basis_type_name,omitempty" gorm:"type:varchar(256)" example:"Daily"`                                             // Basis Type Name
}
