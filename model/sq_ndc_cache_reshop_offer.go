package model

import (
	"math"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SqNdcCacheReshopOffer struct {
	basic.Base
	basic.DataOwner
	SessionID                   *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	OfferID                     *string    `json:"offer_id,omitempty" gorm:"varchar(100)"`
	OwnerCode                   *string    `json:"owner_code,omitempty" gorm:"varchar(5)"`
	DescID                      *string    `json:"desc_id,omitempty" gorm:"varchar(5)"`
	DescText                    *string    `json:"desc_text,omitempty" gorm:"varchar(50)"`
	OrderItemRefID              *string    `json:"order_item_ref_id,omitempty" gorm:"varchar(50)"`
	PenaltyDifferentialAmount   *float64   `json:"penalty_differential_amount,omitempty"`
	DifferentialAmountDueAmount *float64   `json:"differential_amount_due_amount,omitempty"`
	CurrencyCode                *string    `json:"currency_code,omitempty"`
}

func (data *SqNdcCacheReshopOffer) GetTotalExpectedAmountRefund(db *gorm.DB, sessionID *uuid.UUID) (total string) {
	reshop := []SqNdcCacheReshopOffer{}
	db.Model(&reshop).Where(`session_id = ?`, sessionID).
		Where(`LOWER(desc_text) = ?`, `refund`).
		Find(&reshop)
	refund := 0.00
	for _, val := range reshop {
		if val.DifferentialAmountDueAmount != nil {
			refund += math.Abs(*val.DifferentialAmountDueAmount)
		}
	}

	total = lib.FloatToStr(refund, 2)
	return
}
