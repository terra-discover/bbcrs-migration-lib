package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateUserType Model
type CorporateUserType struct {
	basic.Base
	basic.DataOwner
	CorporateUserTypeAPI
	Corporate *Corporate `json:"corporate,omitempty"`
	UserType  *UserType  `json:"user_type,omitempty"`
}

// CorporateUserTypeAPI Model
type CorporateUserTypeAPI struct {
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Corporate ID
	UserTypeID  *uuid.UUID `json:"user_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // UserType ID
	IsDefault   *bool      `json:"is_default,omitempty"`
}
