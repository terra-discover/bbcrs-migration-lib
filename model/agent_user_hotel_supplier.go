package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserHotelSupplier Agent User Hotel Supplier
type AgentUserHotelSupplier struct {
	basic.Base
	basic.DataOwner
	AgentUserHotelSupplierAPI
	AgentUser     *AgentUser     `json:"agent_user,omitempty"`
	HotelSupplier *HotelSupplier `json:"hotel_supplier,omitempty"`
}

// AgentUserHotelSupplierAPI Agent User Hotel Supplier Api
type AgentUserHotelSupplierAPI struct {
	AgentUserID     *uuid.UUID `json:"agent_user_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	HotelSupplierID *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	IsAllowed       *bool      `json:"is_allowed,omitempty" gorm:"type:bool"`
}
