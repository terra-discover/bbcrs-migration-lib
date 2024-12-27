package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Benefit struct is not writable by request body
type Benefit struct {
	basic.Base
	basic.DataOwner
	BenefitAPI
	BenefitTranslation *BenefitTranslation `json:"benefit_translation,omitempty"` // Benefit Translation
	BenefitAsset       *BenefitAsset       `json:"benefit_asset,omitempty"`       // Benefit Asset
}

// BenefitAPI struct is writable by request body
type BenefitAPI struct {
	BenefitName  *string          `json:"benefit_name,omitempty" gorm:"type:varchar(256);not null;index:unique_benefit__benefit_name,unique,where:deleted_at is null" example:"free parking"` // benefit name
	Description  *string          `json:"description,omitempty" gorm:"type:text"`                                                                                                             // description
	BenefitAsset *BenefitAssetAPI `json:"benefit_asset,omitempty" gorm:"-"`                                                                                                                   // benefit assets
}
