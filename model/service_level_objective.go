package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceLevelObjective model
type ServiceLevelObjective struct {
	basic.Base
	basic.DataOwner
	ServiceLevelObjectiveAPI
	ServiceLevel *ServiceLevel `json:"service_level,omitempty"` // Service Level
	TaskType     *TaskType     `json:"task_type,omitempty"`     // Task Type
	Metrics      *Metrics      `json:"metrics,omitempty"`       // Metrics
}

// ServiceLevelObjectiveAPI API
type ServiceLevelObjectiveAPI struct {
	ServiceLevelID *uuid.UUID `json:"service_level_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Service Level ID
	MetricsID      *uuid.UUID `json:"metrics_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`       // Metrics ID
	TaskTypeID     *uuid.UUID `json:"task_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`              // Task Type ID
	Amount         *float64   `json:"amount,omitempty"`                                                                               // Amount
}

type ServiceLevelObjectiveResponse struct {
	ServiceLevelObjective
	Day    *string `json:"day,omitempty"`
	Hour   *string `json:"hour,omitempty"`
	Minute *string `json:"minute,omitempty"`
}

// ConfigurationStandardServicePostAPI API
type ConfigurationStandardServicePostAPI struct {
	ServiceLevelCode *int       `json:"service_level_code,omitempty"`
	TaskTypeID       *uuid.UUID `json:"task_type_id,omitempty"` // Task Type ID
	Amount           *float64   `json:"amount,omitempty"`       // Amount
}

// ConfigurationStandardServicePutAPI API
type ConfigurationStandardServicePutAPI struct {
	TaskTypeID *uuid.UUID `json:"task_type_id,omitempty"` // Task Type ID
	Amount     *float64   `json:"amount,omitempty"`       // Amount
}
