package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationProductVoucherType struct {
	basic.Base
	basic.DataOwner
	ReservationProductVoucherTypeAPI
}

type ReservationProductVoucherTypeAPI struct {
	ReservationID   *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProductID       *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	VoucherTypeID   *uuid.UUID `json:"voucher_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	VoucherTypeCode *string    `json:"voucher_type_code,omitempty" gorm:"type:varchar(36)"`
	VoucherTypeName *string    `json:"voucher_type_name,omitempty" gorm:"type:varchar(256)"`
}
