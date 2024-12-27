package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
Status Flag:
	1 = Active : The reservation is currently active.
	2 = Draft  : This status is used to signify that the reservation is in a draft state,
				 typically employed for reservations undergoing modification in progress.
	3 = Unused : This status has not been utilized so far (as of the current implementation).
	4 = Deleted: The reservation has been marked as deleted.
*/

// Proposal Proposal
type Proposal struct {
	basic.Base
	basic.DataOwner
	ProposalAPI
	ProposalAirItinerary           *ProposalAirItinerary          `json:"proposal_air_itinerary,omitempty" gorm:"foreignKey:ID;references:ProposalID"`
	AirTravelerCriteria            *AirTravelerCriteria           `json:"air_traveler_criteria,omitempty"`
	AirOriginDestinationCriteria   []AirOriginDestinationCriteria `json:"air_origin_destintaiton_criteria,omitempty"`
	TransactionStatus              *Status                        `json:"transaction_status,omitempty" gorm:"foreignKey:ProposalStatus;references:StatusCode"`
	TransactionCreditStatus        *Status                        `json:"transaction_credit_status,omitempty" gorm:"foreignKey:CreditStatus;references:StatusCode"`
	TransactionTravelRequestStatus *Status                        `json:"transaction_travel_request_status,omitempty" gorm:"foreignKey:TravelRequestStatus;references:StatusCode"`
	Agent                          *Agent                         `json:"agent,omitempty"`
	Corporate                      *Corporate                     `json:"corporate,omitempty"`
	ProductType                    *ProductType                   `json:"product_type,omitempty"`
	ProposalProfile                *ProposalProfile               `json:"proposal_profile,omitempty" gorm:"foreignKey:ID;references:ProposalID"`
	TravelPurpose                  *TravelPurpose                 `json:"travel_purpose,omitempty"`
	ProposalBudget                 *ProposalBudget                `json:"proposal_budget,omitempty"`
}

