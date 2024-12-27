package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelAmenityType Integration Partner Hotel Amenity Type
type IntegrationPartnerHotelAmenityType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelAmenityTypeAPI
	IntegrationPartnerHotelAmenityTypeTranslation *IntegrationPartnerHotelAmenityTypeTranslation `json:"integration_partner_hotel_amenity_type_translation,omitempty"` // Integration Partner Hotel Amenity Type Translation
	HotelAmenityType                              *HotelAmenityType                              `json:"hotel_amenity_type,omitempty"`                                 // HotelAmenityType
}

// IntegrationPartnerHotelAmenityTypeAPI Integration Partner Hotel Amenity Type API
type IntegrationPartnerHotelAmenityTypeAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_idx_integration_partner_hotel_amenity_type__hotel_amenity_type_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	HotelAmenityTypeID   *uuid.UUID `json:"hotel_amenity_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_idx_integration_partner_hotel_amenity_type__hotel_amenity_type_code,unique,where:deleted_at is null;not null"`  // Hotel Amenity Type ID
	HotelAmenityTypeCode *string    `json:"hotel_amenity_type_code,omitempty" gorm:"type:varchar(8);index:idx_idx_integration_partner_hotel_amenity_type__hotel_amenity_type_code,unique,where:deleted_at is null;not null"`                                    // Hotel Amenity Type Code
	HotelAmenityTypeName *string    `json:"hotel_amenity_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                                         // Hotel Amenity Type Name
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                                             // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                                                 // Comment
}
