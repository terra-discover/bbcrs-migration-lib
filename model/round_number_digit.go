package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// RoundNumberDigit Round Number Digit
type RoundNumberDigit struct {
	basic.Base
	basic.DataOwner
	RoundNumberDigitAPI
	RoundNumberDigitTranslation *RoundNumberDigitTranslation `json:"round_number_digit_translation,omitempty"` // Round Number Digit Translation
}

// RoundNumberDigitAPI Round Number Digit API
type RoundNumberDigitAPI struct {
	RoundDigitCode *int    `json:"round_digit_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null"`     // Round Digit Code
	RoundDigitName *string `json:"round_digit_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Round Digit Name
}
