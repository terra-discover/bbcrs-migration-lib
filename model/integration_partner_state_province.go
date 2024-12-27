package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerStateProvince Integration Partner State Province
type IntegrationPartnerStateProvince struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerStateProvinceAPI
	IntegrationPartnerStateProvinceTranslation *IntegrationPartnerStateProvinceTranslation `json:"integration_partner_state_province_translation,omitempty"` // Integration Partner State Province Translation
	StateProvince                              *StateProvince                              `json:"state_province,omitempty"`                                 // StateProvince
}

// IntegrationPartnerStateProvinceAPI Integration Partner State Province API
type IntegrationPartnerStateProvinceAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_state_province__state_province_id,unique,where:deleted_at is null;not null"` // Integration Partner ID
	StateProvinceID      *uuid.UUID `json:"state_province_id,omitempty" swaggertype:"string" format:"uuid" gorm:"index:idx_integration_partner_state_province__state_province_id,unique,where:deleted_at is null;not null"`                       // State Province ID
	StateProvinceCode    *string    `json:"state_province_code,omitempty" gorm:"type:varchar(8)"`                                                                                                                                                 // State Province Code
	StateProvinceName    *string    `json:"state_province_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                               // State Province Name
	CountryCode          *string    `json:"country_code,omitempty" gorm:"type:varchar(2)"`                                                                                                                                                        // Country Code
	Latitude             *string    `json:"latitude,omitempty" gorm:"type:varchar(16)"`                                                                                                                                                           // Latitude
	Longitude            *string    `json:"longitude,omitempty" gorm:"type:varchar(16)"`                                                                                                                                                          // Longitude
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                               // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                                   // Comment
}
