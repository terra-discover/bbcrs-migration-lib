package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirTravelPreferenceCriteria struct {
	basic.Base
	basic.DataOwner
	AirTravelPreferenceCriteriaAPI
}

type AirTravelPreferenceCriteriaAPI struct {
	ProposalID *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	CabinCode  *string    `json:"cabin_code,omitempty" gorm:"type:varchar(2)"`
}
