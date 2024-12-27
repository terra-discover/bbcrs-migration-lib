package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentDestinationGroupDestination Model
type AgentDestinationGroupDestination struct {
	basic.Base
	basic.DataOwner
	AgentDestinationGroupDestinationAPI
	City          *City          `json:"city" gorm:"foreignKey:CityID;references:ID"`
	Country       *Country       `json:"country" gorm:"foreignKey:CountryID;references:ID"`
	StateProvince *StateProvince `json:"state_province" gorm:"foreignKey:StateProvinceID;references:ID"`
	Destination   *Destination   `json:"destination" gorm:"foreignKey:DestinationID;references:ID"`
}

// AgentDestinationGroupDestinationAPI API
type AgentDestinationGroupDestinationAPI struct {
	AgentDestinationGroupID *uuid.UUID `json:"agent_destination_group_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	CityID                  *uuid.UUID `json:"city_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	CountryID               *uuid.UUID `json:"country_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	StateProvinceID         *uuid.UUID `json:"state_province_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	DestinationID           *uuid.UUID `json:"destination_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	IsPrimary               *bool      `json:"is_primary,omitempty" gorm:"type:bool"`
}
