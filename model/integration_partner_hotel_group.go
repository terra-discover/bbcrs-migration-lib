package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelGroup Integration Partner Hotel Group
type IntegrationPartnerHotelGroup struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelGroupAPI
	IntegrationPartnerHotelGroupTranslation *IntegrationPartnerHotelGroupTranslation `json:"integration_partner_hotel_group_translation,omitempty"` // Integration Partner Hotel Group Translation
}

// IntegrationPartnerHotelGroupAPI Integration Partner Hotel Group API
type IntegrationPartnerHotelGroupAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_hotel_group__hotel_group_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	HotelGroupID         *uuid.UUID `json:"hotel_group_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_hotel_group__hotel_group_code,unique,where:deleted_at is null;not null"`         // Hotel Group ID
	HotelGroupCode       *string    `json:"hotel_group_code,omitempty" gorm:"type:varchar(16);index:idx_integration_partner_hotel_group__hotel_group_code,unique,where:deleted_at is null;not null"`                                          // Hotel Group Code
	HotelGroupName       *string    `json:"hotel_group_name,omitempty" gorm:"type:varchar(512)"`                                                                                                                                              // Hotel Group Name
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                           // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                               // Comment
}
