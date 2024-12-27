package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Bank Bank
type Bank struct {
	basic.Base
	basic.DataOwner
	BankAPI
	BankTranslation *BankTranslation `json:"bank_translation,omitempty"`
}

// BankAPI Bank API
type BankAPI struct {
	BankCode  *string `json:"bank_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" example:"bk"`    // Bank Code
	BankName  *string `json:"bank_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"bank"` // Bank Name
	SwiftCode *string `json:"swift_code,omitempty" gorm:"type:varchar(64)" example:"BNK"`                                                  // Swift Code
}
