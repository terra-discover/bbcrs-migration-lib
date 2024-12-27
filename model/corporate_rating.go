package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateRating Corporate Rating
type CorporateRating struct {
	basic.Base
	basic.DataOwner
	CorporateRatingAPI
	Corporate           *Corporate           `json:"corporate,omitempty"`             // Corporate
	CorporateRatingType *CorporateRatingType `json:"corporate_rating_type,omitempty"` // Corporate Rating Type
}

// CorporateRatingAPI Corporate Rating Api
type CorporateRatingAPI struct {
	CorporateRatingCode   *int       `json:"corporate_rating_code,omitempty" gorm:"-" example:"1"`                                                                                                                                 // Corporate Rating Code
	CorporateRatingName   *string    `json:"corporate_rating_name,omitempty" gorm:"-" example:"1"`                                                                                                                                 // Corporate Rating Name
	Rating                *float64   `json:"rating,omitempty" example:"1.5"`                                                                                                                                                       // Rating
	CorporateID           *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null;index:idx_corporate_rating__corporate_rating_type_id,where:deleted_at is null" swaggertype:"string" format:"uuid"`             // Corporate Id
	CorporateRatingTypeID *uuid.UUID `json:"corporate_rating_type_id,omitempty" gorm:"type:varchar(36);not null;index:idx_corporate_rating__corporate_rating_type_id,where:deleted_at is null" swaggertype:"string" format:"uuid"` // Corporate Rating Type Id
	IsSummary             *bool      `json:"is_summary,omitempty"`                                                                                                                                                                 // Is Summary
}
