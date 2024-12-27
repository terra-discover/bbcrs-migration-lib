package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// ProductPurchase  Product Purchase
type ProductPurchase struct {
	basic.Base
	basic.DataOwner
	ProductPurchaseAPI
	ProposalProductPurchase    *ProposalProductPurchase    `json:"proposal_product_purchase,omitempty" gorm:"foreignKey:ProductPurchaseID;references:ID"`    // reference ProposalProductPurchase By ProductPurchaseID
	ReservationProductPurchase *ReservationProductPurchase `json:"reservation_product_purchase,omitempty" gorm:"foreignKey:ProductPurchaseID;references:ID"` // reference ReservationProductPurchase By ProductPurchaseID
}

// ProductPurchaseAPI Product Purchase API
type ProductPurchaseAPI struct {
	ProductID                        *uuid.UUID       `json:"product_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AirlineID                        *uuid.UUID       `json:"airline_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	HotelID                          *uuid.UUID       `json:"hotel_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	BusinessPartnerID                *uuid.UUID       `json:"business_partner_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	AirTravelerID                    *uuid.UUID       `json:"air_traveler_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightSegmentID                  *uuid.UUID       `json:"flight_segment_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	PurchaseStatus                   *int             `json:"purchase_status,omitempty" gorm:"type:smallint"`
	RatePlanID                       *uuid.UUID       `json:"rate_plan_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	OfferID                          *uuid.UUID       `json:"offer_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Quantity                         *int             `json:"quantity,omitempty" gorm:"type:smallint"`
	EffectiveDate                    *strfmt.DateTime `json:"effective_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	ExpireDate                       *strfmt.DateTime `json:"expire_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	TicketStatus                     *int             `json:"ticket_status,omitempty" gorm:"type:smallint"`
	TicketDocumentNumber             *string          `json:"ticket_document_number,omitempty" gorm:"type:varchar(64)"`
	BasePriceCurrencyID              *uuid.UUID       `json:"base_price_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	BasePriceAmount                  *float64         `json:"base_price_amount,omitempty"`
	EquivalentPriceCurrencyID        *uuid.UUID       `json:"equivalent_price_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	EquivalentPriceAmount            *float64         `json:"equivalent_price_amount,omitempty"`
	TotalTaxCurrencyID               *uuid.UUID       `json:"total_tax_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	TotalTaxAmount                   *float64         `json:"total_tax_amount,omitempty"`
	TotalFeeCurrencyID               *uuid.UUID       `json:"total_fee_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	TotalFeeAmount                   *float64         `json:"total_fee_amount,omitempty"`
	TotalPriceCurrencyID             *uuid.UUID       `json:"total_price_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	TotalPriceAmount                 *float64         `json:"total_price_amount,omitempty"`
	CurrencyID                       *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	TotalAmountBeforeTax             *float64         `json:"total_amount_before_tax,omitempty"`
	TotalAmountAfterTax              *float64         `json:"total_amount_after_tax,omitempty"`
	CommissionPercent                *float64         `json:"commission_percent,omitempty"`
	CommissionCurrencyID             *uuid.UUID       `json:"commission_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	CommissionAmount                 *float64         `json:"commission_amount,omitempty"`
	CurrencyConversionFromCurrencyID *uuid.UUID       `json:"currency_conversion_from_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	CurrencyConversionToCurrencyID   *uuid.UUID       `json:"currency_conversion_to_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	CurrencyConversionMultiplyRate   *float64         `json:"currency_conversion_multiply_rate,omitempty"`
	BasePenaltyCurrencyID            *uuid.UUID       `json:"base_penalty_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	BasePenaltyAmount                *float64         `json:"base_penalty_amount,omitempty"`
	EquivalentPenaltyCurrencyID      *uuid.UUID       `json:"equivalent_penalty_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	EquivalentPenaltyAmount          *float64         `json:"equivalent_penalty_amount,omitempty"`
	TotalPenaltyTaxCurrencyID        *uuid.UUID       `json:"total_penalty_tax_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	TotalPenaltyTaxAmount            *float64         `json:"total_penalty_tax_amount,omitempty"`
	TotalPenaltyCurrencyID           *uuid.UUID       `json:"total_penalty_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	TotalPenaltyAmount               *float64         `json:"total_penalty_amount,omitempty"`
	PenaltyCurrencyID                *uuid.UUID       `json:"penalty_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	TotalPenaltyAmountBeforeTax      *float64         `json:"total_penalty_amount_before_tax,omitempty"`
	TotalPenaltyAmountAfterTax       *float64         `json:"total_penalty_amount_after_tax,omitempty"`
	CancellationDeadline             *strfmt.DateTime `json:"cancellation_deadline,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	CancelledAt                      *strfmt.DateTime `json:"cancelled_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	CancellationReason               *string          `json:"cancellation_reason" gorm:"text"`
}

