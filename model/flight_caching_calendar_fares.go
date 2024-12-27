package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// FlightCachingCalendarFares
type FlightCachingCalendarFares struct {
	basic.Base
	basic.DataOwner
	SessionID     *uuid.UUID    `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_calendar_fares_session_id;not null"`
	Price         *float64      `json:"price,omitempty"`
	IsCurrent     *bool         `json:"is_current,omitempty"`
	IsVisible     *bool         `json:"is_visible,omitempty"`
	IsLowestPrice *bool         `json:"is_lowest_price,omitempty"`
	CalendarDate  *CalendarDate `json:"calendar_date,omitempty" gorm:"type:text"` // multiple value separates by comma
}

type CalendarDate []string

// Implement Scanner interface
func (t *CalendarDate) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		oti, ok := value.([]uint8)
		if !ok {
			return errors.New(fmt.Sprint("wrong type : ", value))
		}
		val = string(oti)
	}
	*t = CalendarDate(strings.Split(string(val), ","))
	return nil
}

// Implement Valuer interface
func (t CalendarDate) Value() (driver.Value, error) {
	//this check is here if you don't want to save an empty string
	if len(t) == 0 {
		return nil, nil
	}
	return []byte(strings.Join(t, ",")), nil
}

func (CalendarDate) GormDataType() string {
	return "text"
}

func (CalendarDate) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "postgres", "sqlite":
		return "text"
	}
	return ""
}
