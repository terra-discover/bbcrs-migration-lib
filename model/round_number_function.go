package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// RoundNumberFunction Round Number Function
type RoundNumberFunction struct {
	basic.Base
	basic.DataOwner
	RoundNumberFunctionAPI
	RoundNumberFunctionTranslation *RoundNumberFunctionTranslation `json:"round_number_function_translation,omitempty"` // Round Number Function Translation
}

// RoundNumberFunctionAPI Round Number Function API
type RoundNumberFunctionAPI struct {
	RoundFunctionCode *string `json:"round_function_code,omitempty" gorm:"type:varchar(1);index:,unique,where:deleted_at is null;not null"`   // Round Function Code
	RoundFunctionName *string `json:"round_function_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Round Function Name
}
