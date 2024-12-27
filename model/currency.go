package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// Currency struct Currency
type Currency struct {
	basic.Base
	basic.DataOwner
	CurrencyAPI
	CurrencyTranslation *CurrencyTranslation `json:"currency_translation,omitempty"`
}

// CurrencyAPI Currency API
type CurrencyAPI struct {
	CurrencyCode      *string `json:"currency_code,omitempty" gorm:"type:varchar(3);not null;index:idx_currency_code_deleted_at,unique,where:deleted_at is null" example:"IDR"`               // Currency Code
	CurrencyName      *string `json:"currency_name,omitempty" gorm:"type:varchar(64);not null;index:idx_currency_name_deleted_at,unique,where:deleted_at is null" example:"Rupiah Indonesia"` // Currency Name
	CurrencySymbol    *string `json:"currency_symbol,omitempty" gorm:"type:varchar(64)" example:"Rp"`                                                                                         // Currency Symbol
	NumericCode       *string `json:"numeric_code,omitempty" gorm:"type:varchar(3)" example:"360"`                                                                                            // Numeric Code
	MinorUnit         *int    `json:"minor_unit,omitempty" gorm:"type:smallint" example:"0"`                                                                                                  // Minor Unit
	MinorUnitName     *string `json:"minor_unit_name,omitempty" gorm:"type:varchar(64)" example:"Rupiah Indonesia"`                                                                           // Minor Unit Name
	StandardPrecision *int    `json:"standard_precision,omitempty" gorm:"type:smallint" example:"0"`                                                                                          // Standard Precision
	PricePrecision    *int    `json:"price_precision,omitempty" gorm:"type:smallint" example:"14000"`                                                                                         // Price Precision
	Position          *string `json:"position,omitempty" gorm:"type:varchar(8)" example:"Before"`                                                                                             // Position
}

// GetCurrency by query filter
func (data *Currency) GetCurrency(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).First(&data)
}

// Seed Currency
func (currency *Currency) Seed() *Currency {
	seed := Currency{
		CurrencyAPI: CurrencyAPI{
			CurrencyCode:      lib.Strptr("IDR"),
			CurrencyName:      lib.Strptr("Rupiah Indonesia"),
			CurrencySymbol:    lib.Strptr("Rp"),
			NumericCode:       lib.Strptr("360"),
			MinorUnit:         lib.Intptr(0),
			MinorUnitName:     lib.Strptr("Rupiah Indonesia"),
			StandardPrecision: lib.Intptr(0),
			PricePrecision:    lib.Intptr(14000),
			Position:          lib.Strptr("Before"),
		},
	}
	_ = lib.Merge(seed, &currency)
	return currency
}
