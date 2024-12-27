package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TimeZoneTranslation Timezone Translation
type TimeZoneTranslation struct {
	basic.Base
	basic.DataOwner
	TimeZoneTranslationAPI
	TimeZoneID *uuid.UUID `json:"time_zone_id,omitempty" gorm:"type:varchar(36);index:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Time zone id
}

// TimeZoneTranslationAPI Timezone Translation API
type TimeZoneTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);index:,unique,where:deleted_at is null;not null" example:"en"` // Language code example: en, id, cn etc...
	ZoneName     *string `json:"zone_name,omitempty" gorm:"type:varchar(64)" example:"Bangkok, Jakarta"`                                      // Zone Name
	Description  *string `json:"description,omitempty" gorm:"type:text" example:"Bangkok, Jakarta"`                                           // Description
}
