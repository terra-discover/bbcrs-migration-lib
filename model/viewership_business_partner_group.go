package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViewershipBusinessPartnerGroup model
type ViewershipBusinessPartnerGroup struct {
	basic.Base
	basic.DataOwner
	ViewershipBusinessPartnerGroupAPI
	Viewership *Viewership `json:"viewership" gorm:"foreignKey:ViewershipID;references:ID"`
}

// ViewershipBusinessPartnerGroupAPI model
type ViewershipBusinessPartnerGroupAPI struct {
	ViewershipID           *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	BusinessPartnerGroupID *uuid.UUID `json:"business_partner_group_id" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
}
