package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelRatingType Integration Partner Hotel Rating Type
type IntegrationPartnerHotelRatingType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelRatingTypeAPI
	IntegrationPartnerHotelRatingTypeTranslation *IntegrationPartnerHotelRatingTypeTranslation `json:"integration_partner_hotel_rating_type_translation,omitempty"` // Integration Partner Hotel Rating Type Translation
}

// IntegrationPartnerHotelRatingTypeAPI Integration Partner Hotel Rating Type API
type IntegrationPartnerHotelRatingTypeAPI struct {
	IntegrationPartnerID      *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_hotel_rating_type__hotel_rating_type_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	HotelRatingTypeID         *uuid.UUID `json:"hotel_rating_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_hotel_rating_type__hotel_rating_type_code,unique,where:deleted_at is null;not null"`   // Hotel Rating Type ID
	HotelRatingTypeCode       *string    `json:"hotel_rating_type_code,omitempty" gorm:"type:varchar(8);index:idx_integration_partner_hotel_rating_type__hotel_rating_type_code,unique,where:deleted_at is null;not null"`                                     // Hotel Rating Type Code
	HotelRatingTypeName       *string    `json:"hotel_rating_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                                    // Hotel Rating Type Name
	HotelRatingTypeSimpleCode *string    `json:"hotel_rating_type_simple_code,omitempty" gorm:"type:varchar(8);index:,unique,where:deleted_at is null;not null"`                                                                                               // Hotel Rating Type Simple Code
	AccommodationTypeName     *string    `json:"accommodation_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                                   // Accommodation Type Name
	Description               *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                                       // Description
	Comment                   *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                                           // Comment
}
