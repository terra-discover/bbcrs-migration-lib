package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// RewardType Reward Type
type RewardType struct {
	basic.Base
	basic.DataOwner
	RewardTypeAPI
	RewardTypeTranslation *RewardTypeTranslation `json:"reward_type_translation,omitempty"` // Reward Type Translation
}

// RewardTypeAPI Reward Type API
type RewardTypeAPI struct {
	RewardTypeCode *int    `json:"reward_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null"`     // Reward Type Code
	RewardTypeName *string `json:"reward_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Reward Type Name
}

// Seed data
func (s RewardType) Seed() *[]RewardType {
	data := []RewardType{}
	items := []string{
		"Points",
		"Miles",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, RewardType{
			RewardTypeAPI: RewardTypeAPI{
				RewardTypeCode: &code,
				RewardTypeName: &name,
			},
		})
	}

	return &data
}
