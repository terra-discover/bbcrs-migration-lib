package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramContent Loyalty Program Content
type LoyaltyProgramContent struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramContentAPI
}

// LoyaltyProgramContentAPI Loyalty Program Content API
type LoyaltyProgramContentAPI struct {
	LoyaltyProgramID     *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`     // Loyalty Program Id
	ContentDescriptionID *uuid.UUID `json:"content_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Content Description Id
}
