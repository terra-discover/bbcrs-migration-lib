package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

type FlightCachingReservation struct {
	basic.Base
	basic.DataOwner
	SessionID               *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_reservation_session_id"`
	ReservationID           *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_reservation_reservation_id"`
	ReservationCode         *string    `json:"reservation_code,omitempty" gorm:"type:varchar(100)"`
	IsIssueTicketInProgress *bool      `json:"is_issue_ticket_in_progress,omitempty"`
	IssueTicketFailed       *int       `json:"issue_ticket_failed,omitempty"`
	TicketFileID            *string    `json:"ticket_file_id,omitempty" gorm:"type:varchar(36)"` // multimedia description id
	TicketFileURL           *string    `json:"ticket_url,omitempty"`                             // multimedia description url
	IsPrintTicketInProgress *bool      `json:"is_print_ticket_in_progress,omitempty"`
	PrintTicketFailed       *int       `json:"print_ticket_failed,omitempty"`
	IsTicketSent            *bool      `json:"is_ticket_sent,omitempty"`
}

// AfterCreate Data
func (f *FlightCachingReservation) AfterCreate(tx *gorm.DB) error {
	return afterCreateFlightCachingReservation(tx, *f)
}
