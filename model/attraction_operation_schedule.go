package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AttractionOperationSchedule struct
type AttractionOperationSchedule struct {
	basic.Base
	basic.DataOwner
	AttractionID        *uuid.UUID `json:"attraction_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`         // Attraction Id
	OperationScheduleID *uuid.UUID `json:"operation_schedule_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // Operation Schedule Id
}
