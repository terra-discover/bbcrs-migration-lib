package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type InquiryProfile struct {
	basic.Base
	basic.DataOwner
	InquiryProfileAPI
}

type InquiryProfileAPI struct {
	InquiryID          *uuid.UUID `json:"inquiry_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProfileTypeID      *uuid.UUID `json:"profile_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	PersonID           *uuid.UUID `json:"person_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	EmployeeID         *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	UserAccountID      *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	DivisionID         *uuid.UUID `json:"division_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	OfficeID           *uuid.UUID `json:"office_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ContactName        *string    `json:"contact_name,omitempty" gorm:"type:varchar(128)"`
	ContactPhoneNumber *string    `json:"contact_phone_number,omitempty" gorm:"type:varchar(32)"`
	ContactEmail       *string    `json:"contact_email,omitempty" gorm:"type:varchar(256)"`
}
