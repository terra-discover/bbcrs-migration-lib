package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyDestinationGroup Travel Policy Job Title
type TravelPolicyDestinationGroup struct {
	basic.Base
	basic.DataOwner
	TravelPolicyDestinationGroupAPI
	TravelPolicy     *TravelPolicy     `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"`        // travel policy
	DestinationGroup *DestinationGroup `json:"destination_group,omitempty" gorm:"foreignKey:DestinationGroupID;reference:ID"` // destination group
}

// TravelPolicyDestinationGroupAPI Travel Policy Job Title API
type TravelPolicyDestinationGroupAPI struct {
	TravelPolicyID     *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`     // travel policy id
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // destination group id
}
