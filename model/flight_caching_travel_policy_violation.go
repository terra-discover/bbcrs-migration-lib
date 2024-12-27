package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingTravelPolicyViolation struct {
	basic.Base
	basic.DataOwner
	SessionID               *uuid.UUID  `json:"session_id" gorm:"type:varchar(36)"`
	FlightCachingID         *uuid.UUID  `json:"flight_caching_id" gorm:"type:varchar(36);index:idx_flight_caching_travel_policy_violation_flight_caching_id"`
	FlightCachingAirlinesID *uuid.UUID  `json:"flight_caching_airlines_id" gorm:"type:varchar(36);index:idx_flight_caching_travel_policy_violation_flight_caching_airlines_id"`
	FlightCachingCabinsID   *uuid.UUID  `json:"flight_caching_cabins_id" gorm:"type:varchar(36);index:idx_flight_caching_travel_policy_violation_flight_caching_cabins_id"`
	Key                     *string     `json:"key" gorm:"type:varchar(100)"`
	Label                   *string     `json:"label" gorm:"type:varchar(100)"`
	RestrictedRule          *string     `json:"restricted_rule" gorm:"type:varchar(100)"`
	CurrentValue            *string     `json:"current_value" gorm:"type:varchar(100)"`
	Additional              interface{} `json:"additional" gorm:"-"`
}
