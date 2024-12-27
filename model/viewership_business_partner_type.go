package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViewershipBusinessPartnerType Viewership Business Partner Type
type ViewershipBusinessPartnerType struct {
	basic.Base
	basic.DataOwner
	ViewershipBusinessPartnerTypeAPI
	Viewership *Viewership `json:"viewership" gorm:"foreignKey:ViewershipID;references:ID"`
}

// ViewershipBusinessPartnerTypeAPI Viewership Business Partner Type Api
type ViewershipBusinessPartnerTypeAPI struct {
	ViewershipID          *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	BusinessPartnerTypeID *uuid.UUID `json:"business_partner_type_id" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
}
