package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessEntityAccreditation struct
type BusinessEntityAccreditation struct {
	basic.Base
	basic.DataOwner
	BusinessEntityAccreditationAPI
	BusinessEntity *BusinessEntity `json:"business_entity,omitempty"`
}

// BusinessEntityAccreditationAPI for secure request body API
type BusinessEntityAccreditationAPI struct {
	BusinessEntityID    *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36);not null"`
	AccreditationTypeID *uuid.UUID `json:"accreditation_type_id,omitempty" gorm:"type:varchar(36);not null"`
	AccreditationNumber *string    `json:"accreditation_number,omitempty" gorm:"type:varchar(36);not null"`
	IsPrimary           *bool      `json:"is_primary,omitempty"`
}
