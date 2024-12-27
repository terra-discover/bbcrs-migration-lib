package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicy Travel Policy Class
type TravelPolicy struct {
	basic.Base
	basic.DataOwner
	TravelPolicyAPI
	TravelPolicyClass            *TravelPolicyClass             `json:"travel_policy_class,omitempty"`      // multimedia description
	JobTitles                    []JobTitle                     `gorm:"many2many:travel_policy_job_title;"` // job titles
	TravelPolicyProductTypes     []TravelPolicyProductType      `json:"travel_policy_product_types,omitempty"`
	TravelPolicyDestinationGroup []TravelPolicyDestinationGroup `json:"travel_policy_destination_groups,omitempty"`
	TravelPolicyAirline          []TravelPolicyAirline          `json:"travel_policy_airline,omitempty"`
	TravelPolicyAirlineCategory  []TravelPolicyAirlineCategory  `json:"travel_policy_airline_category,omitempty"`
	TravelPolicyCabinType        []TravelPolicyCabinType        `json:"travel_policy_cabin_type,omitempty"`
}

// TravelPolicyAPI Travel Policy Class API
type TravelPolicyAPI struct {
	TravelPolicyClassID         *uuid.UUID  `json:"travel_policy_class_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	TravelPolicyCode            *string     `json:"travel_policy_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null"` // travel policy code
	TravelPolicyName            *string     `json:"travel_policy_name,omitempty" gorm:"type:varchar(36)"`                                        // travel policy name
	ApplicationOrder            *int        `json:"application_order,omitempty"`                                                                 // travel policy name
	AssignAllProductTypes       *bool       `json:"assign_all_product_types,omitempty"`
	AssignAllDestinationGroups  *bool       `json:"assign_all_destination_groups,omitempty"`
	AssignAllAirlineCategories  *bool       `json:"assign_all_airline_categories,omitempty"`
	AssignAllAirlines           *bool       `json:"assign_all_airlines,omitempty"`
	AssignAllCabinTypes         *bool       `json:"assign_all_cabin_types,omitempty"`
	AssignAllPropertyCategories *bool       `json:"assign_all_property_categories,omitempty"`
	AssignAllRatingTypes        *bool       `json:"assign_all_rating_types,omitempty"`
	Comment                     *string     `json:"comment,omitempty" gorm:"type:text"`     // Comment
	Description                 *string     `json:"description,omitempty" gorm:"type:text"` // Description
	JobTitles                   []*JobTitle `json:"job_titles,omitempty" gorm:"-"`
}
