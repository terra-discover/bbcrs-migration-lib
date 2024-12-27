package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreCacheAddons struct {
	basic.Base
	basic.DataOwner
	SessionID       *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	PassengerID     *string    `json:"passenger_id,omitempty" gorm:"type:varchar(20)"`
	PassengerRef    *string    `json:"passenger_ref,omitempty" gorm:"type:varchar(20)"`
	FirstName       *string    `json:"first_name,omitempty" gorm:"type:varchar(100)"`
	LastName        *string    `json:"last_name,omitempty" gorm:"type:varchar(100)"`
	Segment         *string    `json:"segment,omitempty" gorm:"type:varchar(20)"`
	AddonType       *string    `json:"addon_type,omitempty" gorm:"type:varchar(50)"`
	AddonCode       *string    `json:"addon_code,omitempty" gorm:"type:varchar(20)"`
	AddonTitle      *string    `json:"addon_title,omitempty" gorm:"type:varchar(100)"`
	AddonPrice      *float64   `json:"addon_price,omitempty" gorm:"type:float8"`
	AddonIsSelected *bool      `json:"addon_is_selected,omitempty" gorm:"type:boolean;default:false"`
	OfferUUID       *uuid.UUID `json:"offer_uuid,omitempty" gorm:"type:varchar(36)"`
	OfferID         *string    `json:"offer_id,omitempty" gorm:"type:varchar(50)"`
	OfferItemID     *string    `json:"offer_item,omitempty" gorm:"type:varchar(50)"`

	OfferAncillaryRef *string `json:"offer_ancillary_ref,omitempty" gorm:"type:varchar(50)"`
	OfferRuleCode     *string `json:"offer_rule_code,omitempty" gorm:"type:varchar(50)"`
	OfferRuleMethod   *string `json:"offer_rule_method,omitempty" gorm:"type:varchar(50)"`
	AncillaryID       *string `json:"ancillary_id,omitempty" gorm:"type:varchar(50)"`
	AncillaryDefRef   *string `json:"ancillary_def_ref,omitempty" gorm:"type:varchar(50)"`
	AncillaryDefID    *string `json:"ancillary_def_id,omitempty" gorm:"type:varchar(50)"`

	SegmentText    *string `json:"segment_text,omitempty" gorm:"type:text"`
	OfferText      *string `json:"offer_text,omitempty" gorm:"type:text"`
	AncillaryText  *string `json:"ancillary_text,omitempty" gorm:"type:text"`
	DefinitionText *string `json:"definition_text,omitempty" gorm:"type:text"`

	FlightNumber        *string    `json:"flight_number,omitempty" gorm:"type:varchar(20)"`
	ProductCategoryID   *uuid.UUID `json:"product_category_id,omitempty" gorm:"type:varchar(36)"`
	ProductCategoryCode *string    `json:"product_category_code,omitempty" gorm:"type:varchar(20)"`
	ProductCategoryName *string    `json:"product_category_name,omitempty" gorm:"type:varchar(100)"`
	Weight              *float64   `json:"weight,omitempty" gorm:"type:decimal(10,2)"`
}
