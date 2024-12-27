package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessEntityLoyaltyAccount struct
type BusinessEntityLoyaltyAccount struct {
	basic.Base
	basic.DataOwner
	BusinessEntityLoyaltyAccountAPI
	BusinessEntity *BusinessEntity `json:"business_entity,omitempty"`
}

// BusinessEntityLoyaltyAccountAPI for secure request body API
type BusinessEntityLoyaltyAccountAPI struct {
	BusinessEntityID *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36);not null;index:idx_business_entity_loyalty_account__business_entity_id,unique,where:deleted_at is null;"`
	LoyaltyAccountID *uuid.UUID `json:"loyalty_account_id,omitempty" gorm:"type:varchar(36);index:idx_business_entity_loyalty_account__business_entity_id,unique,where:deleted_at is null;"`
	IsPrimary        *bool      `json:"is_primary,omitempty" gorm:"type:bool;index:idx_business_entity_loyalty_account__business_entity_id,unique,where:deleted_at is null;"`
}
