package model

import (
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyClass Travel Policy Class
type TravelPolicyClass struct {
	basic.Base
	basic.DataOwner
	TravelPolicyClassAPI
	CorporateTravelPolicyClass *CorporateTravelPolicyClass `json:"corporate_travel_policy_class,omitempty" swaggerignore:"true"`
	TravelPolicyClassJobTitle  []TravelPolicyClassJobTitle `json:"-" swaggerignore:"true"`
	TravelPolicy               []TravelPolicy              `json:"travel_policy,omitempty"` // multimedia description
}

// TravelPolicyClassAPI Travel Policy Class API
type TravelPolicyClassAPI struct {
	TravelPolicyClassCode *string `json:"travel_policy_class_code,omitempty" gorm:"type:varchar(36);index;not null"` // travel policy class code
	TravelPolicyClassName *string `json:"travel_policy_class_name,omitempty" gorm:"type:varchar(256);not null"`      // travel policy class name
}
