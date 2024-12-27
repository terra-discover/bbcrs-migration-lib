package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelRoomTypeTranslation Integration Partner Hotel Room Type Translation
type IntegrationPartnerHotelRoomTypeTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelRoomTypeTranslationAPI
	IntegrationPartnerHotelRoomTypeID *uuid.UUID `json:"integration_partner_hotel_room_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_hotel_room_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Hotel Room Type id
}

// IntegrationPartnerHotelRoomTypeTranslationAPI Integration Partner Hotel Room Type Translation API
type IntegrationPartnerHotelRoomTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_hotel_room_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoomTypeName *string `json:"room_type_name,omitempty" gorm:"type:varchar(512)"`                                                                                       // Room Type Name
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                                                  // Description
	Comment      *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                      // Comment
}
