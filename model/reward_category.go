package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// RewardCategory Reward Category
type RewardCategory struct {
	basic.Base
	basic.DataOwner
	RewardCategoryAPI
}

// RewardCategoryAPI Reward Category API
type RewardCategoryAPI struct {
	RewardCategoryCode *string `json:"reward_category_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`  //Reward Category Code
	RewardCategoryName *string `json:"reward_category_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Reward Category Name
	Description        *string `json:"description,omitempty" gorm:"type:varchar(4000);not null;index:,unique,where:deleted_at is null"`         // description
}
