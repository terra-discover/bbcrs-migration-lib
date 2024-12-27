package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserTypeHotelSupplier Agent User Type Hotel Supplier
type AgentUserTypeHotelSupplier struct {
	basic.Base
	basic.DataOwner
	AgentUserTypeHotelSupplierAPI
	AgentUserType *AgentUserType `json:"agent_user_type,omitempty"`
	HotelSupplier *HotelSupplier `json:"hotel_supplier,omitempty"`
}

// AgentUserTypeHotelSupplierAPI Agent User Type Hotel Supplier Api
type AgentUserTypeHotelSupplierAPI struct {
	AgentUserTypeID *uuid.UUID `json:"agent_user_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	HotelSupplierID *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	IsAllowed       *bool      `json:"is_allowed,omitempty" gorm:"type:bool"`
}
