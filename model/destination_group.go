package model

import (
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DestinationGroup Destination Group
type DestinationGroup struct {
	basic.Base
	basic.DataOwner
	DestinationGroupAPI
	DestinationGroupTranslation *DestinationGroupTranslation `json:"destination_group_translation,omitempty"` // Destrination Group Translation
	DestinationGroupCountry     []DestinationGroupCountry    `json:"destination_group_country,omitempty"`
	AgentDestinationGroup       []AgentDestinationGroup      `json:"agent_destination_group,omitempty"`
}

// DestinationGroupAPI Destination Group API
type DestinationGroupAPI struct {
	DestinationGroupCode *string `json:"destination_group_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" example:"domestic"`      // Indicates domestic destination
	DestinationGroupName *string `json:"destination_group_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"pulau lombok"` // Destination Group Name
}

func (dg DestinationGroup) Seed() *[]DestinationGroup {
	data := []DestinationGroup{}
	items := []string{
		"domestic|Domestic",
		"international|International",
	}

	for i := range items {
		values := strings.Split(items[i], "|")
		var code string = values[0]
		var name string = values[1]
		data = append(data, DestinationGroup{
			DestinationGroupAPI: DestinationGroupAPI{
				DestinationGroupCode: &code,
				DestinationGroupName: &name,
			},
		})
	}

	return &data
}
