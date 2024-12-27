package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CacheRoomStayCriteria Cache Room Stay Criteria
type CacheRoomStayCriteria struct {
	basic.Base
	basic.DataOwner
	CacheRoomStayCriteriaAPI
	CacheCriteria    *CacheCriteria `json:"cache_criteria,omitempty" gorm:"foreignKey:CacheCriteriaID;references:ID" swaggerignore:"true"` // cache criteria
	DestinationCity  *City          `json:"destination_city,omitempty" gorm:"foreignKey:DestinationCityID;references:ID"`                  // destination city
	DestinationHotel *Hotel         `json:"destination_hotel,omitempty" gorm:"foreignKey:DestinationHotelID;references:ID"`                // destination hotel
}

// CacheRoomStayCriteriaAPI Cache Room Stay Criteria API
type CacheRoomStayCriteriaAPI struct {
	CacheCriteriaID      *uuid.UUID       `json:"cache_criteria_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggerignore:"true"` // cache criteria id
	IndexNumber          *int             `json:"index_number,omitempty" gorm:"type:smallint;" example:"1"`                                        // index number
	DestinationLocation  *string          `json:"destination_location,omitempty" gorm:"type:varchar(256)"`                                         // destination location (Mengikuti pilihan dari FE bisa hotel bisa city)
	DestinationCityID    *uuid.UUID       `json:"destination_city_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                             // destination city id
	DestinationHotelID   *uuid.UUID       `json:"destination_hotel_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                            // destination hotel id
	CityID               *uuid.UUID       `json:"city_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`                                   // destination city id
	AirportID            *uuid.UUID       `json:"airport_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`                                // destination city id
	AdministrativeCityID *uuid.UUID       `json:"administrative_city_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`                    // destination city id
	ZoneID               *uuid.UUID       `json:"zone_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`                                   // destination city id
	DestinationID        *uuid.UUID       `json:"destination_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`                            // destination city id
	HotelID              *uuid.UUID       `json:"hotel_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`                                  // destination city id
	AttractionID         *uuid.UUID       `json:"attraction_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`                             // destination city id
	CheckIn              *strfmt.DateTime `json:"check_in,omitempty" gorm:"type:timestamptz;null" swaggertype:"string" format:"date-time"`         // checkin
	CheckOut             *strfmt.DateTime `json:"check_out,omitempty" gorm:"type:timestamptz;null" swaggertype:"string" format:"date-time"`        // checkout
}
