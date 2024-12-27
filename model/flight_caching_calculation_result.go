package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingCalculationResult struct {
	basic.Base
	basic.DataOwner
	SessionID                                    *uuid.UUID                    `json:"session_id,omitempty" gorm:"type:varchar(36)"`
	ComponentName                                *string                       `json:"component_name,omitempty" gorm:"type:varchar(20)" example:"flight/addon/seat"`
	FareDetailID                                 *uuid.UUID                    `json:"fare_detail_id,omitempty" gorm:"type:varchar(36)"`
	PricePerPassenger                            *float64                      `json:"price_per_passenger,omitempty"`
	PricePerPassengerExcludeTax                  *float64                      `json:"price_per_passenger_exclude_tax,omitempty"`
	PartnerBasePrice                             *float64                      `json:"partner_base_price,omitempty"`
	PartnerTotalTax                              *float64                      `json:"partner_total_tax,omitempty"`
	PartnerPrice                                 *float64                      `json:"partner_price,omitempty"`
	CommissionPerPassenger                       *float64                      `json:"commission_per_passenger,omitempty"`
	PricePerPassengerExcludeCommission           *float64                      `json:"price_per_passenger_exclude_commission,omitempty"`
	PricePerPassengerExcludeCommissionExcludeTax *float64                      `json:"price_per_passenger_exclude_commission_exclude_tax,omitempty"`
	NettToAgentPerPassenger                      *float64                      `json:"nett_to_agent_per_passenger,omitempty"`
	CommissionRetainedPerPassenger               *float64                      `json:"commission_retained_per_passenger,omitempty"`
	ServiceFeePerPassenger                       *float64                      `json:"service_fee_per_passenger,omitempty"`
	ServiceFeeIncluded                           *bool                         `json:"service_fee_included,omitempty"`
	ServiceFeeIsVisible                          *bool                         `json:"service_fee_is_visible,omitempty"`
	ServiceFeeIsSplitInvoice                     *bool                         `json:"service_fee_is_split_invoice,omitempty"`
	MarkupPerPassenger                           *float64                      `json:"markup_per_passenger,omitempty"`
	ServiceFeeTaxPerPassenger                    *float64                      `json:"service_fee_tax_per_passenger,omitempty"`
	TaxFeePerPassenger                           *float64                      `json:"tax_fee_per_passenger,omitempty"`
	MDRFeePerPassenger                           *float64                      `json:"mdr_fee_per_passenger,omitempty"`
	RevalidateFeePerPassenger                    *float64                      `json:"revalidate_fee_per_passenger,omitempty"`
	ReissueFeePerPassenger                       *float64                      `json:"reissue_fee_per_passenger,omitempty"`
	UsedCommissionClaimID                        *uuid.UUID                    `json:"used_commission_claim_id,omitempty" gorm:"type:varchar(36)"`
	UsedMarkupRateID                             *uuid.UUID                    `json:"used_markup_rate_id,omitempty" gorm:"type:varchar(36)"`
	UsedServiceFeeRateID                         *uuid.UUID                    `json:"used_service_fee_rate_id,omitempty" gorm:"type:varchar(36)"`
	UsedTaxRateID                                *uuid.UUID                    `json:"used_tax_rate_id,omitempty" gorm:"type:varchar(36)"`
	UsedServiceFeeTaxRateID                      *uuid.UUID                    `json:"used_service_fee_tax_rate_id,omitempty" gorm:"type:varchar(36)"`
	UsedMDRProcessingFeeRateID                   *uuid.UUID                    `json:"used_mdr_processing_fee_rate_id,omitempty" gorm:"type:varchar(36)"`
	UsedRevalidateProcessingFeeRateID            *uuid.UUID                    `json:"used_revalidate_processing_fee_rate_id,omitempty" gorm:"type:varchar(36)"`
	UsedReissueProcessingFeeRateID               *uuid.UUID                    `json:"used_reissue_processing_fee_rate_id,omitempty" gorm:"type:varchar(36)"`
	FareDetailTaxes                              *[]FlightCachingFareDetailTax `json:"fare_detail_taxes" gorm:"-"`
}
