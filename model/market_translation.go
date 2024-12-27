package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarketTranslation Market Translation
type MarketTranslation struct {
	basic.Base
	basic.DataOwner
	MarketTranslationAPI
	MarketID *uuid.UUID `json:"market_id,omitempty" gorm:"type:varchar(36);uniqueIndex:market_translation_unique;not null" swaggertype:"string" format:"uuid"` // Market id
}

// MarketTranslationAPI Market Translation API
type MarketTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:market_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MarketName   *string `json:"market_name,omitempty" gorm:"type:varchar(256);not null" example:"Corporate"`                                // Market Name
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                     // Description
}
