package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ConvenienceStoreTranslation model
type ConvenienceStoreTranslation struct {
	basic.Base
	basic.DataOwner
	ConvenienceStoreTranslationAPI
	ConvenienceStoreID *uuid.UUID `json:"convenience_store_id,omitempty" gorm:"type:varchar(36);uniqueIndex:convenience_store_translation_unique;not null" swaggertype:"string" format:"uuid"` // Convenience Store id
}

// ConvenienceStoreTranslationAPI model
type ConvenienceStoreTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:convenience_store_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ConvenienceStoreName *string `json:"convenience_store_name,omitempty" gorm:"type:varchar(256)"`                                                             // Convenience Store Name
}
