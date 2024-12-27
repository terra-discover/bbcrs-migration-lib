package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateEmailCategory Corporate Email Category
type CorporateEmailCategory struct {
	basic.Base
	basic.DataOwner
	CorporateEmailCategoryAPI
	Corporate     *Corporate     `json:"corporate,omitempty"`      // Corporate
	EmailCategory *EmailCategory `json:"email_category,omitempty"` // Email Category
}

// CorporateEmailCategoryAPI Corporate Email Category API
type CorporateEmailCategoryAPI struct {
	CorporateID     *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`      // corporate_id
	EmailCategoryID *uuid.UUID `json:"email_category_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // email_category_id
}
