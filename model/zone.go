package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// Zone Zone
type Zone struct {
	basic.Base
	basic.DataOwner
	ZoneAPI
	ZoneTranslation *ZoneTranslation `json:"zone_translation,omitempty"`
	Destination     *Destination     `json:"destination,omitempty"`
}

// ZoneAPI Zone API
type ZoneAPI struct {
	ZoneCode      *string    `json:"zone_code,omitempty" gorm:"type:varchar(8);not null;index:idx_zone_code_deleted_at,unique,where:deleted_at is null" example:"JKT"`       // Zone Code
	ZoneName      *string    `json:"zone_name,omitempty" gorm:"type:varchar(256);not null;index:idx_zone_name_deleted_at,unique,where:deleted_at is null" example:"Jakarta"` // Zone Name
	DestinationID *uuid.UUID `json:"destination_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                                                         // Destination ID
	Description   *string    `json:"description,omitempty" gorm:"type:text" example:"Jakarta adalah Ibukota Negara Indonesia"`                                               // Description
}

// Seed Zone
func (zone *Zone) Seed() *Zone {
	seed := Zone{
		ZoneAPI: ZoneAPI{
			ZoneCode:      lib.Strptr("JKT"),
			ZoneName:      lib.Strptr("Jakarta"),
			DestinationID: lib.UUIDPtr(uuid.New()),
			Description:   lib.Strptr("Jakarta adalah Ibukota Negara Indonesia"),
		},
	}
	_ = lib.Merge(seed, &zone)
	return zone
}
