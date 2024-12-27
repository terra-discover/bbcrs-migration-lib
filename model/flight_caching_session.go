package model

import (
	"database/sql/driver"
	"errors"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type FlightCachingSession struct {
	basic.Base
	basic.DataOwner
	SessionID      *uuid.UUID                  `json:"-" gorm:"type:varchar(36);index:idx_flight_caching_session_request_session_id"`
	SessionPayload FlightCachingSessionPayload `json:"session_payload"`
}

// AfterCreate Data
func (f *FlightCachingSession) AfterCreate(tx *gorm.DB) error {
	return afterCreateFlightCachingSession(tx, *f)
}

// AfterSave Data
func (f *FlightCachingSession) AfterSave(tx *gorm.DB) error {
	return afterSaveFlightCachingSession(tx, *f)
}

// custom data types
type FlightCachingSessionPayload map[string]interface{}

func (p FlightCachingSessionPayload) Value() (driver.Value, error) {
	return lib.JSONMarshal(p)
}

func (p *FlightCachingSessionPayload) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return lib.JSONUnmarshal(b, &p)
}

func (FlightCachingSessionPayload) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

func (FlightCachingSessionPayload) GormDataType() string {
	return "json"
}
