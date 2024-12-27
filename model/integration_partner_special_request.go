package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerSpecialRequest struct
type IntegrationPartnerSpecialRequest struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerSpecialRequestAPI
	//IntegrationPartner *IntegrationPartner
	SpecialRequest *SpecialRequest
}

// IntegrationPartnerSpecialRequestAPI struct
type IntegrationPartnerSpecialRequestAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null"`                                       // IntegrationPartnerID
	SpecialRequestID     *uuid.UUID `json:"special_request_id,omitempty" gorm:"type:varchar(36);not null"`                                           // SpecialRequestID
	SpecialRequestCode   *string    `json:"special_request_code,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // SpecialRequestCode
	SpecialRequestName   *string    `json:"special_request_name,omitempty" gorm:"type:varchar(512);index:,unique,where:deleted_at is null"`          // SpecialRequestName
}
