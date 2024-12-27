package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// StateProvinceCategory State Province Category
type StateProvinceCategory struct {
	basic.Base
	basic.DataOwner
	StateProvinceCategoryAPI
	StateProvinceCategoryTranslation *StateProvinceCategoryTranslation `json:"state_province_category_translation,omitempty"`
}

// StateProvinceCategoryAPI State Province Category API
type StateProvinceCategoryAPI struct {
	StateProvinceCategoryCode *string `json:"state_province_category_code,omitempty" gorm:"type:varchar(8);uniqueIndex:idx_state_province_category_code_deleted_at,where:deleted_at is null;not null" example:"JOG"` // State Province Category Code
	StateProvinceCategoryName *string `json:"state_province_category_name,omitempty" gorm:"type:varchar(256)" example:"Jogjakarta"`                                                                                  // State Province Category Name
}
