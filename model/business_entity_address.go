package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessEntityAddress struct
type BusinessEntityAddress struct {
	basic.Base
	basic.DataOwner
	BusinessEntityAddressAPI
	Address *Address `json:"address,omitempty"` // address
}

// BusinessEntityAddressAPI for secure request body API
type BusinessEntityAddressAPI struct {
	BusinessEntityID *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36);not null;index:idx_business_entity_address__business_entity_id,unique,where:deleted_at is null;"` // business entity id
	AddressTypeID    *uuid.UUID `json:"address_type_id,omitempty" gorm:"type:varchar(36);index:idx_business_entity_address__business_entity_id,unique,where:deleted_at is null;"`             // address type id
	AddressID        *uuid.UUID `json:"address_id,omitempty" gorm:"type:varchar(36)"`                                                                                                         // address id
	IsPrimary        *bool      `json:"is_primary,omitempty" gorm:"type:bool;index:idx_business_entity_address__business_entity_id,unique,where:deleted_at is null;"`                         // is primary
}
