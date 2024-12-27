package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// Aircraft Aircraft
type Aircraft struct {
	basic.Base
	basic.DataOwner
	AircraftAPI
	AircraftTranslation *AircraftTranslation `json:"aircraft_translation,omitempty"`
}

// AircraftAPI Aircraft API
type AircraftAPI struct {
	AircraftCode *string `json:"aircraft_code,omitempty" gorm:"type:varchar(4);not null;index:,unique,where:deleted_at is null"`   // Aircraft Code
	AircraftName *string `json:"aircraft_name,omitempty" gorm:"type:varchar(128);not null;index:,unique,where:deleted_at is null"` // Aircraft Name
	Model        *string `json:"model,omitempty" gorm:"type:varchar(128)"`                                                         // Model
	IcaoCode     *string `json:"icao_code,omitempty" gorm:"type:varchar(4)"`                                                       // Icao Code
}

// GetAircraft by query filter
func (data *Aircraft) GetAircraft(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).First(&data)
}
