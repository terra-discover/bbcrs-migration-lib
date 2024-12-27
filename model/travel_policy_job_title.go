package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyJobTitle Travel Policy Job Title
type TravelPolicyJobTitle struct {
	basic.Base
	basic.DataOwner
	TravelPolicyJobTitleAPI
	TravelPolicy *TravelPolicy `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"` // travel policy
	JobTitle     *JobTitle     `json:"job_title,omitempty" gorm:"foreignKey:JobTitleID;reference:ID"`          // job title
}

// TravelPolicyJobTitleAPI Travel Policy Job Title API
type TravelPolicyJobTitleAPI struct {
	TravelPolicyID *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // travel policy id
	JobTitleID     *uuid.UUID `json:"job_title_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`     // job title id
}
