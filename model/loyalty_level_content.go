package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyLevelContent Loyalty Level Content
type LoyaltyLevelContent struct {
	basic.Base
	basic.DataOwner
	LoyaltyLevelContentAPI
	LoyaltyLevelID     *uuid.UUID          `json:"loyalty_level_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Level ID
	ContentDescription *ContentDescription `json:"content_description,omitempty"`                                                                  // Content Description
}

// LoyaltyLevelContentAPI Loyalty Level Content API
type LoyaltyLevelContentAPI struct {
	ContentDescriptionID *uuid.UUID `json:"content_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Content Description ID
}
