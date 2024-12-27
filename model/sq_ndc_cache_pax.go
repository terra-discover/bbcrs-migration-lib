package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SqNdcCachePax table struct
type SqNdcCachePax struct {
	basic.Base
	basic.DataOwner
	SessionID    *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // session ID
	PaxID        *string    `json:"pax_id,omitempty" gorm:"type:varchar(60)" default:"PAX1"`
	PTC          *string    `json:"ptc,omitempty" gorm:"type:varchar(256)" default:"ADT"`
	PaxRefID     *string    `json:"pax_ref_id,omitempty" gorm:"type:varchar(60)"`
	IndividualID *string    `json:"individual_id,omitempty" gorm:"type:varchar(60)" default:"PAX1"`
	Birthdate    *string    `json:"birth_date,omitempty" gorm:"type:varchar(60)" default:"1995-11-07"`
	GivenName    *string    `json:"given_name,omitempty" gorm:"type:varchar(60)" default:"DIKHI"`
	Surname      *string    `json:"surname,omitempty" gorm:"type:varchar(60)" default:"MARTIN"`
}

// GetSqNdcCachePax by query filter
func (data *SqNdcCachePax) GetSqNdcCachePax(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Take(&data)
}

// GetSqNdcCachePax by query filter
func (data *SqNdcCachePax) GetSqNdcCachePaxs(tx *gorm.DB, queryFilters string) []SqNdcCachePax {
	res := []SqNdcCachePax{}
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Find(&res)
	return res
}
