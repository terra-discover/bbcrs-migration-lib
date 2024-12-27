package model

import (
	"log"
	"time"

	"github.com/terra-discover/bbcrs-migration-lib/lib"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

type AirItinerary struct {
	basic.Base
	basic.DataOwner
	AirItineraryAPI
	TripType                   *TripType                    `json:"trip_type,omitempty"`
	AirTraveler                []AirTraveler                `json:"air_traveler,omitempty" gorm:"foreignKey:AirItineraryID"`
	AirItineraryTotalFare      *AirItineraryTotalFare       `json:"air_itinerary_total_fare,omitempty"`
	AirOriginDestinationOption []AirOriginDestinationOption `json:"air_origin_destination_option,omitempty"`
	AirPassengerTotalFare      []AirPassengerTotalFare      `json:"air_passenger_total_fare,omitempty"`
	AirTicket                  []AirTicket                  `json:"air_ticket,omitempty"`
	AirItineraryPayment        *AirItineraryPayment         `json:"air_itinerary_payment,omitempty"`
}

type AirItineraryAPI struct {
	ItineraryStatus            *int             `json:"itinerary_status,omitempty" gorm:"type:smallint"`                                                        // Indicates the status of the itinerary, e.g. confirmed, cancelled, etc.
	TripTypeID                 *uuid.UUID       `json:"trip_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                      // Indicates type of itinerary, e.g. one way, round trip, multi-city.
	DestinationGroupID         *uuid.UUID       `json:"destination_group_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`              // The reference to the destination group.
	CurrencyID                 *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                       // The reference to the currency of the money amount.
	TotalAmountBeforeTax       *float64         `json:"total_amount_before_tax,omitempty"`                                                                      // The total amount not including any associated tax  (e.g., sales tax, VAT, GST or any associated tax).
	TotalAmountAfterTax        *float64         `json:"total_amount_after_tax,omitempty"`                                                                       // The total amount including all associated taxes  (e.g., sales tax, VAT, GST or any associated tax).
	AirlineTicketingTimeLimit  *strfmt.DateTime `json:"airline_ticketing_time_limit,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // Time limit for ticketing set by the airline, including the timezone.
	AgentTicketingTimeLimit    *strfmt.DateTime `json:"agent_ticketing_time_limit,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`   // Time limit for ticketing adjusted by the agent, including the timezone.
	TicketingTimeLimitFreeText *string          `json:"ticketing_time_limit_free_text,omitempty" gorm:"type:varchar(512)"`                                      // Free text for ticketing time limit, e.g. "TO UK BY 15JUL 1517 BOM TIME ZONE OTHERWISE WILL BE XLD".
	TicketingTimeLimitFullText *string          `json:"ticketing_time_limit_full_text,omitempty" gorm:"type:varchar(512)"`                                      // Free text for ticketing time limit, e.g. "ADTK 1B TO UK BY 15JUL 1517 BOM TIME ZONE OTHERWISE WILL BE XLD".
	CancelledAt                *strfmt.DateTime `json:"cancelled_at,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`                 //
	CancellationReason         *string          `json:"cancellation_reason,omitempty" gorm:"type:text"`                                                         //
}

func (a *AirItinerary) SetAirlineTicketingTimeLimit(db *gorm.DB, ticketingTimeLimit *strfmt.DateTime, agentID *uuid.UUID) {
	if lib.IsZeroStrfmtTimePtr(ticketingTimeLimit) {
		// do nothing if nothing to input
		return
	}

	// save to airline_ticketing_time_limit first
	a.AirlineTicketingTimeLimit = ticketingTimeLimit
	a.AgentTicketingTimeLimit = ticketingTimeLimit //default : same as airline ticketing time limit

	// get agent_ticketing_time_limit by substract from agent_setting.ticketing_time_limit_offset
	if agentID != nil {
		agentSetting := AgentSetting{}
		db.Where("agent_id = ?", agentID).Take(&agentSetting)

		if agentSetting.TicketingTimeLimitOffset != nil {
			// if agent setting offset exists : substract
			newtime := time.Time(*ticketingTimeLimit).Add(time.Duration(*agentSetting.TicketingTimeLimitOffset*-1) * time.Minute)
			a.AgentTicketingTimeLimit = lib.DateTimeptr(strfmt.DateTime(newtime))
		}
	}

	log.Printf("SET AIRLINE TICKETING TIME LIMIT : %+v (AGENT TICKETING TIME LIMIT %+v)", *ticketingTimeLimit, *a.AgentTicketingTimeLimit)
	db.Save(&a)
}
