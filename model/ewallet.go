package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// EWallet EWallet
type EWallet struct {
	basic.Base
	basic.DataOwner
	EWalletAPI
	EWalletTranslation *EWalletTranslation `json:"e_wallet_translation,omitempty"`
}

// EWalletAPI EWallet API
type EWalletAPI struct {
	EWalletCode *string `json:"e_wallet_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" example:"A001"` // E-Wallet Code
	EWalletName *string `json:"e_wallet_name,omitempty" gorm:"type:varchar(256)" example:"Gopay"`                                               // E-Wallet Name
}
