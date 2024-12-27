package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateDivision Model
type CorporateDivision struct {
	basic.Base
	basic.DataOwner
	CorporateDivisionAPI
	Corporate *Corporate `json:"corporate,omitempty"`
	Division  *Division  `json:"division,omitempty" gorm:"foreignKey:DivisionID;references:ID"`
}

// CorporateDivisionAPI Model
type CorporateDivisionAPI struct {
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Corporate ID
	DivisionID  *uuid.UUID `json:"division_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`  // Division ID
}
