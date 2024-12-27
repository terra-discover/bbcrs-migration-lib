package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// BankAccountType Bank Account Type
type BankAccountType struct {
	basic.Base
	basic.DataOwner
	BankAccountTypeAPI
	BankAccountTypeTranslation *BankAccountTypeTranslation `json:"bank_account_type_translation,omitempty"` // Bank Account Type Translation
}

// BankAccountTypeAPI Bank Account Type API
type BankAccountTypeAPI struct {
	BankAccountTypeCode *int    `json:"bank_account_type_code,omitempty" gorm:"type:int;index:,unique,where:deleted_at is null;not null"`          // Bank Account Type Code
	BankAccountTypeName *string `json:"bank_account_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Bank Account Type Name
}
