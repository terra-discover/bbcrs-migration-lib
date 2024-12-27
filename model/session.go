package model

import (
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Session Model
type Session struct {
	basic.Base
	basic.DataOwner
	SessionKey    *string          `json:"session_key,omitempty" gorm:"type:varchar(256);not null"`                              // Session Key
	ExpireAt      *strfmt.DateTime `json:"expire_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`  // Expire At
	UserAccountID *uuid.UUID       `json:"user_account_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // User Account Id
	IPAddress     *string          `json:"ip_address,omitempty" gorm:"type:varchar(256)"`                                        // Ip Address
	UserAgent     *string          `json:"user_agent,omitempty" gorm:"type:text"`                                                // User Agent
	DeviceID      *uuid.UUID       `json:"device_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`       // Device Id
	Payload       *string          `json:"payload,omitempty" gorm:"type:text"`                                                   // Payload
	ProposalID    *uuid.UUID       `json:"proposal_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`     // Proposal Id
	ReservationID *uuid.UUID       `json:"reservation_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`  // Reservation Id
	MessageID     *uuid.UUID       `json:"message_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`      // Message Id
	Device        *Device          `json:"device,omitempty" swaggerignore:"true"`                                                // Device
	UserAccount   *UserAccount     `json:"user_account,omitempty" swaggerignore:"true"`                                          // User Account
}

// afterCreateFlightCachingSession Data -
func afterCreateFlightCachingSession(db *gorm.DB, f FlightCachingSession) error {
	// Logging for make hooks more traceable
	setLog := func(message string) {
		log.Printf("hooks afterCreateFlightCachingSession session id: %s , message: %s", f.SessionID, message)
	}

	setLog("START")
	defer setLog("END")

	if lib.IsEmptyUUIDPtr(f.SessionID) {
		setLog("empty session id")
		return nil
	}

	// Don't return error when flight_caching_request not found
	var flightCachingRequest FlightCachingRequest
	if err := db.Where("session_id = ?", f.SessionID).
		Order(`created_at DESC`).
		First(&flightCachingRequest).Error; err != nil {
		setLog("flight_caching_request not found")
		return nil
	}

	// Don't return error when session not found
	var session Session
	if err := db.Where("user_account_id = ?", flightCachingRequest.UserID).
		Order(`created_at DESC`).
		First(&session).Error; err != nil {
		setLog("session not found")
		return nil
	}

	// Set session id
	session.ID = f.SessionID

	// Set session key
	session.SessionKey = lib.Strptr(lib.GenUUIDString())

	if len(f.SessionPayload) > 0 {
		// Set session payload
		// Don't return error when flight_caching_payload not found
		payloadValue, err := f.SessionPayload.Value()
		if err != nil {
			setLog("session payload invalid value: " + err.Error())
			return nil
		}
		payload := string(payloadValue.([]byte))
		session.Payload = lib.Strptr(payload)
	} else {
		session.Payload = nil
	}

	// Don't return error when failed create session
	err := db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&session).Error
	if err != nil {
		setLog("failed create session")
		return nil
	}

	return nil
}

// afterSaveFlightCachingSession Data -
func afterSaveFlightCachingSession(db *gorm.DB, f FlightCachingSession) error {
	// Logging for make hooks more traceable
	setLog := func(message string) {
		log.Printf("hooks afterSaveFlightCachingSession session id: %s , message: %s", f.SessionID, message)
	}

	// setLog("START")
	// defer setLog("END")

	if lib.IsEmptyUUIDPtr(f.SessionID) || len(f.SessionPayload) == 0 {
		setLog("empty session id or session payload")
		return nil
	}

	// Don't return error when session not found
	var session Session
	if err := db.Where("id = ?", f.SessionID).
		Order(`created_at DESC`).
		First(&session).Error; err != nil {
		setLog("session not found")
		return nil
	}

	// Set session payload
	// Don't return error when flight_caching_payload not found
	payloadValue, err := f.SessionPayload.Value()
	if err != nil {
		setLog("session payload invalid value: " + err.Error())
		return nil
	}
	payload := string(payloadValue.([]byte))
	session.Payload = lib.Strptr(payload)

	// Only update if session_key still empty
	if lib.IsEmptyStrPtr(session.SessionKey) {
		// Set session key
		session.SessionKey = lib.Strptr(lib.GenUUIDString())
	}

	// Don't return error when failed update session
	err = db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Updates(&session).Error
	if err != nil {
		setLog("failed update session")
		return nil
	}

	return nil
}

// afterCreateFlightCachingProposal Data
func afterCreateFlightCachingProposal(db *gorm.DB, f FlightCachingProposal) error {
	// Logging for make hooks more traceable
	setLog := func(message string) {
		log.Printf("hooks afterCreateFlightCachingProposal session id: %s , message: %s", f.SessionID, message)
	}

	// setLog("START")
	// defer setLog("END")

	if lib.IsEmptyUUIDPtr(f.SessionID) || lib.IsEmptyUUIDPtr(f.ProposalID) {
		setLog("empty session id or proposal id")
		return nil
	}

	// Don't return error when session not found
	var session Session
	if err := db.Where("id = ?", f.SessionID).
		Order(`created_at DESC`).
		First(&session).Error; err != nil {
		setLog("session not found")
		return nil
	}

	// Set session proposal_id
	session.ProposalID = f.ProposalID

	// Don't return error when failed update session
	err := db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Updates(&session).Error
	if err != nil {
		setLog("failed update session")
		return nil
	}

	return nil
}

// afterCreateFlightCachingReservation Data
func afterCreateFlightCachingReservation(db *gorm.DB, f FlightCachingReservation) error {
	// Logging for make hooks more traceable
	setLog := func(message string) {
		log.Printf("hooks afterCreateFlightCachingReservation session id: %s , message: %s", f.SessionID, message)
	}

	setLog("START")
	defer setLog("END")

	if lib.IsEmptyUUIDPtr(f.SessionID) || lib.IsEmptyUUIDPtr(f.ReservationID) {
		setLog("empty session id or reservation id")
		return nil
	}

	// Don't return error when session not found
	var session Session
	if err := db.Where("id = ?", f.SessionID).
		Order(`created_at DESC`).
		First(&session).Error; err != nil {
		setLog("session not found")
		return nil
	}

	// Set session reservation_id
	session.ReservationID = f.ReservationID

	// Don't return error when failed update session
	err := db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Updates(&session).Error
	if err != nil {
		setLog("failed update session")
		return nil
	}

	return nil
}
