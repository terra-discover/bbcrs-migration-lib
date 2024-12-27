package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateTravelPolicyClass Corporate Travel Policy Class
type CorporateTravelPolicyClass struct {
	basic.Base
	basic.DataOwner
	CorporateTravelPolicyClassAPI
	Corporate         *Corporate         `json:"corporate,omitempty" gorm:"foreignKey:CorporateID;references:ID"`                   // corporate
	TravelPolicyClass *TravelPolicyClass `json:"travel_policy_class,omitempty" gorm:"foreignKey:TravelPolicyClassID;references:ID"` // travel policy class
}

// CorporateTravelPolicyClassAPI Corporate Travel Policy Class API
type CorporateTravelPolicyClassAPI struct {
	CorporateID         *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_travel_policy_class__travel_policy_class_id,unique,where:deleted_at is null;not null;" format:"uuid"`           // corporate id
	TravelPolicyClassID *uuid.UUID `json:"travel_policy_class_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_travel_policy_class__travel_policy_class_id,unique,where:deleted_at is null;not null;" format:"uuid"` // travel policy class id
}
