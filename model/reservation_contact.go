package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationContact struct {
	basic.Base
	basic.DataOwner
	ReservationContactAPI
}

type ReservationContactAPI struct {
	ReservationID      *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ContactTypeID      *uuid.UUID `json:"contact_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	PersonID           *uuid.UUID `json:"person_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	EmployeeID         *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	CountryAccessCode  *string    `json:"country_access_code,omitempty" gorm:"type:varchar(3)"`
	ContactPhoneNumber *string    `json:"contact_phone_number,omitempty" gorm:"type:varchar(32)"`
	ContactEmail       *string    `json:"contact_email,omitempty" gorm:"type:varchar(256)"`
	IsPrimary          *bool      `json:"is_primary,omitempty"`
}
