package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Schedule Model
type Schedule struct {
	basic.Base
	basic.DataOwner
	ScheduleAPI
	OffsetDropTime *OffsetDropTime `json:"offset_drop_time,omitempty"`
	OffsetTimeUnit *OffsetTimeUnit `json:"offset_time_unit,omitempty"`
	RecurrenceType *RecurrenceType `json:"recurrence_type,omitempty"`
}

// ScheduleAPI API
type ScheduleAPI struct {
	ScheduleName         *string          `json:"schedule_name,omitempty" gorm:"type:varchar(256)"`
	Start                *strfmt.DateTime `json:"start,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	End                  *strfmt.DateTime `json:"end,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	AbsoluteDeadline     *strfmt.DateTime `json:"absolute_deadline,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	OffsetTimeUnitID     *uuid.UUID       `json:"offset_time_unit_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	OffsetUnitMultiplier *int             `json:"offset_unit_multiplier,omitempty" gorm:"type:smallint;" example:"1"`
	OffsetDropTimeID     *uuid.UUID       `json:"offset_drop_time_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	RecurrenceTypeID     *uuid.UUID       `json:"recurrence_type_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	Seconds              *string          `json:"seconds,omitempty" gorm:"type:varchar(64)"`
	Minutes              *string          `json:"minutes,omitempty" gorm:"type:varchar(64)"`
	Hours                *string          `json:"hours,omitempty" gorm:"type:varchar(64)"`
	DayOfTheMonth        *string          `json:"day_of_the_month,omitempty" gorm:"type:varchar(64)"`
	Months               *string          `json:"months,omitempty" gorm:"type:varchar(64)"`
	DayOfTheWeek         *string          `json:"day_of_the_week,omitempty" gorm:"type:varchar(64)"`
	Years                *string          `json:"years,omitempty" gorm:"type:varchar(64)"`
	Description          *string          `json:"description,omitempty" gorm:"type:text"`
	ScheduleAPIAddition
}

// ScheduleRequest Schedule Request
type ScheduleRequest struct {
	FirstReminder      *ScheduleAPI `json:"first_reminder,omitempty"`
	RecurrenceReminder *ScheduleAPI `json:"recurrence_reminder,omitempty"`
}

// ScheduleAPIAddition Schedule Addition
type ScheduleAPIAddition struct {
	From *int `json:"from,omitempty" gorm:"-"` // From
	To   *int `json:"to,omitempty" gorm:"-"`   // To
}
