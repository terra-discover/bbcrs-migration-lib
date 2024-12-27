package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerRoomAmenityType Integration Partner Room Amenity Type
type IntegrationPartnerRoomAmenityType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerRoomAmenityTypeAPI
	IntegrationPartnerRoomAmenityTypeTranslation *IntegrationPartnerRoomAmenityTypeTranslation `json:"integration_partner_room_amenity_type_translation,omitempty"` // Integration Partner Room Amenity Type Translation
}

// IntegrationPartnerRoomAmenityTypeAPI Integration Partner Room Amenity Type API
type IntegrationPartnerRoomAmenityTypeAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_room_amenity_type__room_amenity_type_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	RoomAmenityTypeID    *uuid.UUID `json:"room_amenity_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_room_amenity_type__room_amenity_type_code,unique,where:deleted_at is null;not null"`   // Room Amenity Type ID
	RoomAmenityTypeCode  *string    `json:"room_amenity_type_code,omitempty" gorm:"type:varchar(8);index:idx_integration_partner_room_amenity_type__room_amenity_type_code,unique,where:deleted_at is null;not null"`                                     // Room Amenity Type Code
	RoomAmenityTypeName  *string    `json:"room_amenity_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                                    // Room Amenity Type Name
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                                       // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                                           // Comment
}
