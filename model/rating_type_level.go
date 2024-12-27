package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RatingTypeLevel Rating Type Level
type RatingTypeLevel struct {
	basic.Base
	basic.DataOwner
	RatingTypeLevelAPI
	RatingTypeLevelTranslation *RatingTypeLevelTranslation `json:"rating_type_level_translation,omitempty"`                                                                                                                                                                                                                            // Rating Type Level Translation
	RatingTypeID               *uuid.UUID                  `json:"rating_type_id,omitempty" gorm:"type:varchar(36);index:idx_rating_type_level_rating_type_level_code,unique,where:deleted_at is null;index:idx_rating_type_level_rating_type_level_name,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Rating type id
}

// RatingTypeLevelAPI Rating Type Level API
type RatingTypeLevelAPI struct {
	RatingTypeLevelCode *int64   `json:"rating_type_level_code,omitempty" gorm:"type:smallint;index:idx_rating_type_level_rating_type_level_code,unique,where:deleted_at is null;not null" example:"1"`       // Rating type level code
	RatingTypeLevelName *string  `json:"rating_type_level_name,omitempty" gorm:"type:varchar(256);index:idx_rating_type_level_rating_type_level_name,unique,where:deleted_at is null;not null" example:"Low"` // Rating type level name
	Rating              *float64 `json:"rating,omitempty" example:"1"`                                                                                                                                        // Rating value
}
