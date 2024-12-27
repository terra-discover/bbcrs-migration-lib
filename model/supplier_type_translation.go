package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SupplierTypeTranslation Supplier Type Translation
type SupplierTypeTranslation struct {
	basic.Base
	basic.DataOwner
	SupplierTypeTranslationAPI
	SupplierTypeID *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:supplier_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Supplier Type id
}

// SupplierTypeTranslationAPI Supplier Type Translation API
type SupplierTypeTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:supplier_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	SupplierTypeName *string `json:"supplier_type_name,omitempty" gorm:"type:varchar(256)" example:"Direct"`                                            // Supplier Type Name
}
