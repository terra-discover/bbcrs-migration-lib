package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerRoomAmenityTypeTranslation Integration Partner Room Amenity Type Translation
type IntegrationPartnerRoomAmenityTypeTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerRoomAmenityTypeTranslationAPI
	IntegrationPartnerRoomAmenityTypeID *uuid.UUID `json:"integration_partner_room_amenity_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_room_amenity_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Room Amenity Type id
}

// IntegrationPartnerRoomAmenityTypeTranslationAPI Integration Partner Room Amenity Type Translation API
type IntegrationPartnerRoomAmenityTypeTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_room_amenity_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoomAmenityTypeName *string `json:"room_amenity_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Room Amenity Type Name
	Description         *string `json:"description,omitempty" gorm:"type:text"`                                                                                                    // Description
	Comment             *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                        // Comment
}
