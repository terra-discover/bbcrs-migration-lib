package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyAccountVoucher Operation Schedule
type LoyaltyAccountVoucher struct {
	basic.Base
	basic.DataOwner
	LoyaltyAccountVoucherAPI
}

// LoyaltyAccountVoucherAPI Operation Schedule API
type LoyaltyAccountVoucherAPI struct {
	LoyaltyAccountID *uuid.UUID       `json:"loyalty_account_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml // Operation Schedule Name
	VoucherID        *uuid.UUID       `json:"voucher_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`         // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml // Operation Schedule Name
	RedeemedAt       *strfmt.DateTime `json:"redeemed_at,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`            // Redeemed At
	UsedAt           *strfmt.DateTime `json:"used_at,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`                // Used At                                                                                                                 // Description
}
