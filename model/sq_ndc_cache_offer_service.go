package model

import (
	"log"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SqNdcCacheService table struct
type SqNdcCacheOfferService struct {
	basic.Base
	basic.DataOwner
	SessionID              *uuid.UUID         `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // SessionID
	OfferID                *string            `json:"offer_id,omitempty" gorm:"type:varchar(50)"`
	OfferItemID            *string            `json:"offer_item_id,omitempty" gorm:"type:varchar(60)" default:"SP1F-17622158948335060224-1-2"` // OfferItemID
	ServiceDefinitionRefID *string            `json:"service_definition_ref_id,omitempty" gorm:"type:varchar(60)" default:"SRV1"`              // ServiceDefinitionRefID
	PaxRefID               *string            `json:"pax_ref_id,omitempty" gorm:"type:varchar(60)" default:"PAX1"`                             // PaxRefID
	PaxSegmentRefID        *string            `json:"pax_segment_ref_id,omitempty" gorm:"type:varchar(60)" default:"SEG5"`                     // PaxSegmentRefID
	BaseAmount             *float64           `json:"base_amount,omitempty" gorm:"type:decimal(22,2)"`                                         // price information
	TotalTaxAmount         *float64           `json:"total_tax_amount,omitempty" gorm:"type:decimal(22,2)"`                                    // price information
	TotalAmount            *float64           `json:"total_amount,omitempty" gorm:"type:decimal(22,2)"`                                        // price information
	DiscountAmount         *float64           `json:"discount_amount,omitempty" gorm:"type:decimal(22,2)"`                                     // price information
	CurrencyCode           *string            `json:"currency_code,omitempty" gorm:"type:varchar(3)"`                                          // Currency Code
	DiscountDescText       *string            `json:"discount_desc_text,omitempty" gorm:"type:varchar(100)" default:"SQDGENBSDTDL"`            // DiscountDescText
	SqNdcCacheServiceID    *uuid.UUID         `json:"sq_ndc_cache_service_id,omitempty"`                                                       // ServiceID
	Service                *SqNdcCacheService `json:"service,omitempty" gorm:"foreignKey:ID;references:SqNdcCacheServiceID"`                   // Service
	Pax                    *SqNdcCachePax     `json:"pax,omitempty" gorm:"foreignKey:PaxID;references:PaxRefID"`                               // Pax
	ShoppingResponseID     *string            `json:"shopping_response_id,omitempty" gorm:"type:varchar(60)" default:"SP2F-1451030339040814687"`
	// Segment                *SqNdcCachePaxSegment `json:"segment,omitempty" gorm:"foreignKey:PaxSegmentID;references:PaxSegmentRefID"`               // Segment [DEPRECATED]
}

// GetSqNdcCacheOfferService by query filter
func (data *SqNdcCacheOfferService) GetSqNdcCacheOfferServices(tx *gorm.DB, queryFilters string) (res *[]SqNdcCacheOfferService) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	err := tx.Where(qf, wf...).Preload(`Service`).Preload(`Pax`).Find(&res)
	// err := tx.Where(qf, wf...).Preload(`Service`).Preload(`Segment`).Preload(`Pax`).Find(&res) [DEPRECATED]
	if err.Error != nil {
		log.Println(err.Error)
	}
	return
}
