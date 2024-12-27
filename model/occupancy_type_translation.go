package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OccupancyTypeTranslation Occupancy Type Translation
type OccupancyTypeTranslation struct {
	basic.Base
	basic.DataOwner
	OccupancyTypeTranslationAPI
	OccupancyTypeID *uuid.UUID `json:"occupancy_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:occupancy_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Occupancy Type id
}

// OccupancyTypeTranslationAPI Occupancy Type Translation API
type OccupancyTypeTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:occupancy_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	OccupancyTypeName *string `json:"occupancy_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Occupancy Type Name
}
