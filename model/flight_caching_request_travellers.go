package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightCachingRequestTravellers model | many to one relation
type FlightCachingRequestTravellers struct {
	basic.Base
	basic.DataOwner
	SessionID        *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_request_travellers_session_id;not null"`
	FlightRequestID  *uuid.UUID `json:"flight_request_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // FlightRequestID
	EmployeeID       *string    `json:"employee_id,omitempty" gorm:"type:varchar(36)"`                     // EmployeeID
	FirstName        *string    `json:"first_name,omitempty" gorm:"type:varchar(255)"`
	IsGuest          *bool      `json:"is_guest,omitempty"`
	PassengerType    *string    `json:"passenger_type,omitempty"` // temporary attribute to classify of passenge types e.g:child,adult,infant
	LastName         *string    `json:"last_name,omitempty" gorm:"type:varchar(255)"`
	MiddleName       *string    `json:"middle_name,omitempty" gorm:"type:varchar(255)"`
	NamePrefix       *string    `json:"name_prefix,omitempty" gorm:"type:varchar(255)"`
	Age              *int       `json:"age,omitempty"`
	EmployeeName     *string    `json:"employee_name,omitempty" gorm:"-"`
	OrderIndex       *int       `json:"order_index,omitempty" gorm:"-"`
	EmployeeJobTitle *string    `json:"employee_job_title,omitempty" gorm:"-"`
	Employee         *Employee  `json:"-"`
}
