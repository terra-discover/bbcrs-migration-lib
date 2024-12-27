package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OperationTime struct
type OperationTime struct {
	basic.Base
	basic.DataOwner
	OperationTimeAPI
	OperationSchedule *OperationSchedule `json:"operation_schedule,omitempty" swaggerignore:"true"` // Operation Schedule
}

// OperationTimeAPI struct
type OperationTimeAPI struct {
	OperationScheduleID *uuid.UUID `json:"operation_schedule_id,omitempty" gorm:"type:varchar(36)"` // Operation Schedule Id
	Monday              *bool      `json:"monday,omitempty" gorm:"default:false"`                   // Active on Monday
	Tuesday             *bool      `json:"tuesday,omitempty" gorm:"default:false"`                  // Active on Tuesday
	Wednesday           *bool      `json:"wednesday,omitempty" gorm:"default:false"`                // Active on Wednesday
	Thursday            *bool      `json:"thursday,omitempty" gorm:"default:false"`                 // Active on Thursday
	Friday              *bool      `json:"friday,omitempty" gorm:"default:false"`                   // Active on Friday
	Saturday            *bool      `json:"saturday,omitempty" gorm:"default:false"`                 // Active on Saturday
	Sunday              *bool      `json:"sunday,omitempty" gorm:"default:false"`                   // Active on Sunday
	StartTime           *string    `json:"start_time,omitempty" gorm:"type:varchar(36)"`            // Start Time
	EndTime             *string    `json:"end_time,omitempty" gorm:"type:varchar(36)"`              // End Time
	Description         *string    `json:"description,omitempty" gorm:"type:text"`                  // Description
}
