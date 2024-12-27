package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardRateHotelSupplier Reward Rate Hotel Supplier
type RewardRateHotelSupplier struct {
	basic.Base
	basic.DataOwner
	RewardRateHotelSupplierAPI
}

// RewardRateHotelSupplierAPI Reward Rate Hotel Supplier API
type RewardRateHotelSupplierAPI struct {
	RewardRateID    *uuid.UUID `json:"reward_rate_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`     //Reward Rate Id
	HotelSupplierID *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Hotel Supplier Id
	IsIncluded      *bool      `json:"is_included,omitempty"`                                                                                // Is Included
}
