package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateProductTypeRestriction Model
type CorporateProductTypeRestriction struct {
	basic.Base
	basic.DataOwner
	CorporateProductTypeRestrictionAPI
	Corporate   *Corporate   `json:"corporate,omitempty" gorm:"foreignKey:CorporateID;referenes:ID"`       // Corporate
	ProductType *ProductType `json:"product_type,omitempty" gorm:"foreignKey:ProductTypeID;references:ID"` // ProductType
}

// CorporateProductTypeRestrictionAPI Model
type CorporateProductTypeRestrictionAPI struct {
	CorporateID              *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`    // Corporate ID
	ProductTypeID            *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // ProductType ID
	MinAdvancedBookingOffset *int64     `json:"min_advanced_booking_offset,omitempty"`
}

type CorporateProductTypeRestrictionResponse struct {
	MinAdvancedBookingOffset *int64 `json:"min_advanced_booking_offset,omitempty"`
}
