package model

import (
	"time"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

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

// Reservation Reservation
type Reservation struct {
	basic.Base
	basic.DataOwner
	ReservationAPI
	ReservationAirItinerary        *ReservationAirItinerary `json:"reservation_air_itinerary,omitempty" gorm:"foreignKey:ID;references:ReservationID"`
	ReservationProfile             []ReservationProfile     `json:"reservation_profile,omitempty"`
	TransactionStatus              *Status                  `json:"transaction_status,omitempty" gorm:"foreignKey:ReservationStatus;references:StatusCode"`
	TransactionCreditStatus        *Status                  `json:"transaction_credit_status,omitempty" gorm:"foreignKey:CreditStatus;references:StatusCode"`
	TransactionTravelRequestStatus *Status                  `json:"transaction_travel_request_status,omitempty" gorm:"foreignKey:TravelRequestStatus;references:StatusCode"`
	Agent                          *Agent                   `json:"agent,omitempty"`
	Corporate                      *Corporate               `json:"corporate,omitempty"`
	Proposal                       *Proposal                `json:"proposal,omitempty"`
	ProductType                    *ProductType             `json:"product_type,omitempty"`
	TravelPurpose                  *TravelPurpose           `json:"travel_purpose,omitempty"`
}

// ReservationAPI Reservation API
type ReservationAPI struct {
	ReservationCode             *string          `json:"reservation_code,omitempty" gorm:"type:varchar(36);not null"`                                            // The code of the reservation.
	AgentID                     *uuid.UUID       `json:"agent_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                          // The reference to the agent.
	CorporateID                 *uuid.UUID       `json:"corporate_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                      // The reference to the corporate.
	ReservationStatus           *int             `json:"reservation_status,omitempty" gorm:"type:smallint"`                                                      //  Indicates the status of the reservation, e.g. confirmed, cancelled, etc.
	ProposalID                  *uuid.UUID       `json:"proposal_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                       // The reference to the proposal.
	ProductTypeID               *uuid.UUID       `json:"product_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                   //  The reference to the type of product, e.g. flight, hotel, tour, etc.
	CurrencyID                  *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                       // The reference to the currency of the money amount.
	TotalAmountBeforeTax        *float64         `json:"total_amount_before_tax,omitempty"`                                                                      // The total amount not including any associated tax  (e.g., sales tax, VAT, GST or any associated tax).
	TotalAmountAfterTax         *float64         `json:"total_amount_after_tax,omitempty"`                                                                       // The total amount including all associated taxes  (e.g., sales tax, VAT, GST or any associated tax).
	LanguageID                  *uuid.UUID       `json:"language_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                       // The reference to the language.
	SpecialRequest              *string          `json:"special_request,omitempty" gorm:"type:text"`                                                             // Special request for the whole reservation.
	TravelPurposeID             *uuid.UUID       `json:"travel_purpose_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                 // The reference to the travel purpose.
	IsProject                   *bool            `json:"is_project,omitempty"`                                                                                   // Indicates whether this reservation is for a project.
	IsSelfBookingTool           *bool            `json:"is_self_booking_tool,omitempty"`                                                                         // Indicates whether this reservation comes from corporate self booking tool.
	IsPersonalTravel            *bool            `json:"is_personal_travel,omitempty"`                                                                           // Indicates whether this reservation is for corporate personal travel.
	IsBundling                  *bool            `json:"is_bundling,omitempty"`                                                                                  // Indicates whether this reservation is for booking flight+hotel.
	CreditStatus                *int             `json:"credit_status,omitempty" gorm:"type:smallint"`                                                           // Indicates the status of the credit, e.g. OK, Not OK, being reviewed, etc.
	CreditComment               *string          `json:"credit_comment,omitempty" gorm:"type:text"`                                                              // The comment related to over credit authorization of the reservation.
	TravelRequestStatus         *int             `json:"travel_request_status,omitempty" gorm:"type:smallint"`                                                   // Indicates the status of the travel request, e.g. required, verified, etc.
	TravelRequestComment        *string          `json:"travel_request_comment,omitempty" gorm:"type:text"`                                                      // The comment related to the travel request submission.
	ModificationStatus          *int             `json:"modification_status,omitempty" gorm:"type:smallint"`                                                     // Indicates the status of the modification, e.g. pending modification, modification ready, modification accepted, modification rejected, etc.
	ModificationComment         *string          `json:"modification_comment,omitempty" gorm:"type:text"`                                                        // The comment related to modification of the reservation.
	CancellationStatus          *int             `json:"cancellation_status,omitempty" gorm:"type:smallint"`                                                     // Indicates the status of the cancellation, e.g. pending cancellation, cancellation ready, cancellation accepted, cancellation rejected, etc.
	CancellationComment         *string          `json:"cancellation_comment,omitempty" gorm:"type:text"`                                                        // The comment related to cancellation of the reservation.
	Comment                     *string          `json:"comment,omitempty" gorm:"type:text"`                                                                     // The internal comment of the reservation.
	ConsumerViewableComment     *string          `json:"consumer_viewable_comment,omitempty" gorm:"type:text"`                                                   // The comment that is shown to the consumer.
	TotalPenaltyAmountBeforeTax *float64         `json:"total_penalty_amount_before_tax,omitempty"`                                                              // The total penalty amount not including any associated tax  (e.g., sales tax, VAT, GST or any associated tax).
	TotalPenaltyAmountAfterTax  *float64         `json:"total_penalty_amount_after_tax,omitempty"`                                                               // The total penalty amount including all associated taxes  (e.g., sales tax, VAT, GST or any associated tax).
	CancellationReason          *string          `json:"cancellation_reason,omitempty" gorm:"type:text"`                                                         // The reason for cancellation.
	ReservationTimestamp        *strfmt.DateTime `json:"reservation_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`        // Timestamp of the creation of the reservation (in UTC).
	ModificationTimestamp       *strfmt.DateTime `json:"modification_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`       // Timestamp of the last modification (in UTC).
	CancellationTimestamp       *strfmt.DateTime `json:"cancellation_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`       // Timestamp of the cancellation (in UTC).
	ReservationLocalTimestamp   *strfmt.DateTime `json:"reservation_local_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`  // Timestamp of the creation of the reservation (in Agent's Time Zone).
	ModificationLocalTimestamp  *strfmt.DateTime `json:"modification_local_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // Timestamp of the last modification (in Agent's Time Zone).
	CancellationLocalTimestamp  *strfmt.DateTime `json:"cancellation_local_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // Timestamp of the cancellation (in Agent's Time Zone).
}

// GetReservation by query filter
func (data *Reservation) GetReservation(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).First(&data)
}

// SetReservation
func (data *Reservation) SetReservation(tx *gorm.DB) {
	data.ReservationTimestamp = lib.DateTimeptr(strfmt.DateTime(time.Now()))
	data.ReservationLocalTimestamp = lib.DateTimeptr(strfmt.DateTime(time.Now()))
	tx.Create(data)
}

// SetHistory will save current reservation to its history, based on current userIDs (only 1st index used)
func (data *Reservation) SetHistory(tx *gorm.DB, userIDs ...*uuid.UUID) {
	var userID *uuid.UUID
	if len(userIDs) > 0 {
		userID = userIDs[0]
	}

	// create a history record by current latest reservation state

	var description *string // TODO : get description from Event Tracking result
	rh := ReservationHistory{
		ReservationHistoryAPI: ReservationHistoryAPI{
			ReservationID:        data.ID,
			ReservationCode:      data.ReservationCode,
			ReservationStatus:    data.ReservationStatus,
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
			Timestamp:            data.ReservationTimestamp,
			LocalTimestamp:       data.ReservationLocalTimestamp,
		},
	}
	tx.Create(&rh)
}
