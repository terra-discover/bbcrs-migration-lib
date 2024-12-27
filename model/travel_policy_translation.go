package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyTranslation Travel Policy Translation
type TravelPolicyTranslation struct {
	basic.Base
	basic.DataOwner
	TravelPolicyTranslationAPI
	TravelPolicy *TravelPolicy `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"` // travel policy
}

// TravelPolicyTranslationAPI Travel Policy Translation API
type TravelPolicyTranslationAPI struct {
	TravelPolicyID   *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);uniqueIndex:travel_policy_translation_unique;not null" swaggertype:"string" swaggerignore:"true" format:"uuid"` // Travel Policy id
	LanguageCode     *string    `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:travel_policy_translation_unique;not null" example:"en"`                                                // Language code example: en, id, cn etc...
	TravelPolicyName *string    `json:"travel_policy_name,omitempty" gorm:"type:varchar(256);"`                                                                                                           // travel policy name
}
