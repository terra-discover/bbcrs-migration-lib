package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViewershipLoyaltyLevel Viewership Loyalty Level
type ViewershipLoyaltyLevel struct {
	basic.Base
	basic.DataOwner
	ViewershipLoyaltyLevelAPI
	Viewership *Viewership `json:"viewership" gorm:"foreignKey:ViewershipID;references:ID"`
}

// ViewershipLoyaltyLevelAPI Viewership Loyalty Level Api
type ViewershipLoyaltyLevelAPI struct {
	ViewershipID   *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	LoyaltyLevelID *uuid.UUID `json:"loyalty_level_id" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
}
