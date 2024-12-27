package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardRateSupplierType Reward Rate Supplier Type
type RewardRateSupplierType struct {
	basic.Base
	basic.DataOwner
	RewardRateSupplierTypeAPI
}

// RewardRateSupplierTypeAPI Reward Rate Supplier Type API
type RewardRateSupplierTypeAPI struct {
	RewardRateID   *uuid.UUID `json:"reward_rate_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`   //Reward Rate Id
	SupplierTypeID *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"` // Supplier Type Id
	IsIncluded     *bool      `json:"is_included,omitempty"`                                                                              // Is Included
}
