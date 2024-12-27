package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgeQualifyingGroup struct
type AgeQualifyingGroup struct {
	basic.Base
	basic.DataOwner
	AgeQualifyingGroupAPI
	AgeQualifyingType *AgeQualifyingType `json:"age_qualifying_type,omitempty"`
}

// AgeQualifyingGroupAPI struct
type AgeQualifyingGroupAPI struct {
	AgeQualifyingTypeID *uuid.UUID `json:"age_qualifying_type_id,omitempty" gorm:"type:varchar(36);not null"`
	MinAge              *int       `json:"min_age,omitempty" gorm:"type:smallint"`
	MaxAge              *int       `json:"max_age,omitempty" gorm:"type:smallint"`
}
