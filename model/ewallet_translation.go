package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EWalletTranslation EWallet Translation
type EWalletTranslation struct {
	basic.Base
	basic.DataOwner
	EWalletTranslationAPI
	EWalletID *uuid.UUID `json:"e_wallet_id,omitempty" gorm:"type:varchar(36);uniqueIndex:e_wallet_translation_unique;not null" swaggertype:"string" format:"uuid"` // EWallet id
}

// EWalletTranslationAPI EWallet Translation API
type EWalletTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:e_wallet_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	EWalletName  *string `json:"e_wallet_name,omitempty" gorm:"type:varchar(256)" example:"Gopay"`                                             // E-Wallet Name
}
