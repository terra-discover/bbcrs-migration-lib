package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CallOut Call Out
type CallOut struct {
	basic.Base
	basic.DataOwner
	CallOutAPI
	CallOutTranslation *CallOutTranslation `json:"call_out_translation,omitempty"`
	CallOutAsset       *CallOutAsset       `json:"call_out_asset,omitempty"`
}

// CallOutAPI Call Out API
type CallOutAPI struct {
	CallOutName     *string       `json:"call_out_name,omitempty" gorm:"type:varchar(256);not null"` // Call Out Name
	Color           *string       `json:"color,omitempty" gorm:"type:varchar(36)"`                   // Color
	BackgroundColor *string       `json:"background_color,omitempty" gorm:"type:varchar(36)"`        // Background Color
	Description     *string       `json:"description,omitempty" gorm:"type:text"`                    // Description
	CallOutAsset    *CallOutAsset `json:"call_out_asset,omitempty" gorm:"-"`                         // CallOutAsset
}

// Seed CallOut
func (callOut *CallOut) Seed() *CallOut {
	seed := CallOut{
		CallOutAPI: CallOutAPI{
			CallOutName:     lib.Strptr("New"),
			Color:           lib.Strptr("Orange"),
			BackgroundColor: lib.Strptr("Blue"),
			Description:     nil,
		},
	}
	_ = lib.Merge(seed, &callOut)
	return callOut
}

// CallOuts CallOuts seeder
type CallOuts []CallOut

// Seed CallOuts Seed
func (c *CallOuts) Seed() *CallOuts {
	name := []string{"New", "Important", "Hot Deal", "Promo"}
	for i := 0; i < len(name); i++ {
		*c = append(*c, CallOut{CallOutAPI: CallOutAPI{
			CallOutName: lib.Strptr(name[i]),
		}})
	}
	return c
}
