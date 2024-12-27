package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Viewership Viewership
type Viewership struct {
	basic.Base
	basic.DataOwner
	ViewershipAPI
}

// ViewershipAPI Viewership Api
type ViewershipAPI struct {
	ViewershipCode                 *uuid.UUID `json:"viewership_code,omitempty" gorm:"type:varchar(36);not null index:,unique,where:deleted_at is null"` //Viewership Code
	ViewershipName                 *string    `json:"viewership_name,omitempty" gorm:"type:varchar(64);not null,where:deleted_at is null"`               //Viewership Name
	AssignAllBusinessPartnerType   *int       `json:"assign_all_business_partner_type,omitempty" gorm:"type:smallint;"`                                  // Assign All Business Partner Type
	AssignAllBusinessPartnerGroups *int       `json:"assign_all_business_partner_group,omitempty" gorm:"type:smallint;"`                                 // Assign All Business Partner Groups
	AssignAllBusinessPartners      *int       `json:"assign_all_business_partners,omitempty" gorm:"type:smallint;"`                                      // Assign All Business Partners
	AssignAllDestinationSystem     *int       `json:"assign_all_destination_system,omitempty" gorm:"type:smallint;"`                                     // Assign All Destination Systems
	AssignAllLocation              *int       `json:"assign_all_location,omitempty" gorm:"type:smallint;"`                                               // Assign All Location
	AssignAllDevice                *int       `json:"assign_all_device,omitempty" gorm:"type:smallint;"`                                                 // Assign All Device
	AssignAllLoyalty               *int       `json:"assign_all_Lloyalty,omitempty" gorm:"type:smallint;"`                                               // Assign All Loyalty
	IsAccessCodeRequired           *int       `json:"is_access_code_required,omitempty" gorm:"type:smallint;"`                                           // Is Access Code Required
	Comment                        *string    `json:"comment,omitempty" gorm:"type:text;"`                                                               // Comment
	Description                    *string    `json:"description,omitempty" gorm:"type:text;"`                                                           // Description
}
