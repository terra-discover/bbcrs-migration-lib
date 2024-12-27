package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerSetting Integration Partner Setting
type IntegrationPartnerSetting struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerSettingAPI
	IntegrationPartner *IntegrationPartner `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"` // Integration Partner
	Currency           *Currency           `json:"currency" gorm:"foreignKey:ForeignExchangeBaseCurrencyID;references:ID"`   // Currency
}

// IntegrationPartnerSettingAPI Integration Partner Setting API
type IntegrationPartnerSettingAPI struct {
	IntegrationPartnerID          *uuid.UUID       `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid"`                                       // The reference to the integration partner.
	CanBookSeat                   *bool            `json:"can_book_seat,omitempty"`                                                                                   // Indicates whether allowing to book seat.
	CanBookSsr                    *bool            `json:"can_book_ssr,omitempty"`                                                                                    // Indicates whether allowing to book special service request.
	BookSeatCanHoldTicket         *bool            `json:"book_seat_can_hold_ticket,omitempty"`                                                                       // Indicates whether allowing to book now and issue flight ticket later when booking seat.
	BookSsrCanHoldTicket          *bool            `json:"book_ssr_can_hold_ticket,omitempty"`                                                                        // Indicates whether allowing to book now and issue flight ticket later when booking special service request.
	CanCancelTicket               *bool            `json:"can_cancel_ticket,omitempty"`                                                                               // Indicates whether allowing to cancel or void issued ticket.
	OnlyAvailableRoom             *bool            `json:"only_available_room,omitempty"`                                                                             // Indicates whether searching availability for available room only.
	OnlyDirectContractTariff      *bool            `json:"only_direct_contract_tariff,omitempty"`                                                                     // Indicates whether searching availability for direct contract tariff only.
	OnlyBestPrice                 *bool            `json:"only_best_price,omitempty"`                                                                                 // Indicates whether searching availability for best price/ lowest rate category only.
	RoomRatePerCategory           *bool            `json:"room_rate_per_category,omitempty"`                                                                          // Indicates whether searching availability with full category rate.
	DailyRateBreakdown            *bool            `json:"daily_rate_breakdown,omitempty"`                                                                            // Indicates whether searching availability with daily rate breakdown.
	RoomRateWithHotelSupplier     *bool            `json:"room_rate_with_hotel_supplier,omitempty"`                                                                   // Indicates whether room rates shown with hotel supplier.
	NeedNationalIdentityNumber    *bool            `json:"need_national_identity_number,omitempty"`                                                                   // Indicates whether requires national identity number.
	NeedPassport                  *bool            `json:"need_passport,omitempty"`                                                                                   // Indicates whether requires passport for international travel.
	NeedBirthDate                 *bool            `json:"need_birth_date,omitempty"`                                                                                 // Indicates whether requires birth date.
	HasNationality                *bool            `json:"has_nationality,omitempty"`                                                                                 // Indicates whether has nationality.
	NeedNationality               *bool            `json:"need_nationality,omitempty"`                                                                                // Indicates whether requires nationality.
	HasFrequentFlyer              *bool            `json:"has_frequent_flyer,omitempty"`                                                                              // Indicates whether has frequent flyer program.
	ForeignExchangeBaseCurrencyID *uuid.UUID       `json:"foreign_exchange_base_currency_id,omitempty" swaggertype:"string" format:"uuid"`                            // The reference to the base currency for foreign exchange.
	HotelCodeTimestamp            *strfmt.DateTime `json:"hotel_code_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`            // Timestamp of the last update of hotel code (in UTC).
	HotelInformationTimestamp     *strfmt.DateTime `json:"hotel_information_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`     // Timestamp of the last update of hotel information (in UTC).
	CorporateInformationTimestamp *strfmt.DateTime `json:"corporate_information_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // Timestamp of the last update of corporate information (in UTC).
	PartnerCabinTypeEnabled       *bool            `json:"partner_cabin_type_enabled,omitempty"`                                                                      // Partner Cabin Type Enabled
	PartnerMealPlanTypeEnabled    *bool            `json:"partner_meal_plan_type_enabled,omitempty"`                                                                  // Partner Meal Plan Type Enabled
	PartnerHotelEnabled           *bool            `json:"partner_hotel_enabled,omitempty"`                                                                           // Partner Hotel Enabled
	PartnerCityEnabled            *bool            `json:"partner_city_enabled,omitempty"`                                                                            // Partner City Enabled
	PartnerCountryEnabled         *bool            `json:"partner_country_enabled,omitempty"`                                                                         // Partner Country Enabled
	PartnerStateProvinceEnabled   *bool            `json:"partner_state_province_enabled,omitempty"`                                                                  // Partner State Province Enabled
	PartnerCurrencyEnabled        *bool            `json:"partner_currency_enabled,omitempty"`                                                                        // Partner Currency Enabled
	PartnerFeeTaxTypeEnabled      *bool            `json:"partner_fee_tax_type_enabled,omitempty"`                                                                    // Partner Fee Tax Type Enabled
	PartnerCredentialEnabled      *bool            `json:"partner_credential_enabled,omitempty"`                                                                      // Partner Credential Enabled
	PartnerCorporateEnabled       *bool            `json:"partner_corporate_enabled,omitempty"`                                                                       // Partner Corporate Enabled
	PartnerEventEnabled           *bool            `json:"partner_event_enabled,omitempty"`                                                                           // Partner Event Enabled
	PartnerPaymentGatewayEnabled  *bool            `json:"partner_payment_gateway_enabled,omitempty"`                                                                 // Partner Payment Gateway Enabled
	PartnerHotelSupplierEnabled   *bool            `json:"partner_hotel_supplier_enabled,omitempty"`                                                                  // Partner Hotel Supplier Enabled
	MaxNameLength                 *int             `json:"max_name_length,omitempty" gorm:"default:29"`                                                               // Based on MoM, not documented in the specification. Added in 2023, Sept 12
}
