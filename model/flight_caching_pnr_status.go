package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightCachingPnrStatus
type FlightCachingPnrStatus struct {
	basic.Base
	basic.DataOwner
	SessionID               *uuid.UUID       `json:"session_id,omitempty" gorm:"index:idx_flight_caching_pnr_status__pnr_id_session_id,unique;type:varchar(36);not null"`
	SourceType              *string          `json:"source_type,omitempty" gorm:"type:varchar(20)"`                                                                    // SABER, OPSIGO, NDC, ...
	PnrID                   *string          `json:"pnr_id,omitempty" gorm:"index:idx_flight_caching_pnr_status__pnr_id_session_id,unique;type:varchar(100);not null"` // PnrID from partner booking response
	BookingCode             *string          `json:"booking_code,omitempty" gorm:"type:varchar(100)"`
	AgentBookingCode        *string          `json:"agent_booking_code,omitempty" gorm:"type:varchar(100)"`
	AirlineBookingCode      *string          `json:"airline_booking_code,omitempty" gorm:"type:varchar(100)"`
	AutoIssue               *bool            `json:"auto_issue,omitempty"`          // mark that this data need to be issued automatically after success booking
	ProgressPercentage      *int             `json:"progress_percentage,omitempty"` // 0-100
	BookingDate             *strfmt.DateTime `json:"booking_date" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	TicketingTimeLimit      *strfmt.DateTime `json:"ticketing_time_limit" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`                 //date & time of ticketing limit
	AgentTicketingTimeLimit *strfmt.DateTime `json:"agent_ticketing_time_limit,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // Time limit for ticketing adjusted by the agent, including the timezone.
	ReservedDate            *strfmt.DateTime `json:"reserved_date" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	TicketedDate            *strfmt.DateTime `json:"ticketed_date" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	PartnerTicketedDate     *strfmt.DateTime `json:"partner_ticketed_date" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	CanceledDate            *strfmt.DateTime `json:"canceled_date" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	// CancellationOffer       *string          `json:"cancellation_offer" gorm:"varchar(20)"`    // no_cancellation, void_booking, void_ticket, refund_ticket
	Status       *string `json:"status,omitempty" gorm:"type:varchar(20)"` // pending, booked,
	InvoiceNo    *string `json:"invoice_no,omitempty" gorm:"type:varchar(100)"`
	ErrorMessage *string `json:"error_message,omitempty"`
}
