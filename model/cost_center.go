package model

import (
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CostCenter Cost Center
type CostCenter struct {
	basic.Base
	basic.DataOwner
	CostCenterAPI
	CostCenterTranslation *CostCenterTranslation `json:"cost_center_translation,omitempty"` // Cost Center Translation
}

// CostCenterAPI Cost Center API
type CostCenterAPI struct {
	CostCenterCode *string `json:"cost_center_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null and cost_center_code is not null"` // Cost Center Code
	CostCenterName *string `json:"cost_center_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"`                        // Cost Center Name
}
