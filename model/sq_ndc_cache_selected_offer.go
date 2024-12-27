package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SqNdcCacheSelectedOffer
type SqNdcCacheSelectedOffer struct {
	basic.Base
	basic.DataOwner
	SessionID                *uuid.UUID                   `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                                                                                            // session ID
	OfferID                  *string                      `json:"offer_id,omitempty" gorm:"type:varchar(60);index:idx_offer_id,unique,where:deleted_at is null;not null" default:"SP2F-7761831007179611393-1"` // OfferID
	OfferExpirationDateTime  *string                      `json:"offer_expiration_date_time,omitempty" gorm:"type:timestamptz"`                                                                                // OfferExpirationDateTime
	PaymentTimeLimitDateTime *string                      `json:"payment_time_limit_date_time,omitempty" gorm:"type:timestamptz"`                                                                              // PaymentTimeLimitDateTime
	OfferItem                *SqNdcCacheSelectedOfferItem `json:"offer_item,omitempty" gorm:"foreignKey:OfferID;references:OfferID"`
}

// SqNdcCacheSelectedOfferItem
type SqNdcCacheSelectedOfferItem struct {
	basic.Base
	basic.DataOwner
	SessionID          *uuid.UUID                           `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                                                                                                          // session ID
	OfferItemID        *string                              `json:"offer_item_id,omitempty" gorm:"type:varchar(265);index:idx_offer_item_id,unique,where:deleted_at is null;not null" default:"SP1F-15505250137536741898-1-1"` // OfferItemID
	OfferID            *string                              `json:"offer_id,omitempty" gorm:"type:varchar(60)" default:"SP2F-7761831007179611393-1"`                                                                           // OfferID
	PaxID              *string                              `json:"pax_id,omitempty" gorm:"type:varchar(60)" default:"PAX1"`
	Type               *string                              `json:"type,omitempty" gorm:"type:varchar(30)"` // eg : seat,offer,addons
	BaseAmount         *float64                             `json:"base_amount,omitempty"`                  // price information
	TotalTaxAmount     *float64                             `json:"total_tax_amount,omitempty"`             // price information
	TotalAmount        *float64                             `json:"total_amount,omitempty"`                 // price information
	PrevBaseAmount     *float64                             `json:"prev_base_amount,omitempty"`             // price information
	PrevTotalTaxAmount *float64                             `json:"prev_total_tax_amount,omitempty"`        // price information
	PrevTotalAmount    *float64                             `json:"prev_total_amount,omitempty"`            // price information
	PenaltyPrice       *float64                             `json:"penalty_price,omitempty"`                // Represents the penalty price fee, if applicable | used for module modify flight after issued only
	PriceDifference    *float64                             `json:"price_difference,omitempty"`             // Difference between new and previous prices (exclude penalty price). Positive: Upgrade, Negative: Downgrade | used for module modify flight before or after issued only
	PriceAmountDue     *float64                             `json:"price_amount_due,omitempty"`             // Represents the amount due to pay for the current pricing |  used for module modify flight after issued only
	RouteID            *uuid.UUID                           `json:"route_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	FareID             *uuid.UUID                           `json:"fare_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	PersonID           *uuid.UUID                           `json:"person_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	FareDetail         *[]SqNdcCacheSelectedOfferFareDetail `json:"fare_detail,omitempty" gorm:"foreignKey:OfferItemID;references:OfferItemID"`
}

