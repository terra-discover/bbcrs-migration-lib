package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SqNdcCacheService table struct
type SqNdcCacheService struct {
	basic.Base
	basic.DataOwner
	SessionID                *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                              // SessionID
	ServiceDefinitionID      *string    `json:"service_definition_id,omitempty" gorm:"type:varchar(60)" default:"SRV1"`        // ServiceDefinitionID
	Name                     *string    `json:"name,omitempty" gorm:"type:varchar(60)" default:"ONE KILOGRAM BAGGAGE"`         // Name
	ServiceCode              *string    `json:"service_code,omitempty" gorm:"type:varchar(60)" default:"XBAG"`                 // Airline specific service code which can be associated with Reason for Issuance Sub Code. Example : XLEG for Extra Leg Room Seat, XBAG for Extra Bag
	ReasonForIssuanceCode    *string    `json:"reason_for_issueance_code,omitempty" gorm:"type:varchar(60)" default:"C"`       // A (AirTransportation), C (Baggage), E (AirportServices), F (Merchandise), G (In-flightServices)
	ReasonForIssuanceSubCode *string    `json:"reason_for_issueance_sub_code,omitempty" gorm:"type:varchar(60)" default:"0AA"` // ReasonForIssuanceSubCode
	Description              *string    `json:"description,omitempty"`                                                         // Description
	ServiceBundle            *string    `json:"service_bundle,omitempty"`                                                      // ServiceBundle
}

// GetSqNdcCacheService by query filter
func (data *SqNdcCacheService) GetSqNdcCacheService(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Take(&data)
}

type ServiceDescription struct {
	DescID          string `json:"desc_id"`
	DescText        string `json:"desc_text"`
	MarkupStyleText string `json:"markup_style_text"`
}

type ServiceBundle struct {
	ServiceDefinitionRefID string `json:"service_definition_ref_id"`
}
