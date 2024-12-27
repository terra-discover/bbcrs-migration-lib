package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardRateProductType Reward Rate Product Type
type RewardRateProductType struct {
	basic.Base
	basic.DataOwner
	RewardRateProductTypeAPI
}

// RewardRateProductTypeAPI Reward Rate Product Type API
type RewardRateProductTypeAPI struct {
	RewardRateID  *uuid.UUID `json:"reward_rate_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`  //Reward Rate Id
	ProductTypeID *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"` // Product Type Id
}
