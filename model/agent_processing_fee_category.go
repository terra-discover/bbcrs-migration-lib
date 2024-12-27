package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentProcessingFeeCategory struct
type AgentProcessingFeeCategory struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeCategory *ProcessingFeeCategory `json:"processing_fee_category,omitempty"` // Processing Fee Category
	AgentProcessingFeeCategoryAPI
}

// AgentProcessingFeeCategoryAPI struct
type AgentProcessingFeeCategoryAPI struct {
	AgentID                 *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;index:idx_agent_processing_category_id,unique,where:deleted_at is null"`                   // Agent Id
	ProcessingFeeCategoryID *uuid.UUID `json:"processing_fee_category_id,omitempty" gorm:"type:varchar(36);not null;index:idx_agent_processing_category_id,unique,where:deleted_at is null"` // Processing Fee Category Id
}

// AgentProcessingFeeCategoryDetail struct
type AgentProcessingFeeCategoryDetail struct {
	basic.Base
	basic.DataOwner
	AgentProcessingFeeCategoryDetailAPI
	Used                        *bool              `json:"used,omitempty"`
	DomesticReissue             *ProcessingFeeRate `json:"domestic_reissue,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticRevalidate          *ProcessingFeeRate `json:"domestic_revalidate,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticRefund              *ProcessingFeeRate `json:"domestic_refund,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticVoid                *ProcessingFeeRate `json:"domestic_void,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticFrp                 *ProcessingFeeRate `json:"domestic_frp,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticNonGds              *ProcessingFeeRate `json:"domestic_non_gds,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalReissue        *ProcessingFeeRate `json:"international_reissue,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalRevalidate     *ProcessingFeeRate `json:"international_revalidate,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalRefund         *ProcessingFeeRate `json:"international_refund,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalVoid           *ProcessingFeeRate `json:"international_void,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalFrp            *ProcessingFeeRate `json:"international_frp,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalNonGds         *ProcessingFeeRate `json:"international_non_gds,omitempty" gorm:"foreignKey:ID;references:ID"`
	OtherEmergencyService       *ProcessingFeeRate `json:"other_emergency_service,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticModifyHotelFee      *ProcessingFeeRate `json:"domestic_modify_hotel_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalModifyHotelFee *ProcessingFeeRate `json:"international_modify_hotel_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticHotelRefundFee      *ProcessingFeeRate `json:"domestic_hotel_refund_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalHotelRefundFee *ProcessingFeeRate `json:"international_hotel_refund_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticNonGdsHotelFee      *ProcessingFeeRate `json:"domestic_non_gds_hotel_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalNonGdsHotelFee *ProcessingFeeRate `json:"international_non_gds_hotel_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
}

// AgentProcessingFeeCategoryDetailAPI struct
type AgentProcessingFeeCategoryDetailAPI struct {
	ProcessingFeeCategoryName   *string                    `json:"processing_fee_category_name" example:"Flight Processing Fee 1" validate:"required"`
	ProcessingFeeCategoryID     *uuid.UUID                 `json:"processing_fee_category_id"`
	Description                 *string                    `json:"description" example:"Flight Processing Fee 1 Description"`
	DomesticReissue             *AgentProcessingFeeRateAPI `json:"domestic_reissue" gorm:"-"`
	DomesticRevalidate          *AgentProcessingFeeRateAPI `json:"domestic_revalidate" gorm:"-"`
	DomesticRefund              *AgentProcessingFeeRateAPI `json:"domestic_refund" gorm:"-"`
	DomesticVoid                *AgentProcessingFeeRateAPI `json:"domestic_void" gorm:"-"`
	DomesticFrp                 *AgentProcessingFeeRateAPI `json:"domestic_frp" gorm:"-"`
	DomesticNonGds              *AgentProcessingFeeRateAPI `json:"domestic_non_gds" gorm:"-"`
	InternationalReissue        *AgentProcessingFeeRateAPI `json:"international_reissue" gorm:"-"`
	InternationalRevalidate     *AgentProcessingFeeRateAPI `json:"international_revalidate" gorm:"-"`
	InternationalRefund         *AgentProcessingFeeRateAPI `json:"international_refund" gorm:"-"`
	InternationalVoid           *AgentProcessingFeeRateAPI `json:"international_void" gorm:"-"`
	InternationalFrp            *AgentProcessingFeeRateAPI `json:"international_frp" gorm:"-"`
	InternationalNonGds         *AgentProcessingFeeRateAPI `json:"international_non_gds" gorm:"-"`
	OtherEmergencyService       *AgentProcessingFeeRateAPI `json:"other_emergency_service" gorm:"-"`
	DomesticModifyHotelFee      *AgentProcessingFeeRateAPI `json:"domestic_modify_hotel_fee" gorm:"-"`
	DomesticHotelRefundFee      *AgentProcessingFeeRateAPI `json:"domestic_hotel_refund_fee" gorm:"-"`
	DomesticNonGdsHotelFee      *AgentProcessingFeeRateAPI `json:"domestic_non_gds_hotel_fee" gorm:"-"`
	InternationalModifyHotelFee *AgentProcessingFeeRateAPI `json:"international_modify_hotel_fee" gorm:"-"`
	InternationalHotelRefundFee *AgentProcessingFeeRateAPI `json:"international_hotel_refund_fee" gorm:"-"`
	InternationalNonGdsHotelFee *AgentProcessingFeeRateAPI `json:"international_non_gds_hotel_fee" gorm:"-"`
}

// AgentProcessingFeeRateAPI writeable struct
type AgentProcessingFeeRateAPI struct {
	ChargeTypeID   *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // Charge Type Id
	Percent        *float64   `json:"percent,omitempty" gorm:"type:decimal(19,4)"`                    // Percent
	Amount         *float64   `json:"amount,omitempty" gorm:"type:decimal(19,4)" example:"3127.14"`   // Amount
	IsTaxInclusive *bool      `json:"is_tax_inclusive,omitempty"`                                     // Is Tax Inclusive
}

// AgentProcessingFeeCategoryProcessingFeeRateDetailAPI struct
type AgentProcessingFeeCategoryProcessingFeeRateDetailAPI struct {
	ProductCategoryID *uuid.UUID                 `json:"product_category_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // Product Category Id
	Markup            *AgentProcessingFeeRateAPI `json:"markup" gorm:"-"`
}
