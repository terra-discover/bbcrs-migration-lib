package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationTaxCategory struct {
	basic.Base
	basic.DataOwner
	ReservationTaxCategoryAPI
	Reservation *Reservation `json:"reservation,omitempty" gorm:"foreignKey:ReservationID;references:ID"`
}

type ReservationTaxCategoryAPI struct {
	ReservationID   *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null;"`  // The reference to the reservation
	TaxCategoryID   *uuid.UUID `json:"tax_category_id,omitempty" gorm:"type:varchar(36);not null;"` // The reference to the tax category.
	TaxCategoryCode *string    `json:"tax_category_code,omitempty" gorm:"type:varchar(36)"`         // The code of the tax category.
	TaxCategoryName *string    `json:"tax_category_name,omitempty" gorm:"type:varchar(256)"`        // The name of the tax category.
	Description     *string    `json:"description,omitempty" gorm:"type:text"`
}
