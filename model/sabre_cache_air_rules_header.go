package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreCacheAirRulesHeader struct {
	basic.Base
	basic.DataOwner
	SessionID            *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	SabreCacheAirRulesID *uuid.UUID `json:"sabre_cache_air_rules_id" gorm:"type:varchar(36)"`
	Type                 *string    `json:"type" gorm:"type:text"`
	Text                 *string    `json:"text" gorm:"type:text"`
}
