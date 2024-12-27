package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// BusinessServiceType Business Service Type
type BusinessServiceType struct {
	basic.Base
	basic.DataOwner
	BusinessServiceTypeAPI
	BusinessServiceTypeTranslation *BusinessServiceTypeTranslation `json:"business_service_type_translation,omitempty"`
}

// BusinessServiceTypeAPI Business Service Type API
type BusinessServiceTypeAPI struct {
	BusinessServiceTypeCode *int    `json:"business_service_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null" example:"1"` // Business Service Type Code
	BusinessServiceName     *string `json:"business_service_name,omitempty" gorm:"type:varchar(256)" example:"Maintenance"`                                        // Business Service Name
}
