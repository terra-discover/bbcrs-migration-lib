package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentSupplierHotel Agent Supplier Hotel
type AgentSupplierHotel struct {
	basic.Base
	basic.DataOwner
	AgentSupplierHotelAPI
	Agent         *Agent         `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	HotelSupplier *HotelSupplier `json:"hotel_supplier" gorm:"foreignKey:HotelSupplierID;references:ID"`
}

// AgentSupplierHotelAPI Agent Supplier Hotel Api
type AgentSupplierHotelAPI struct {
	AgentID         *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	HotelSupplierID *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
}