// ProposalAPI Proposal API
type ProposalAPI struct {
	ProposalCode               *string          `json:"proposal_code,omitempty" gorm:"type:varchar(36);not null"`                                               // The code of the proposal reffer from sessionID.
	AgentID                    *uuid.UUID       `json:"agent_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                          // The reference to the agent.
	CorporateID                *uuid.UUID       `json:"corporate_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                      // The reference to the corporate.
	ProposalStatus             *int             `json:"proposal_status,omitempty" gorm:"type:smallint"`                                                         // Indicates the status of the proposal. 0=pending, 1=booked 2=issued 3=canceled ...
	ProductTypeID              *uuid.UUID       `json:"product_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                   // The reference to the type of product, e.g. flight, hotel, tour, etc.
	CurrencyID                 *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                       // The reference to the currency of the money amount.
	TotalAmountBeforeTax       *float64         `json:"total_amount_before_tax,omitempty"`                                                                      // The total amount not including any associated tax  (e.g., sales tax, VAT, GST or any associated tax).
	TotalAmountAfterTax        *float64         `json:"total_amount_after_tax,omitempty"`                                                                       // The total amount including all associated taxes  (e.g., sales tax, VAT, GST or any associated tax).
	LanguageID                 *uuid.UUID       `json:"language_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                       // The reference to the language.
	SpecialRequest             *string          `json:"special_request,omitempty" gorm:"type:text"`                                                             // Special request for the whole proposal.
	TravelPurposeID            *uuid.UUID       `json:"travel_purpose_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                 // The reference to the travel purpose.
	IsProject                  *bool            `json:"is_project,omitempty"`                                                                                   // Indicates whether this proposal is for a project.
	IsSelfBookingTool          *bool            `json:"is_self_booking_tool,omitempty"`                                                                         // Indicates whether this proposal comes from corporate self booking tool.
	IsPersonalTravel           *bool            `json:"is_personal_travel,omitempty"`                                                                           // Indicates whether this proposal is for corporate personal travel.
	IsApprovedToIssue          *bool            `json:"is_approved_to_issue,omitempty"`                                                                         // Indicates whether this proposal is for corporate personal travel.
	Opening                    *string          `json:"opening,omitempty" gorm:"type:text"`                                                                     // The comment related to over credit authorization of the proposal.
	Closing                    *string          `json:"closing,omitempty" gorm:"type:text"`                                                                     // The comment related to over credit authorization of the proposal.
	CreditStatus               *int             `json:"credit_status,omitempty" gorm:"type:smallint"`                                                           // Indicates the status of the credit, e.g. OK, Not OK, being reviewed, etc.
	CreditComment              *string          `json:"credit_comment,omitempty" gorm:"type:text"`                                                              // The comment related to over credit authorization of the proposal.
	TravelRequestStatus        *int             `json:"travel_request_status,omitempty" gorm:"type:smallint"`                                                   // Indicates the status of the travel request, e.g. required, verified, etc.
	TravelRequestComment       *string          `json:"travel_request_comment,omitempty" gorm:"type:text"`                                                      // The comment related to the travel request submission.
	ModificationStatus         *int             `json:"modification_status,omitempty" gorm:"type:smallint"`                                                     // Indicates the status of the modification, e.g. pending modification, modification ready, modification accepted, modification rejected, etc.
	ModificationComment        *string          `json:"modification_comment,omitempty" gorm:"type:text"`                                                        // The comment related to modification of the proposal.
	CancellationStatus         *int             `json:"cancellation_status,omitempty" gorm:"type:smallint"`                                                     // Indicates the status of the cancellation, e.g. pending cancellation, cancellation ready, cancellation accepted, cancellation rejected, etc.
	CancellationComment        *string          `json:"cancellation_comment,omitempty" gorm:"type:text"`                                                        // The comment related to cancellation of the proposal.
	Comment                    *string          `json:"comment,omitempty" gorm:"type:text"`                                                                     // The internal comment of the proposal.
	ConsumerViewableComment    *string          `json:"consumer_viewable_comment,omitempty" gorm:"type:text"`                                                   // The comment that is shown to the consumer.
	ProposalTimestamp          *strfmt.DateTime `json:"proposal_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`           // Timestamp of the creation of the proposal (in UTC).
	ModificationTimestamp      *strfmt.DateTime `json:"modification_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`       // Timestamp of the last modification (in UTC).
	CancellationTimestamp      *strfmt.DateTime `json:"cancellation_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`       //
	ProposalLocalTimestamp     *strfmt.DateTime `json:"proposal_local_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`     //
	ModificationLocalTimestamp *strfmt.DateTime `json:"modification_local_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` //
	CancellationLocalTimestamp *strfmt.DateTime `json:"cancellation_local_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` //
	CancellationCharge         *float64         `json:"cancellation_charge,omitempty"`
	CancellationReason         *string          `json:"cancellation_reason,omitempty" gorm:"type:text"` // The reason for cancellation.
}

// GetProposal by query filter
func (data *Proposal) GetProposal(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Preload(`ProposalAirItinerary.AirItinerary`).Where(qf, wf...).First(&data)
}

// SetProposal
func (data *Proposal) SetProposal(tx *gorm.DB) {
	data.ProposalTimestamp = lib.DateTimeptr(strfmt.DateTime(time.Now()))
	data.ProposalLocalTimestamp = lib.DateTimeptr(strfmt.DateTime(time.Now()))
	tx.Create(data)
}

// SetHistory will save current proposal to its history, based on current userIDs (only 1st index used)
func (data *Proposal) SetHistory(tx *gorm.DB, userIDs ...*uuid.UUID) {
	var userID *uuid.UUID
	if len(userIDs) > 0 {
		userID = userIDs[0]
	}

	// create a history record by current latest proposal state

	var description *string // TODO : get description from Event Tracking result

	ph := ProposalHistory{
		ProposalHistoryAPI: ProposalHistoryAPI{
			ProposalID:           data.ID,
			ProposalCode:         data.ProposalCode,
			ProposalStatus:       data.ProposalStatus,
			CreditStatus:         data.CreditStatus,
			CreditComment:        data.CreditComment,
			TravelRequestStatus:  data.TravelRequestStatus,
			TravelRequestComment: data.TravelRequestComment,
			ModificationStatus:   data.ModificationStatus,
			ModificationComment:  data.ModificationComment,
			CancellationStatus:   data.CancellationStatus,
			CancellationComment:  data.CancellationComment,
			UserAccountID:        userID,
			Description:          description,
			Timestamp:            data.ProposalTimestamp,
			LocalTimestamp:       data.ProposalLocalTimestamp,
		},
	}
	tx.Create(&ph)
}
