package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelPolicy Hotel Policy
type HotelPolicy struct {
	basic.Base
	basic.DataOwner
	HotelPolicyAPI
	HotelPolicyTranslation *HotelPolicyTranslation `json:"hotel_policy_translation,omitempty"` // Hotel Policy Translation
}

// HotelPolicyAPI Hotel Policy API
type HotelPolicyAPI struct {
	HotelID         *uuid.UUID `json:"hotel_id,omitempty" swaggertype:"string" format:"uuid" gorm:"index:,unique,where:deleted_at is null;not null"` // Hotel ID
	CheckInTime     *string    `json:"check_in_time,omitempty" gorm:"type:varchar(16)"`                                                              // Check In Time
	CheckOutTime    *string    `json:"check_out_time,omitempty" gorm:"type:varchar(16)"`                                                             // Check Out Time
	TotalAdultCount *int       `json:"total_adult_count,omitempty" gorm:"type:smallint"`                                                             // Total Adult Count
	TotalChildCount *int       `json:"total_child_count,omitempty" gorm:"type:smallint"`                                                             // Total Child Count
	MaxChildAge     *int       `json:"max_child_age,omitempty" gorm:"type:smallint"`                                                                 // Max Child Age
}
