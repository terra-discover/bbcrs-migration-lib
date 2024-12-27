package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViewershipBusinessPartner Viewership Business Partner
type ViewershipBusinessPartner struct {
	basic.Base
	basic.DataOwner
	ViewershipBusinessPartnerTypeAPI
	Viewership *Viewership `json:"viewership" gorm:"foreignKey:ViewershipID;references:ID"`
}

// ViewershipBusinessPartnerAPI Viewership Business Partner Api
type ViewershipBusinessPartnerAPI struct {
	ViewershipID      *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	BusinessPartnerID *uuid.UUID `json:"business_partner_id" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
}
