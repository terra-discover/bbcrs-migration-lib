package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OpsigoFlightClassObject model
type OpsigoFlightClassCaching struct {
	basic.Base
	basic.DataOwner
	SessionID     *uuid.UUID `json:"session_id" gorm:"type:varchar(36)"`
	Number        *string    `json:"number,omitempty" gorm:"type:varchar(36)"`
	ClassID       *string    `json:"class_id,omitempty" gorm:"type:varchar(4000)"`
	FlightID      *string    `json:"flight_id,omitempty" gorm:"type:varchar(4000)"`
	DepartDate    *string    `json:"depart_date,omitempty" gorm:"type:varchar(15)"`
	DepartTime    *string    `json:"depart_time,omitempty" gorm:"type:varchar(15)"`
	ArriveDate    *string    `json:"arrive_date,omitempty" gorm:"type:varchar(15)"`
	ArriveTime    *string    `json:"arrive_time,omitempty" gorm:"type:varchar(15)"`
	Code          *string    `json:"code,omitempty" gorm:"type:varchar(36)"`
	Category      *string    `json:"category,omitempty" gorm:"type:varchar(36)"`
	Seat          *int       `json:"seat,omitempty" gorm:"type:smallint"`
	Fare          *float64   `json:"fare,omitempty" gorm:"type:decimal(22,2)"`
	Tax           *float64   `json:"tax,omitempty" gorm:"type:decimal(22,2)"`
	FareBasisCode *string    `json:"fare_basis_code,omitempty" gorm:"type:varchar(10000)"`
	Sequence      *int       `json:"sequence,omitempty" gorm:"type:smallint"`
	ClassGroupId  *string    `json:"class_group_id,omitempty" gorm:"type:varchar(255)"`
	TransitTime   *string    `json:"transit_time,omitempty" gorm:"type:varchar(30)"`
	ExtraData     *string    `json:"extra_data,omitempty" gorm:"type:varchar(10000)"`
}
