package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// TravelPurpose model
type TravelPurpose struct {
	basic.Base
	basic.DataOwner
	TravelPurposeAPI
	TravelPurposeTranslation *TravelPurposeTranslation `json:"travel_purpose_translation,omitempty"`
}

// TravelPurposeAPI model
type TravelPurposeAPI struct {
	TravelPurposeCode *string `json:"travel_purpose_code,omitempty" gorm:"type:varchar(36);where:deleted_at is null;not null"`
	TravelPurposeName *string `json:"travel_purpose_name,omitempty" gorm:"type:varchar(256);not null;where:deleted_at is null"`
}
