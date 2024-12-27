package model

import (
	"github.com/go-openapi/strfmt"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OperationSchedule Operation Schedule
type OperationSchedule struct {
	basic.Base
	basic.DataOwner
	OperationScheduleAPI
	OperationScheduleTranslation *OperationScheduleTranslation `json:"operation_schedule_translation,omitempty"`
	OperationTime                []*OperationTime              `json:"operation_time,omitempty"`
}

// OperationScheduleAPI Operation Schedule API
type OperationScheduleAPI struct {
	OperationScheduleName *string          `json:"operation_schedule_name,omitempty" gorm:"type:varchar(256);uniqueIndex:idx_operation_schedule_name_parent_operation_schedule_id,where:deleted_at is null" example:"cleaning"` // Operation Schedule Name
	Start                 *strfmt.DateTime `json:"start,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`                                                                                             // Start
	End                   *strfmt.DateTime `json:"end,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`                                                                                               // End
	Description           *string          `json:"description,omitempty" gorm:"type:text" example:"deskripsi"`                                                                                                                  // Description
}
