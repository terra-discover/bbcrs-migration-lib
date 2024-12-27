package model

import (
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// SqNdcCacheOffersFareComponent one to many relation
type SqNdcCacheOffersFareComponent struct {
	basic.Base
	basic.DataOwner
	SessionID *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	SqNdcCacheOffersFareComponentAPI
}

// SqNdcCacheOffersFareComponentAPI
type SqNdcCacheOffersFareComponentAPI struct {
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

// // Seed SqNdcCacheOffersFareComponent
// func (d *SqNdcCacheOffersFareComponent) Seed() *SqNdcCacheOffersFareComponent {
// 	sessionID, _ := uuid.Parse(viper.GetString("SESSION_ID"))
// 	fareDetailID, _ := uuid.Parse("e3b58926-b85a-4b55-9a38-d5f50ea98c09")
// 	seed := SqNdcCacheOffersFareComponent{
// 		SessionID: &sessionID,
// 		SqNdcCacheOffersFareComponentAPI: SqNdcCacheOffersFareComponentAPI{
// 			FareDetailID:             &fareDetailID,
// 			PriceClassRef:            lib.Strptr("FF51SIN"),
// 			SegmentRefs:              lib.Strptr("SEG2"),
// 			FareBasisCode:            lib.Strptr("V16SGO"),
// 			FareRulesRemarksCategory: lib.Strptr("FT"),
// 			FareRulesRemarksText:     lib.Strptr("NEGOTIATED FARE"),
// 			Rbd:                      lib.Strptr("V"),
// 			CabinTypeCode:            lib.Strptr("M"),
// 			CabinTypeName:            lib.Strptr("ECO"),
// 		},
// 	}
// 	_ = lib.Merge(seed, &d)
// 	return d
// }
