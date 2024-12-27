package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SabreCacheAirRulesStatus model
type SabreCacheAirRulesStatus struct {
	basic.Base
	basic.DataOwner
	SessionID            *uuid.UUID       `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	SabreCacheAirRulesID *uuid.UUID       `json:"sabre_cache_air_rules_id" gorm:"type:varchar(36)"`
	Status               *strfmt.DateTime `json:"status" gorm:"type:varchar(60)"`
	Timestamp            *string          `json:"timestamp" gorm:"type:timestamp"`
	// HostCommand          *[]SabreStatusMessage `json:"host_command" gorm:"foreignKey:SabreCacheAirRulesStatusID;references:ID"`
	// Message              *[]SabreStatusMessage `json:"message" gorm:"foreignKey:SabreCacheAirRulesStatusID;references:ID"`
}

// SabreStatusMessage struct
// type SabreStatusMessage struct {
// 	SessionID                  *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
// 	SabreCacheAirRulesStatusID *uuid.UUID `json:"sabre_cache_air_rules_status_id" gorm:"type:varchar(36)"`
// 	LNIATA                     *string    `json:"lniata" gorm:"type:varchar(60)"`
// 	Text                       *string    `json:"text" gorm:"type:text"`
// }
