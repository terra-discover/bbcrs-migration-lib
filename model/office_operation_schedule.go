package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OfficeOperationSchedule Model
type OfficeOperationSchedule struct {
	basic.Base
	basic.DataOwner
	OfficeOperationScheduleAPI
	Office            *Office            `json:"office" gorm:"foreignKey:OfficeID;references:ID"`
	OperationSchedule *OperationSchedule `json:"operation_schedule" gorm:"foreignKey:OperationScheduleID;references:ID"`
}

// OfficeOperationScheduleAPI API
type OfficeOperationScheduleAPI struct {
	OfficeID            *uuid.UUID `json:"office_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	OperationScheduleID *uuid.UUID `json:"operation_schedule_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
