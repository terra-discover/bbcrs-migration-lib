package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotel Integration Partner Hotel
type IntegrationPartnerHotel struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"index:idx_idx_integration_partner_hotel__hotel_name,unique,where:deleted_at is null;not null"` // Integration Partner ID
	IntegrationPartnerHotelAPI
	IntegrationPartnerHotelTranslation *IntegrationPartnerHotelTranslation `json:"integration_partner_hotel_translation,omitempty"` // Integration Partner Hotel Translation
	Hotel                              *Hotel                              `json:"hotel,omitempty"`                                 // Hotel
	Currency                           *Currency                           `json:"currency,omitempty"`                              // Currency
}

// IntegrationPartnerHotelAPI Integration Partner Hotel API
type IntegrationPartnerHotelAPI struct {
	HotelID                      *uuid.UUID `json:"hotel_id,omitempty" swaggertype:"string" format:"uuid" gorm:"index:idx_idx_integration_partner_hotel__hotel_name,unique,where:deleted_at is null;not null"` // Hotel ID
	WebServiceUserName           *string    `json:"web_service_user_name,omitempty" gorm:"type:varchar(256)"`                                                                                                  // Web Service User Name
	WebServicePassword           *string    `json:"web_service_password,omitempty" gorm:"type:varchar(256)"`                                                                                                   // Web Service Password
	HotelCode                    *string    `json:"hotel_code,omitempty" gorm:"type:varchar(36)"`                                                                                                              // Hotel Code
	HotelName                    *string    `json:"hotel_name,omitempty" gorm:"type:varchar(128);index:idx_idx_integration_partner_hotel__hotel_name,unique,where:deleted_at is null;not null"`                // Hotel Name
	CurrencyID                   *uuid.UUID `json:"currency_id,omitempty" swaggertype:"string" format:"uuid"`                                                                                                  // Currency ID
	ExtranetURL                  *string    `json:"extranet_url,omitempty" gorm:"type:varchar(256)"`                                                                                                           // Extranet Url
	ExtranetUserName             *string    `json:"extranet_user_name,omitempty" gorm:"type:varchar(256)"`                                                                                                     // Extranet User Name
	ExtranetPassword             *string    `json:"extranet_password,omitempty" gorm:"type:varchar(256)"`                                                                                                      // Extranet Password
	RoomRateEnabled              *bool      `json:"room_rate_enabled,omitempty"`                                                                                                                               // Room Rate Enabled
	RoomAvailabilityEnabled      *bool      `json:"room_availability_enabled,omitempty"`                                                                                                                       // Room Availability Enabled
	ClosedOnArrivalEnabled       *bool      `json:"closed_on_arrival_enabled,omitempty"`                                                                                                                       // Closed On Arrival Enabled
	ClosedOnDepartureEnabled     *bool      `json:"closed_on_departure_enabled,omitempty"`                                                                                                                     // Closed On Departure Enabled
	LengthOfStayEnabled          *bool      `json:"length_of_stay_enabled,omitempty"`                                                                                                                          // Length Of Stay Enabled
	BookingAdvancedOffsetEnabled *bool      `json:"booking_advanced_offset_enabled,omitempty"`                                                                                                                 // Booking Advanced Offset Enabled
	AvailabilityStatusEnabled    *bool      `json:"availability_status_enabled,omitempty"`                                                                                                                     // Availability Status Enabled
	ReservationEnabled           *bool      `json:"reservation_enabled,omitempty"`                                                                                                                             // Reservation Enabled
	HotelCodeEnabled             *bool      `json:"hotel_code_enabled,omitempty"`                                                                                                                              // Hotel Code Enabled
	HotelInformationEnabled      *bool      `json:"hotel_information_enabled,omitempty"`                                                                                                                       // Hotel Information Enabled
	IsEnabled                    *bool      `json:"is_enabled,omitempty"`                                                                                                                                      // Is Enabled
	IsPrimary                    *bool      `json:"is_primary,omitempty"`                                                                                                                                      // Is Primary
	HotelDetails                 *string    `json:"hotel_details,omitempty" gorm:"type:text"`                                                                                                                  // Hotel Details
	Comment                      *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                        // Comment
}

// IntegrationPartnerHotelDetail Integration Partner Hotel Detail
type IntegrationPartnerHotelDetail struct {
	IntegrationPartnerHotel
	Address *Address `json:"address,omitempty" gorm:"-"` // Address
}
