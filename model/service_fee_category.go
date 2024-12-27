package model

import (
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeCategory Service Fee Category
type ServiceFeeCategory struct {
	basic.Base
	basic.DataOwner
	ServiceFeeCategoryAPI
	ServiceFeeCategoryTranslation *ServiceFeeCategoryTranslation `json:"service_fee_category_translation,omitempty"` // Service Fee Category Translation
	ServiceFeeRate                *[]ServiceFeeRate              `json:"service_fee_rate"`
	CorporateServiceFeeCategory   *CorporateServiceFeeCategory   `json:"corporate_service_fee_category,omitempty" gorm:"foreignKey:ServiceFeeCategoryID;"`
}

// ServiceFeeCategoryAPI Service Fee Category API
type ServiceFeeCategoryAPI struct {
	ServiceFeeCategoryCode *string `json:"service_fee_category_code,omitempty" gorm:"type:varchar(36);index:service_fee_category__service_fee_category_code,unique,where:deleted_at is null"`  // Service Fee Category Code
	ServiceFeeCategoryName *string `json:"service_fee_category_name,omitempty" gorm:"type:varchar(256);index:service_fee_category__service_fee_category_code,unique,where:deleted_at is null"` // Service Fee Category Name
	Description            *string `json:"description,omitempty" gorm:"type:text"`                                                                                                             // Description
}
