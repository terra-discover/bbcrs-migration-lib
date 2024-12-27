package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViewershipDestinationSystem Viewership Destination System
type ViewershipDestinationSystem struct {
	basic.Base
	basic.DataOwner
	ViewershipBusinessPartnerTypeAPI
	Viewership *Viewership `json:"viewership" gorm:"foreignKey:ViewershipID;references:ID"`
}

// ViewershipDestinationSystemAPI Viewership Destination System Api
type ViewershipDestinationSystemAPI struct {
	ViewershipID       *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	DestinatonSystemID *uuid.UUID `json:"destination_system_id" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
}
