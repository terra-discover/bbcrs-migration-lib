package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// MessageType Message Type
type MessageType struct {
	basic.Base
	basic.DataOwner
	MessageTypeAPI
	MessageTypeTranslation *MessageTypeTranslation `json:"message_type_translation,omitempty"` // Message Type Translation
	MessageCategory        *MessageCategory        `json:"message_category,omitempty"`         // Message Category
}

// MessageTypeAPI Message Type API
type MessageTypeAPI struct {
	MessageTypeCode   *int       `json:"message_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null" example:"1"`            // Message Type Code
	MessageTypeName   *string    `json:"message_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Feedback"` // Message Type Name
	IsSystem          *bool      `json:"is_system,omitempty" example:"true"`                                                                                      // Is System
	IsTransactional   *bool      `json:"is_transactional,omitempty" example:"true"`                                                                               // Is Transactional
	IsCampaign        *bool      `json:"is_campaign,omitempty" example:"true"`                                                                                    // Is Campaign
	IsInvoice         *bool      `json:"is_invoice,omitempty" example:"true"`                                                                                     // Is Invoice
	MessageCategoryID *uuid.UUID `json:"message_category_id,omitempty" gorm:"type:varchar(36)"`                                                                   // Message Type ID
	Priority          *int       `json:"priority,omitempty" example:"1" description:"smaller numbers indicate higher priority"`                                   // Priority
}

// Seed MessageType
func (messageType *MessageType) Seed() *MessageType {
	seed := MessageType{
		MessageTypeAPI: MessageTypeAPI{
			MessageTypeCode:   lib.Intptr(1),
			MessageTypeName:   lib.Strptr("Message Type Name"),
			IsSystem:          lib.Boolptr(false),
			IsTransactional:   lib.Boolptr(false),
			IsCampaign:        lib.Boolptr(false),
			IsInvoice:         lib.Boolptr(false),
			MessageCategoryID: nil,
			Priority:          lib.Intptr(0),
		},
	}
	_ = lib.Merge(seed, &messageType)
	return messageType
}

// MessageTypes migrator
type MessageTypes []MessageType

// Seed data
func (messageTypes *MessageTypes) Seed() *MessageTypes {
	nameString := "General inquiry \nFeedback \nRequest for Over Credit Authorization \nPromotion \nNewsletter \nSurvey \nSystem \nTransactional \nRequest for Proposal Review \nReminder for Request for Proposal Review \nInvoice \nRequest for Travel Request Submission \nReminder for Travel Request Submission \nRequest for Travel Request Revision \nRequest for Credit Card Online Payment \nReminder for Credit Card Online Payment \nApproved to Issue \nReminder for Approved to Issue \nTicketing Time Limit \nReminder for Ticketing Time Limit \nBooking Confirmation \nFlight Ticket \nHotel Voucher \nRequest for Payment Settlement \nReminder for Payment Settlement \nThank You and Survey \nDuring Stay Email \nFlight Booking Confirmation \nHotel Booking Confirmation \nFlight Issuance Confirmation \nInternal Flight Booking Confirmation \nInternal Hotel Booking Confirmation \nInternal Flight Issuance Confirmation \nCancellation Confirmation \nFlight Cancellation Confirmation \nHotel Cancellation Confirmation \nInternal Flight Cancellation Confirmation \nInternal Hotel Cancellation Confirmation \nInternal Ticketing Time Limit \nReminder for Internal Ticketing Time Limit \nSurvey Response \nInternal User Activation \nInternal User Welcome \nCorporate User Activation \nCorporate User Welcome \nMember Activation \nMember Welcome \nPassword Change \nPassword Restore \nOne Time Password \nInvoice Per Transaction \nInvoice Batch \nRequest for Over Credit Authorization Approved \nRequest for Over Credit Authorization Rejected \nCancellation Deadline \nReminder for Cancellation Deadline \nInternal Cancellation Deadline \nReminder for Internal Cancellation Deadline \nTicket Not Found \nReminder for Invoice Per Transaction \nReminder for Invoice Batch \nReminder for Request for Over Credit Authorization \nTransactional - Flight \nTransactional - Hotel \nTransactional - Tour \nReminder for Travel Request Revision \nSend to Mail \nSchedule Change - Change Flight Number \nSchedule Change - Flight Retimed \nTravel Request Submission \nProposal Option to Book \nReminder for Proposal Option to Book \nInvoice Per Transaction - Flight \nInvoice Per Transaction - Hotel \nInvoice Per Transaction - Other \nInvoice Batch - Flight \nInvoice Batch - Hotel \nInvoice Batch - Other \nReminder for Invoice Per Transaction - Flight \nReminder for Invoice Per Transaction - Hotel \nReminder for Invoice Per Transaction - Other \nReminder for Invoice Batch - Flight \nReminder for Invoice Batch - Hotel \nReminder for Invoice Batch - Other \nPassport Expiration \nReminder for Passport Expiration \nCorporate Contract Expiration \nReminder for Corporate Contract Expiration \nRequest for Hotel Booking Acknowledgment \nReminder for Hotel Booking Acknowledgment \nHotel Booking Acknowledged \nInternal Hotel Booking Acknowledged \nCorporate Data Verification \nRequest to Reissue Ticket \nReminder for Request to Reissue Ticket \nRequest to Revalidate Ticket \nReminder for Request to Revalidate Ticket \nRequest to Reissue Ticket (Name Change) \nReminder for Request to Reissue Ticket (Name Change) \nRequest to Modify Hotel Voucher \nReminder for Request to Modify Hotel Voucher \nRequest to Reissue Ticket Completed \nRequest to Revalidate Ticket Completed \nRequest to Reissue Ticket (Name Change) Completed \nRequest to Modify Hotel Voucher Completed \nInvoice Per Transaction - Tour \nInvoice Per Transaction - Document \nInvoice Batch - Tour \nInvoice Batch - Document \nReminder for Invoice Per Transaction - Tour \nReminder for Invoice Per Transaction - Document \nReminder for Invoice Batch - Tour \nReminder for Invoice Batch - Document \nCredit Limit Almost Used Up \nFlight Itinerary \ne-Ticket \ne-Itinerary \ne-Voucher \ne-Invoice Per Transaction \ne-Invoice Batch"
	name := strings.Split(nameString, " \n")
	for i := 0; i < len(name); i++ {
		*messageTypes = append(*messageTypes, MessageType{
			MessageTypeAPI: MessageTypeAPI{
				MessageTypeCode:   lib.Intptr(i + 1),
				MessageTypeName:   lib.Strptr(name[i]),
				IsSystem:          lib.Boolptr(false),
				IsTransactional:   lib.Boolptr(false),
				IsCampaign:        lib.Boolptr(false),
				IsInvoice:         lib.Boolptr(false),
				MessageCategoryID: nil,
				Priority:          lib.Intptr(0),
			},
		})
	}
	return messageTypes
}
