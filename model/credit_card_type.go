package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// CreditCardType Credit Card Type
type CreditCardType struct {
	basic.Base
	basic.DataOwner
	CreditCardTypeAPI
	CreditCardTypeTranslation *CreditCardTypeTranslation `json:"credit_card_type_translation,omitempty"` // Credit Card Type Translation
	CreditCardTypeAsset       *CreditCardTypeAsset       `json:"credit_card_type_asset,omitempty"`       // Credit Card Type Asset
}

// CreditCardTypeAPI Credit Card Type API
type CreditCardTypeAPI struct {
	CreditCardTypeCode       *string                 `json:"credit_card_type_code,omitempty" gorm:"type:varchar(8);index:,unique,where:deleted_at is null;not null" example:"bnk"`    // Credit Card Type Code
	CreditCardTypeName       *string                 `json:"credit_card_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"bank"` // Credit Card Type Name
	CardIdentificationRemark *string                 `json:"card_identification_remark,omitempty" gorm:"type:text" example:"12345"`                                                   // Card Identification Remark
	CreditCardTypeAsset      *CreditCardTypeAssetAPI `json:"credit_card_type_asset,omitempty" gorm:"-"`                                                                               // Credit Card Type Asset
}
