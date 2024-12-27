package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Air Itinerary Cancel Policy
type AirItineraryCancelPolicy struct {
	basic.Base
	basic.DataOwner
	AirItineraryCancelPolicyAPI
}

// AirItineraryCancelPolicyAPI
type AirItineraryCancelPolicyAPI struct {
	AirItineraryID        *uuid.UUID `json:"air_itinerary_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	CancelPolicyID        *uuid.UUID `json:"cancel_policy_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	CancelPolicyCode      *string    `json:"cancel_policy_code,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	CancelPolicyName      *string    `json:"cancel_policy_name,omitempty" gorm:"type:varchar(256)" swaggertype:"string"`
	OffsetTimeUnitID      *uuid.UUID `json:"offset_time_unit_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	BasisTypeID           *uuid.UUID `json:"basis_type_id,omitempty"  gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	IsTaxInclusive        *bool      `json:"is_tax_inclusive,omitempty"`
	IsServiceFeeInclusive *bool      `json:"is_service_fee_inclusive,omitempty"`
	IsFeeInclusive        *bool      `json:"is_fee_inclusive,omitempty"`
	NotCancellable        *bool      `json:"not_cancellable,omitempty"`
	NotRefundable         *bool      `json:"not_refundable,omitempty"`
	Description           *string    `json:"description,omitempty" gorm:"type:text"`
}
