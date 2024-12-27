package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyPropertyCategory Travel Policy Job Title
type TravelPolicyPropertyCategory struct {
	basic.Base
	basic.DataOwner
	TravelPolicyPropertyCategoryAPI
	TravelPolicy     *TravelPolicy     `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"`        // travel policy
	PropertyCategory *PropertyCategory `json:"property_category,omitempty" gorm:"foreignKey:PropertyCategoryID;reference:ID"` // property category
}

// TravelPolicyPropertyCategoryAPI Travel Policy Job Title API
type TravelPolicyPropertyCategoryAPI struct {
	TravelPolicyID     *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`     // travel policy id
	PropertyCategoryID *uuid.UUID `json:"property_category_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // property category id
	IsRestricted       *bool      `json:"is_restricted,omitempty" example:"true"`
}
