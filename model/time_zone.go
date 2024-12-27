package model

import (
	"time"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// TimeZone Timezone
type TimeZone struct {
	basic.Base
	basic.DataOwner
	TimeZoneAPI
	TimeZoneTranslation *TimeZoneTranslation `json:"time_zone_translation,omitempty"` // Timezone Translation
}

// TimeZoneAPI Timezone API
type TimeZoneAPI struct {
	UTCOffset                       *string  `json:"utc_offset,omitempty" gorm:"type:varchar(8)" example:"+07:00"`                      // UTC Offset
	UTCOffsetHour                   *float64 `json:"utc_offset_hour,omitempty" example:"7.0"`                                           // UTC Offset Hour
	UTCDaylightSavingTimeOffset     *string  `json:"utc_daylight_saving_time_offset,omitempty" gorm:"type:varchar(8)" example:"+07:00"` // UTC Daylight Saving Time Offset
	UTCDaylightSavingTimeOffsetHour *float64 `json:"utc_daylight_saving_time_offset_hour,omitempty" example:"7.0"`                      // UTC Daylight Saving Time Offset Hour
	ZoneName                        *string  `json:"zone_name,omitempty" gorm:"type:varchar(64)" example:"Bangkok, Jakarta"`            // Zone Name
	Description                     *string  `json:"description,omitempty" gorm:"type:text" example:"Bangkok, Jakarta"`                 // Description
}

// Seed TimeZone
func (b *TimeZone) Seed(agentID uuid.UUID) *[]TimeZone {
	now := strfmt.DateTime(time.Now().UTC())

	data := []TimeZone{}

	if lib.IsEmptyUUID(agentID) {
		agentID = uuid.New()
	}

	timeZoneID := agentID
	initial := TimeZone{}
	initial.ID = &timeZoneID
	initial.UTCOffset = lib.Strptr("+07:00")
	initial.UTCOffsetHour = lib.Float64ptr(+7)
	initial.ZoneName = lib.Strptr("(UTC+07:00) Bangkok, Hanoi, Jakarta")
	initial.CreatedAt = &now
	data = append(data, initial)

	return &data
}
