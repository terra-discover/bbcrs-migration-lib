package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProductPurchaseReward struct {
	basic.Base
	basic.DataOwner
	ProductPurchaseRewardAPI
}

type ProductPurchaseRewardAPI struct {
	ProductPurchaseID          *uuid.UUID `json:"product_purchase_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	RewardBalanceTransactionID *uuid.UUID `json:"reward_balance_transaction_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	VoucherTransactionID       *uuid.UUID `json:"voucher_transaction_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Comment                    *string    `json:"comment,omitempty" gorm:"type:text"`
	Description                *string    `json:"description,omitempty" gorm:"type:text"`
}