type SqNdcCacheSelectedOfferFareDetail struct {
	basic.Base
	basic.DataOwner
	SessionID          *uuid.UUID                                  `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	OfferItemID        *string                                     `json:"offer_item_id,omitempty" gorm:"type:varchar(60)" default:"SP1F-15505250137536741898-1-1"`
	PassengerRefs      *string                                     `json:"passenger_refs,omitempty" gorm:"type:varchar(60)"`
	BaseAmount         *float64                                    `json:"base_amount,omitempty"`           // price information
	TotalTaxAmount     *float64                                    `json:"total_tax_amount,omitempty"`      // price information
	TotalAmount        *float64                                    `json:"total_amount,omitempty"`          // price information
	PrevBaseAmount     *float64                                    `json:"prev_base_amount,omitempty"`      // price information
	PrevTotalTaxAmount *float64                                    `json:"prev_total_tax_amount,omitempty"` // price information
	PrevTotalAmount    *float64                                    `json:"prev_total_amount,omitempty"`     // price information
	PenaltyPrice       *float64                                    `json:"penalty_price,omitempty"`         // Represents the penalty price fee, if applicable | used for module modify flight after issued only
	PriceDifference    *float64                                    `json:"price_difference,omitempty"`      // Difference between new and previous prices (exclude penalty price). Positive: Upgrade, Negative: Downgrade | used for module modify flight before or after issued only
	PriceAmountDue     *float64                                    `json:"price_amount_due,omitempty"`      // Represents the amount due to pay for the current pricing |  used for module modify flight after issued only
	BaggageAllowance   *[]SqNdcCacheSelectedOfferFareDetailBaggage `json:"baggage_allowance" gorm:"foreignKey:FareDetailID;references:ID"`
	FareTax            *[]SqNdcCacheSelectedOfferFareDetailTax     `json:"fare_tax,omitempty"  gorm:"foreignKey:FareDetailID;references:ID"`
	FareComponent      *[]SqNdcCacheSelectedFareComponent          `json:"fare_component" gorm:"foreignKey:FareDetailID;references:ID"`
}

type SqNdcCacheSelectedFareComponent struct {
	basic.Base
	basic.DataOwner
	SessionID                *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	FareDetailID             *uuid.UUID `json:"fare_detail_id,omitempty" gorm:"type:varchar(36)"`
	PriceClassRef            *string    `json:"price_class_ref,omitempty" gorm:"type:varchar(60)" default:"FF31LON"`
	SegmentRefs              *string    `json:"segment_refs,omitempty" gorm:"type:varchar(60)" default:"SEG2"`
	FareBasisCode            *string    `json:"fare_basis_code,omitempty" gorm:"type:varchar(60)" default:"H14GBRL"`                  // FareBasisCode.Code
	FareRulesRemarksCategory *string    `json:"fare_rules_remarks_category,omitempty" gorm:"type:varchar(60)" default:"FT"`           // FareRulesRemarks
	FareRulesRemarksText     *string    `json:"fare_rules_remarks_text,omitempty" gorm:"type:varchar(256)" default:"NEGOTIATED FARE"` // FareRulesRemarks
	Rbd                      *string    `json:"rbd,omitempty" gorm:"type:varchar(10)" default:"H"`                                    // FareBasis
	CabinTypeCode            *string    `json:"cabin_type_code,omitempty" gorm:"type:varchar(10)" default:"M"`                        // CabinType
	CabinTypeName            *string    `json:"cabin_type_name,omitempty" gorm:"type:varchar(256)" default:"ECO"`                     // CabinType
}

type SqNdcCacheSelectedOfferFareDetailTax struct {
	basic.Base
	basic.DataOwner
	SessionID    *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	FareDetailID *uuid.UUID `json:"fare_detail_id,omitempty" gorm:"type:varchar(36)"`
	TaxCode      *string    `json:"tax_code,omitempty" gorm:"type:varchar(5)"`
	Amount       *float64   `json:"amount,omitempty"`
	Nation       *string    `json:"nation,omitempty" gorm:"type:varchar(5)"`
	CurrencyCode *string    `json:"currency_code,omitempty" gorm:"type:varchar(36)"`
	CurrencyName *string    `json:"currencey_name,omitempty" gorm:"type:varchar(50)"`
}

type SqNdcCacheSelectedOfferFareDetailBaggage struct {
	basic.Base
	basic.DataOwner
	SessionID            *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	FareDetailID         *uuid.UUID `json:"fare_detail_id,omitempty" gorm:"type:varchar(36)"`
	MaximumWeightMeasure *float64   `json:"maximum_weight_measure,omitempty"`
	BaggageAllowanceID   *string    `json:"baggage_allowance_id,omitempty" gorm:"type:varchar(10)"`
	UnitCode             *string    `json:"unit_code,omitempty" gorm:"type:varchar(20)"`
}

// GetSelectedOffer
func (data *SqNdcCacheSelectedOffer) GetSelectedOffer(tx *gorm.DB, sessionID *uuid.UUID, types string) {
	tx.Model(&SqNdcCacheSelectedOffer{}).
		Select(`
			sq_ndc_cache_selected_offer.id,
			sq_ndc_cache_selected_offer.session_id,
			sq_ndc_cache_selected_offer.offer_id,
			sq_ndc_cache_selected_offer.offer_expiration_date_time,
			sq_ndc_cache_selected_offer.payment_time_limit_date_time,
			item.offer_item_id "OfferItem__offer_item_id",
			item.offer_id "OfferItem__offer_id",
			item.pax_id "OfferItem__pax_id",
			item.type "OfferItem__type",
			item.base_amount "OfferItem__base_amount",
			item.total_tax_amount "OfferItem__total_tax_amount"
		`).
		Where("sq_ndc_cache_selected_offer.session_id = ?", sessionID).
		Joins(`LEFT JOIN (
			SELECT a.* FROM 
			"sq_ndc_cache_selected_offer_item" a
			WHERE a.type = '` + types + `'
		) item ON item.offer_id = sq_ndc_cache_selected_offer.offer_id`).
		Take(&data)
}

// GetTotalAmount
func (s *SqNdcCacheSelectedOfferItem) GetTotalAmount(tx *gorm.DB, sessionID *uuid.UUID) (total string) {
	tx.Model(&s).Select("SUM(total_amount)").Where("session_id = ?", sessionID).Row().Scan(&total)
	return
}

// GetPriceAmountDue
func (s *SqNdcCacheSelectedOfferItem) GetPriceAmountDue(tx *gorm.DB, sessionID *uuid.UUID) (total string) {
	tx.Model(&s).Select("SUM(price_amount_due)").Where("session_id = ?", sessionID).Row().Scan(&total)
	return
}

// SetOfferItem sets a selected offer item in the SqNdcCacheSelectedOffer table.
//
// Args:
//
//	tx (*gorm.DB): The gorm transaction.
//	sessionID (*uuid.UUID): The session ID.
//	personID (*uuid.UUID): The person ID.
//	routeID (*uuid.UUID): The route ID.
//	fareID (*uuid.UUID): The fare ID.
//	paxID (*string): The passenger ID.
//	offerItemID (*string): The offer item ID.
//	types (*string): The types.
//	TotalAmount (*float64): The total amount.
//
// Returns:
//
//	error: The error, if any.
func (s *SqNdcCacheSelectedOffer) SetOfferItem(
	tx *gorm.DB,
	sessionID,
	personID,
	routeID,
	fareID *uuid.UUID,
	paxID,
	offerItemID *string,
	types *string,
	TotalAmount *float64,
) error {
	s.GetSelectedOffer(tx, sessionID, "offer")

	offerItem := SqNdcCacheSelectedOfferItem{}
	notEmptyAllCond := !lib.IsEmptyUUIDPtr(sessionID) && !lib.IsEmptyUUIDPtr(routeID) && !lib.IsEmptyUUIDPtr(fareID) && !lib.IsEmptyUUIDPtr(personID)
	if notEmptyAllCond {
		tx.Unscoped().Where(tx.Where(SqNdcCacheSelectedOfferItem{
			SessionID: sessionID,
			Type:      types,
			RouteID:   routeID,
			FareID:    fareID,
			PersonID:  personID,
		})).Delete(&offerItem)
	}

	offerItem.SessionID = sessionID
	offerItem.OfferID = s.OfferID
	offerItem.PaxID = paxID
	offerItem.Type = types
	offerItem.OfferItemID = offerItemID
	offerItem.TotalAmount = TotalAmount
	offerItem.RouteID = routeID
	offerItem.FareID = fareID
	offerItem.PersonID = personID

	if err := tx.Create(&offerItem); err != nil {
		return err.Error
	}
	return nil
}
