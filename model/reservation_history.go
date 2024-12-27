package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReservationHistory Reservation History
type ReservationHistory struct {
	basic.Base
	basic.DataOwner
	ReservationHistoryAPI
}

// ReservationHistoryAPI Reservation History API
type ReservationHistoryAPI struct {
	ReservationID        *uuid.UUID       `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`  // The reference to the reservation.
	ReservationCode      *string          `json:"reservation_code,omitempty" gorm:"type:varchar(36)"`                                            // The code of the reservation.
	ReservationStatus    *int             `json:"reservation_status,omitempty" gorm:"type:smallint"`                                             // Indicates the status of the reservation, e.g. confirmed, cancelled, etc.
	UserAccountID        *uuid.UUID       `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // The reference to the active user account.
	CreditStatus         *int             `json:"credit_status,omitempty" gorm:"type:smallint"`                                                  // Indicates the status of the credit, e.g. OK, Not OK, being reviewed, etc.
	CreditComment        *string          `json:"credit_comment,omitempty" gorm:"type:text"`                                                     // The comment related to over credit authorization of the reservation.
	TravelRequestStatus  *int             `json:"travel_request_status,omitempty" gorm:"type:smallint"`                                          // Indicates the status of the travel request, e.g. required, verified, etc.
	TravelRequestComment *string          `json:"travel_request_comment,omitempty" gorm:"type:text"`                                             // The comment related to the travel request submission.
	ModificationStatus   *int             `json:"modification_status,omitempty" gorm:"type:smallint"`                                            // Indicates the status of the modification, e.g. pending modification, modification ready, modification accepted, modification rejected, etc.
	ModificationComment  *string          `json:"modification_comment,omitempty" gorm:"type:text"`                                               // The comment related to modification of the reservation.
	CancellationStatus   *int             `json:"cancellation_status,omitempty" gorm:"type:smallint"`                                            // Indicates the status of the cancellation, e.g. pending cancellation, cancellation ready, cancellation accepted, cancellation rejected, etc.
	CancellationComment  *string          `json:"cancellation_comment,omitempty" gorm:"type:text"`                                               // The comment related to cancellation of the reservation.
	Description          *string          `json:"description,omitempty" gorm:"type:text"`                                                        // The description of the activity.
	Timestamp            *strfmt.DateTime `json:"timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`           // Timestamp of the activity (in UTC).
	LocalTimestamp       *strfmt.DateTime `json:"local_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`     // Timestamp of the activity (in Agent's Time Zone).
}
