package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Voucher Voucher
type Voucher struct {
	basic.Base
	basic.DataOwner
	VoucherAPI
	VoucherTranslation *VoucherTranslation `json:"voucher_translation,omitempty"` // Voucher Translation
}

// VoucherAPI Voucher API
type VoucherAPI struct {
	VoucherCode           *string          `json:"voucher_code,omitempty" example:"vouchercode" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`  // Voucher Code
	VoucherName           *string          `json:"voucher_name,omitempty" example:"voucher code" gorm:"type:varchar(64);index:,unique,where:deleted_at is null;not null"` // Voucher Name
	VoucherTypeID         *uuid.UUID       `json:"voucher_type_id,omitempty" swaggertype:"string" format:"uuid"`                                                          // Voucher Type ID
	CurrencyID            *uuid.UUID       `json:"currency_id,omitempty" swaggertype:"string" format:"uuid"`                                                              // Currency ID
	Amount                *float64         `json:"amount,omitempty" gorm:"type:decimal(19,4)" example:"5000.0"`                                                           // Amount
	RewardTypeID          *uuid.UUID       `json:"reward_type_id,omitempty" swaggertype:"string" format:"uuid"`                                                           // Reward Type ID
	RewardAmount          *float64         `json:"reward_amount,omitempty" gorm:"type:decimal(19,4)" example:"5000.0"`                                                    // Reward Amount
	Value                 *float64         `json:"value,omitempty" gorm:"type:decimal(19,4)" example:"5000.0"`                                                            // Value
	Percent               *float64         `json:"percent,omitempty" gorm:"type:decimal(19,4)" example:"5.0"`                                                             // Percent
	MaxAmount             *float64         `json:"max_amount,omitempty" gorm:"type:decimal(19,4)" example:"5000.0"`                                                       // Max Amount
	Use                   *int             `json:"use,omitempty" gorm:"type:smallint"`                                                                                    // Use
	UnlimitedUse          *bool            `json:"unlimited_use,omitempty" example:"false"`                                                                               // Unlimited Use
	EffectiveDate         *strfmt.DateTime `json:"effective_date,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`                              // Effective Date
	ExpireDate            *strfmt.DateTime `json:"expire_date,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`                                 // Expire Date
	VoucherUsed           *int             `json:"voucher_used,omitempty" gorm:"type:smallint"`                                                                           // Voucher Used
	AssignAllProductTypes *bool            `json:"assign_all_product_types,omitempty" example:"false"`                                                                    // Assign All Product Types
	AssignAllAirlanes     *bool            `json:"assign_all_airlanes,omitempty" example:"false"`                                                                         // Assign All Airlanes
	AssignAllHotelBrands  *bool            `json:"assign_all_hotel_brands,omitempty" example:"false"`                                                                     // Assign All Hotel Brands
	AssignAllHotels       *bool            `json:"assign_all_hotels,omitempty" example:"false"`                                                                           // Assign All Hotels
	VoucherTransactionID  *uuid.UUID       `json:"voucher_transaction_id,omitempty" swaggertype:"string" format:"uuid"`                                                   // Voucher Transaction ID
	Comment               *string          `json:"comment,omitempty" example:"comment" gorm:"type:text"`                                                                  // Comment
	Description           *string          `json:"description,omitempty" example:"descriptions" gorm:"type:text"`                                                         // Description
}
