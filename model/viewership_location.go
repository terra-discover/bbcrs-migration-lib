package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViewershipLocation Viewership Location
type ViewershipLocation struct {
	basic.Base
	basic.DataOwner
	ViewershipLocationAPI
	Viewership *Viewership `json:"viewership" gorm:"foreignKey:ViewershipID;references:ID"`
}

// ViewershipLocationAPI Viewership Location Api
type ViewershipLocationAPI struct {
	ViewershipID  *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	CountryID     *uuid.UUID `json:"country_id" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	CityID        *uuid.UUID `json:"city_id" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	DestinationID *uuid.UUID `json:"destination_id" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	ZoneID        *uuid.UUID `json:"zone_id" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
}
