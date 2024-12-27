package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Division Model
type Division struct {
	basic.Base
	basic.DataOwner
	DivisionAPI
	AgentDivision       *AgentDivision       `json:"agent_division,omitempty" gorm:"foreignKey:ID;references:DivisionID" swaggerignore:"true"` //Agent Division
	DivisionTranslation *DivisionTranslation `json:"division_translation,omitempty"  gorm:"foreignKey:DivisionID;references:ID"`               // Division Translation
	ParentDivision      *Division            `json:"parent_division,omitempty" swaggerignore:"true"`                                           // Parent Division
}

// DivisionAPI Division API
type DivisionAPI struct {
	DivisionCode     *string    `json:"division_code,omitempty" gorm:"type:varchar(36);not null"`                                // Division Code
	DivisionName     *string    `json:"division_name,omitempty" gorm:"type:varchar(256);not null"`                               // Division Name
	ParentDivisionID *uuid.UUID `json:"parent_division_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // Parent Division ID
	ManagerID        *uuid.UUID `json:"manager_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`         // Manager ID / Employee ID
	Depth            *int       `json:"depth,omitempty"`                                                                         // Depth
}
