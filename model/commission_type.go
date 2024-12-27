package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// CommissionType Commission Type
type CommissionType struct {
	basic.Base
	basic.DataOwner
	CommissionTypeAPI
	CommissionTypeTranslation *CommissionTypeTranslation `json:"commission_type_translation,omitempty"` // Commission Type Translation
}

// CommissionTypeAPI Commission Type API
type CommissionTypeAPI struct {
	CommissionTypeCode *int    `json:"commission_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null"`     // Commission Type Code
	CommissionTypeName *string `json:"commission_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Commission Type Name
	Description        *string `json:"description,omitempty" gorm:"type:text"`                                                                  // Description
}
