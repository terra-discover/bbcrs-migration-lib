package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// VoucherType Voucher Type
type VoucherType struct {
	basic.Base
	basic.DataOwner
	VoucherTypeAPI
	VoucherTypeTranslation *VoucherTypeTranslation `json:"voucher_type_translation,omitempty"` // Voucher Type Translation
}

// VoucherTypeAPI Voucher Type API
type VoucherTypeAPI struct {
	VoucherTypeCode       *string `json:"voucher_type_code,omitempty" example:"vouchercode" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`   // Voucher Type Code
	VoucherTypeName       *string `json:"voucher_type_name,omitempty" example:"voucher code" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Voucher Type Name
	AssignAllProductTypes *bool   `json:"assign_all_product_types,omitempty" example:"false"`                                                                          // Assign All Product Types
	AssignAllAirlanes     *bool   `json:"assign_all_airlanes,omitempty" example:"false"`                                                                               // Assign All Airlanes
	AssignAllHotelBrands  *bool   `json:"assign_all_hotel_brands,omitempty" example:"false"`                                                                           // Assign All Hotel Brands
	AssignAllHotels       *bool   `json:"assign_all_hotels,omitempty" example:"false"`                                                                                 // Assign All Hotels
	Comment               *string `json:"comment,omitempty" example:"comments" gorm:"type:text"`                                                                       // Comment
	Description           *string `json:"description,omitempty" example:"descriptions" gorm:"type:text"`                                                               // Description
}
