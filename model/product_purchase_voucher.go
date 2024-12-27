package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProductPurchaseVoucher struct {
	basic.Base
	basic.DataOwner
	ProductPurchaseVoucherAPI
}

type ProductPurchaseVoucherAPI struct {
	ProductPurchaseID *uuid.UUID `json:"product_purchase_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	VoucherID         *uuid.UUID `json:"voucher_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
}