type ProductPurchaseTicket struct {
	ID                           *uuid.UUID `json:"id,omitempty"`
	ProductCategoryName          *string    `json:"product_category_name,omitempty"`
	ProductCategoryCode          *string    `json:"product_category_code,omitempty"`
	AirlineName                  *string    `json:"airline_name,omitempty"`
	ProductCode                  *string    `json:"product_code,omitempty"`
	ProductName                  *string    `json:"product_name,omitempty"`
	TicketStatus                 *int       `json:"ticket_status,omitempty"`
	TicketDocumentNumber         *string    `json:"ticket_document_number,omitempty"`
	MarketingAirlineCode         *string    `json:"marketing_airline_code,omitempty"`
	MarketingAirlineFlightNumber *string    `json:"marketing_airline_flight_number,omitempty"`
	OperatingAirlineCode         *string    `json:"operating_airline_code,omitempty"`
	OperatingAirlineFlightNumber *string    `json:"operating_airline_flight_number,omitempty"`
	DepartureAirportLocationCode *string    `json:"departure_airport_location_code,omitempty"`
	ArrivalAirportLocationCode   *string    `json:"arrival_airport_location_code,omitempty"`
	DepartureAirportTerminal     *string    `json:"departure_airport_terminal,omitempty"`
	ArrivalAirportTerminal       *string    `json:"arrival_airport_terminal,omitempty"`
	TotalAmountBeforeTax         *float64   `json:"total_amount_before_tax,omitempty"`
	TotalAmountAfterTax          *float64   `json:"total_amount_after_tax,omitempty"`
	ReservationID                *uuid.UUID `json:"reservation_id,omitempty"`
	ProductID                    *uuid.UUID `json:"product_id,omitempty"`
	ProductCategoryID            *uuid.UUID `json:"product_category_id,omitempty"`
	AirlineID                    *uuid.UUID `json:"airline_id,omitempty"`
	AirTravelerID                *uuid.UUID `json:"air_traveler_id,omitempty"`
	FlightSegmentID              *uuid.UUID `json:"flight_segment_id,omitempty"`
	PersonID                     *uuid.UUID `json:"person_id,omitempty"`
}

// GetProductPurchaseTicket by query filter
func (data *ProductPurchase) GetProductPurchaseTicket(tx *gorm.DB, reservationID *uuid.UUID) *[]ProductPurchaseTicket {
	var products []ProductPurchaseTicket
	tx.Model(&data).
		Select(`
		product_purchase.id,
		product_category.product_category_code,
		product_category.product_category_name,
		airline.airline_name,
		product.product_code,
		product.product_name,
		product_purchase.ticket_status,
		product_purchase.ticket_document_number,
		flight_segment.marketing_airline_code as marketing_airline_code,
		flight_segment.marketing_airline_flight_number as marketing_airline_flight_number,
		flight_segment.operating_airline_code as operating_airline_code,
		flight_segment.operating_airline_flight_number as operating_airline_flight_number,
		flight_segment.departure_airport_location_code as departure_airport_location_code,
		flight_segment.arrival_airport_location_code as arrival_airport_location_code,
		flight_segment.departure_airport_terminal as departure_airport_terminal,
		flight_segment.arrival_airport_terminal as arrival_airport_terminal,
		product_purchase.total_amount_before_tax as total_amount_before_tax,
		product_purchase.total_amount_after_tax as total_amount_after_tax,
		reservation_product_purchase.reservation_id,
		product_purchase.product_id,
		product_category.id as product_category_id,
		product_purchase.airline_id,
		product_purchase.air_traveler_id,
		product_purchase.flight_segment_id,
		air_traveler.person_id
	`).
		Joins("LEFT JOIN product ON product_purchase.product_id = product.id").
		Joins("LEFT JOIN product_category_product ON product_category_product.product_id = product.id").
		Joins("LEFT JOIN product_category ON product_category_product.product_category_id = product_category.id").
		Joins("LEFT JOIN airline ON product_purchase.airline_id = airline.id").
		Joins("LEFT JOIN reservation_product_purchase ON reservation_product_purchase.product_purchase_id = product_purchase.id").
		Joins("LEFT JOIN air_traveler ON product_purchase.air_traveler_id = air_traveler.id").
		Joins("LEFT JOIN flight_segment ON product_purchase.flight_segment_id = flight_segment.id").
		Where("reservation_product_purchase.reservation_id = ?", reservationID).
		Find(&products)
	return &products
}
