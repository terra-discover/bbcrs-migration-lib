package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Market Market
type Market struct {
	basic.Base
	basic.DataOwner
	MarketAPI
	MarketTranslation *MarketTranslation `json:"market_translation,omitempty"` // Market Translation
}

// MarketAPI Market API
type MarketAPI struct {
	MarketCode  *int    `json:"market_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`             // Market Code
	MarketName  *string `json:"market_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Corporate"` // Market Name
	Description *string `json:"description,omitempty" gorm:"type:text"`                                                                             // Description
}

// Seed data
func (Market) Seed() *[]Market {
	data := []Market{}
	items := []string{
		"Corporate",
		"Retail",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, Market{
			MarketAPI: MarketAPI{
				MarketCode: &code,
				MarketName: &name,
			},
		})
	}

	return &data
}
