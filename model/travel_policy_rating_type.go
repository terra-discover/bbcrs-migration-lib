package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyRatingType Travel Policy Job Title
type TravelPolicyRatingType struct {
	basic.Base
	basic.DataOwner
	TravelPolicyRatingTypeAPI
	TravelPolicy *TravelPolicy `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"` // travel policy
	RatingType   *RatingType   `json:"rating_type,omitempty" gorm:"foreignKey:RatingTypeID;reference:ID"`      // rating type
}

// TravelPolicyRatingTypeAPI Travel Policy Job Title API
type TravelPolicyRatingTypeAPI struct {
	TravelPolicyID *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // travel policy id
	RatingTypeID   *uuid.UUID `json:"rating_type_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`   // rating type id
	Rating         *int       `json:"rating,omitempty" gorm:"type:smallint;not null" example:"1"`
}
