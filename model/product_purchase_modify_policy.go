package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProductPurchaseModifyPolicy struct {
	basic.Base
	basic.DataOwner
	ProductPurchaseModifyPolicyAPI
}

type ProductPurchaseModifyPolicyAPI struct {
	ProductPurchaseID    *uuid.UUID `json:"product_purchase_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ModifyPolicyID       *uuid.UUID `json:"modify_policy_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ModifyPolicyCode     *string    `json:"modify_policy_code,omitempty" gorm:"type:varchar(36)"`
	ModifyPolicyName     *string    `json:"modify_policy_name,omitempty" gorm:"type:varchar(256)"`
	OffsetTimeUnitID     *uuid.UUID `json:"offset_time_unit_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	OffsetUnitMultiplier *int       `json:"offset_unit_multiplier,omitempty" gorm:"type:smallint"`
	OffsetDropTimeID     *uuid.UUID `json:"offset_drop_time_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	NotModifiable        *bool      `json:"not_modifiable,omitempty"`
	NotRefundable        *bool      `json:"not_refundable,omitempty"`
	Description          *string    `json:"description,omitempty" gorm:"type:text"`
}
