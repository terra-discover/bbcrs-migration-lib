package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TeamTranslation Team Translation
type TeamTranslation struct {
	basic.Base
	basic.DataOwner
	TeamTranslationAPI
	TeamID *uuid.UUID `json:"team_id,omitempty" gorm:"type:varchar(36);uniqueIndex:team_translation_unique;not null" swaggertype:"string" format:"uuid"` // Team id
}

// TeamTranslationAPI Team Translation API
type TeamTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:team_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	TeamName     *string `json:"team_name,omitempty" gorm:"type:varchar(256)" example:"Sales Team 1"`                                      // Team Name
}
