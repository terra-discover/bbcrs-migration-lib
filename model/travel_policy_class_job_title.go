package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyClassJobTitle Travel Policy Class Job Title
type TravelPolicyClassJobTitle struct {
	basic.Base
	basic.DataOwner
	TravelPolicyClassJobTitleAPI
	TravelPolicyClass *TravelPolicyClass `json:"travel_policy_class,omitempty" gorm:"foreignKey:TravelPolicyClassID;references:ID"` // travel policy class
	JobTitle          *JobTitle          `json:"job_title,omitempty" gorm:"foreignKey:JobTitleID;reference:ID"`                     // job title
}

// TravelPolicyClassJobTitleAPI Travel Policy Class Job Title API
type TravelPolicyClassJobTitleAPI struct {
	TravelPolicyClassID *uuid.UUID `json:"travel_policy_class_id,omitempty" gorm:"type:varchar(36);index:idx_travel_policy_class_job_title__job_title_id,unique,where:deleted_at is null;not null;" format:"uuid"` // travel policy class id
	JobTitleID          *uuid.UUID `json:"job_title_id,omitempty" gorm:"type:varchar(36);index:idx_travel_policy_class_job_title__job_title_id,unique,where:deleted_at is null;not null;" format:"uuid"`           // job title id
}
