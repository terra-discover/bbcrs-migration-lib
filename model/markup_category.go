package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// MarkupCategory Markup Category
type MarkupCategory struct {
	basic.Base
	basic.DataOwner
	MarkupCategoryAPI
	MarkupCategoryTranslation *MarkupCategoryTranslation `json:"markup_category_translation,omitempty"`
	MarkupRate                []MarkupRate               `json:"markup_rate" gorm:"foreignKey:MarkupCategoryID"`
}

// MarkupCategoryAPI Markup Category API
type MarkupCategoryAPI struct {
	MarkupCategoryCode *string `json:"markup_category_code" gorm:"type:varchar(36);unique"` //MarkupCategoryCode
	MarkupCategoryName *string `json:"markup_category_name" gorm:"type:varchar(256)"`       //MarkupCategoryName
	Description        *string `json:"description" gorm:"type:varchar(4000)"`               //Description
}
