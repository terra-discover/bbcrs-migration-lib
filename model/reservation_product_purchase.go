package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Reservation Product Purchase
type ReservationProductPurchase struct {
	basic.Base
	basic.DataOwner
	ReservationProductPurchaseAPI
	ProductPurchase *ProductPurchase `json:"product_purchase,omitempty" gorm:"foreignKey:ProductPurchaseID;references:ID"` // ProductPurchase
	Reservation     *Reservation     `json:"reservation,omitempty" gorm:"foreignKey:ReservationID;references:ID"`          // Reservation
}

// ReservationProductPurchaseAPI Proposal Product Purchase API
type ReservationProductPurchaseAPI struct {
	ReservationID     *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProductPurchaseID *uuid.UUID `json:"product_purchase_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
