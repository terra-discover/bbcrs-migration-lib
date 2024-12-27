package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// ProcessingFeeCategory Processing Fee Category
type ProcessingFeeCategory struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeCategoryAPI
	ProcessingFeeCategoryTranslation *ProcessingFeeCategoryTranslation `json:"processing_fee_category_translation,omitempty"` // Processing Fee Category Translation
	ProcessingFeeRate                []ProcessingFeeRate               `json:"processing_fee_rate"`
}

// ProcessingFeeCategoryAPI Processing Fee Category API
type ProcessingFeeCategoryAPI struct {
	ProcessingFeeCategoryCode *string `json:"processing_fee_category_code,omitempty" format:"uuid" gorm:"type:varchar(36);index:,unique,where:deleted_at is null and processing_fee_category_code is not null"`     // Processing Fee Category Code
	ProcessingFeeCategoryName *string `json:"processing_fee_category_name,omitempty" example:"Flight Preset" gorm:"type:varchar(256);index:,where:deleted_at is null and processing_fee_category_name is not null"` // Processing Fee Category Name
	Description               *string `json:"description,omitempty" gorm:"type:text"`                                                                                                                               // Description
}
