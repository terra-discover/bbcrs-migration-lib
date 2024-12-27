package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelFacility Hotel Facility
type HotelFacility struct {
	basic.Base
	basic.DataOwner
	HotelFacilityAPI
	HotelFacilityTranslation *HotelFacilityTranslation `json:"hotel_facility_translation,omitempty"` // Hotel Facility Translation
	Proximity                *Proximity                `json:"proximity,omitempty"`                  // Proximity
	OperationSchedule        *OperationSchedule        `json:"operation_schedule,omitempty"`         // Operation Schedule
	HotelFacilityAsset       *HotelFacilityAsset       `json:"hotel_facility_asset,omitempty"`       // Hotel Facility Asset
}

// HotelFacilityAPI Hotel Facility API
type HotelFacilityAPI struct {
	HotelID               *uuid.UUID             `json:"hotel_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`                 // Hotel ID
	HotelFacilityName     *string                `json:"hotel_facility_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Hotel Facility Name
	IsAvailableToAnyGuest *bool                  `json:"is_available_to_any_guest,omitempty"`                                                                    // Is Available To Any Guest
	ProximityID           *uuid.UUID             `json:"proximity_id,omitempty" swaggertype:"string" format:"uuid"`                                              // Proximity ID
	OperationScheduleID   *uuid.UUID             `json:"operation_schedule_id,omitempty" swaggertype:"string" format:"uuid"`                                     // Operation Schedule ID
	RelativePositionID    *uuid.UUID             `json:"relative_position_id,omitempty" swaggertype:"string" format:"uuid"`                                      // Relative Position ID
	Description           *string                `json:"description,omitempty" gorm:"type:text"`                                                                 // Description
	HotelFacilityAsset    *HotelFacilityAssetAPI `json:"hotel_facility_asset,omitempty" gorm:"-"`                                                                // Hotel Facility Asset
}
