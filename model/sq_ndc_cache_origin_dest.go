package model

import (
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// SqNdcCacheOriginDest table struct
type SqNdcCacheOriginDest struct {
	basic.Base
	basic.DataOwner
	SessionID       *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	OriginDestID    *string    `json:"origin_dest_id,omitempty" gorm:"type:varchar(60)" default:"CGKSIN"`
	OriginCode      *string    `json:"origin_code,omitempty" gorm:"type:varchar(60)" default:"CGK"`
	DestCode        *string    `json:"dest_code,omitempty"  gorm:"type:varchar(60)" default:"SIN"`
	PaxJourneyRefID *string    `json:"pax_journey_ref_id,omitempty"` // multiple with comas
}

// // Seed SqNdcCacheOriginDest
// func (d *SqNdcCacheOriginDest) Seed() *SqNdcCacheOriginDest {
// 	sessionID, _ := uuid.Parse(viper.GetString("SESSION_ID"))

// 	seed := SqNdcCacheOriginDest{
// 		SessionID:       &sessionID,
// 		OriginDestID:    lib.Strptr("SINCGK"),
// 		OriginCode:      lib.Strptr("SIN"),
// 		DestCode:        lib.Strptr("CGK"),
// 		PaxJourneyRefID: lib.Strptr("FLT2"),
// 	}
// 	_ = lib.Merge(seed, &d)
// 	return d
// }
