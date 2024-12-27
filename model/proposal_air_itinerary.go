package model

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// ProposalAirItinerary Proposal Proposal
type ProposalAirItinerary struct {
	basic.Base
	basic.DataOwner
	ProposalAirItineraryAPI
	AirItinerary *AirItinerary `json:"air_itinerary,omitempty" gorm:"foreignKey:AirItineraryID"`
}

// ProposalAirItineraryAPI ProposalAPI Proposal API
type ProposalAirItineraryAPI struct {
	ProposalID     *uuid.UUID `json:"proposal_id" gorm:"type:varchar(36);not null;"`
	AirItineraryID *uuid.UUID `json:"air_itinerary_id" gorm:"type:varchar(36);not null"`
}

// GetProposalAirItinerary by query filter
func (data *ProposalAirItinerary) GetProposalAirItinerary(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Take(&data)
}
