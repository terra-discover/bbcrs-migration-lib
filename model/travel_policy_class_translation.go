package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyClassTranslation Travel Policy Class Translation
type TravelPolicyClassTranslation struct {
	basic.Base
	basic.DataOwner
	TravelPolicyClassTranslationAPI
	TravelPolicyClass *TravelPolicyClass `json:"travel_policy_class,omitempty" gorm:"foreignKey:TravelPolicyClassID;references:ID"` // travel policy class
}

// TravelPolicyClassTranslationAPI Travel Policy Class Translation API
type TravelPolicyClassTranslationAPI struct {
	TravelPolicyClassID   *uuid.UUID `json:"travel_policy_class_id,omitempty" gorm:"type:varchar(36);uniqueIndex:travel_policy_class_translation_unique;not null" swaggertype:"string" swaggerignore:"true" format:"uuid"` // Travel Policy Class id
	LanguageCode          *string    `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:travel_policy_class_translation_unique;not null" example:"en"`                                                      // Language code example: en, id, cn etc...
	TravelPolicyClassName *string    `json:"travel_policy_class_name,omitempty" gorm:"type:varchar(256);"`                                                                                                                 // travel policy class name
}
