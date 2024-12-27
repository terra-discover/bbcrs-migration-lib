package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationBusinessPartner struct {
	basic.Base
	basic.DataOwner
	ReservationBusinessPartnerAPI
}

type ReservationBusinessPartnerAPI struct {
	ReservationID     *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	BusinessPartnerID *uuid.UUID `json:"business_partner_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
