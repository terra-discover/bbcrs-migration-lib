package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelPolicyTranslation Hotel Policy Translation
type HotelPolicyTranslation struct {
	basic.Base
	basic.DataOwner
	HotelPolicyTranslationAPI
	HotelPolicyID *uuid.UUID `json:"hotel_policy_id,omitempty" gorm:"type:varchar(36);uniqueIndex:hotel_policy_translation_unique;not null" swaggertype:"string" format:"uuid"` // Hotel Policy id
}

// HotelPolicyTranslationAPI Hotel Policy Translation API
type HotelPolicyTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:hotel_policy_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CheckInTime  *string `json:"check_in_time,omitempty" gorm:"type:varchar(16)"`                                                                  // Check In Time
	CheckOutTime *string `json:"check_out_time,omitempty" gorm:"type:varchar(16)"`                                                                 // Check Out Time
}
