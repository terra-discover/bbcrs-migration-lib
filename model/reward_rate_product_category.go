package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardRateProductCategory Reward Rate Product Category
type RewardRateProductCategory struct {
	basic.Base
	basic.DataOwner
	RewardRateProductCategoryAPI
}

// RewardRateProductCategoryAPI Reward Rate Product Category API
type RewardRateProductCategoryAPI struct {
	RewardRateID      *uuid.UUID `json:"reward_rate_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`       //Reward Rate Id
	ProductCategoryID *uuid.UUID `json:"product_category_id,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Product Category Id
}
