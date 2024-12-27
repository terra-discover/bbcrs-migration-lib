package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type Queue struct {
	basic.Base
	basic.DataOwner
	QueueAPI
}

type QueueAPI struct {
	AgentID               *uuid.UUID       `json:"agent_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"`      // The reference to the agent.
	QueueTypeID           *uuid.UUID       `json:"queue_type_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"` // The reference to the queue type.
	VisibleAt             *strfmt.DateTime `json:"visible_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"`        // The date and time that a particular task is visible to be processed.
	ExpireAt              *strfmt.DateTime `json:"expire_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"`         // The date and time that a particular task is no longer to be processed.
	Payload               *string          `json:"payload" gorm:"type:text"`                                                                    // The task to be processed.
	Acknowledgement       *string          `json:"acknowledgement,omitempty" gorm:"type:varchar(36)"`                                           // An acknowledgement identity which is assigned when a message is retrieved for processing.
	Tries                 *int             `json:"tries,omitempty"`                                                                             // The number of times the message has been retrieved from queue. Increment this value for every retrieval.
	CorporateID           *uuid.UUID       `json:"corporate_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`           // The reference to the corporate.
	ProposalID            *uuid.UUID       `json:"proposal_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`            // A reference to a proposal.
	ReservationID         *uuid.UUID       `json:"reservation_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`         // A reference to a reservation.
	TicketNumber          *string          `json:"ticket_number,omitempty" gorm:"type:varchar(36)"`                                             // A flight ticket number.
	RoomStayID            *uuid.UUID       `json:"room_stay_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`           // A reference to a room stay.
	InvoiceID             *uuid.UUID       `json:"invoice_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`             // A reference to an invoice.
	InvoiceGroupID        *uuid.UUID       `json:"invoice_group_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`       // A reference to an invoice group.
	InquiryID             *uuid.UUID       `json:"inquiry_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`             // A reference to an inquiry.
	MessageID             *uuid.UUID       `json:"message_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`             // A reference to a message.
	PersonID              *uuid.UUID       `json:"person_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`              // A reference to a person.
	EmployeeID            *uuid.UUID       `json:"employee_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`            // A reference to an employee.
	UserAccountID         *uuid.UUID       `json:"user_account_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`        // A reference to a user account.
	AttachmentID          *uuid.UUID       `json:"attachment_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`          // A reference to an attachment (multimedia description).
	DocumentID            *uuid.UUID       `json:"document_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`            // A reference to a travel document.
	CityID                *uuid.UUID       `json:"city_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`                // A reference to a city.
	EmailTemplateID       *uuid.UUID       `json:"email_template_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`      // A reference to an email template.
	EmailTemplateOfficeID *uuid.UUID       `json:"email_template_office_id,omitempty" gorm:"varchar(36)" format:"uuid" swaggertype:"string"`    // A reference to an email template office.
	Reminders             *int             `json:"reminders,omitempty"`                                                                         // The number of reminder for the same message.
	IsTest                *bool            `json:"is_test,omitempty"`                                                                           // Indicates whether this queue is for testing only.
	IsResent              *bool            `json:"is_resent,omitempty"`                                                                         // Indicates whether this queue is for resent.
}

type QueueRequest struct {
	QueueTypeCode *int    `json:"queue_type_code,omitempty"`
	QueueTypeName *string `json:"queue_type_name,omitempty"`
	QueueAPI
}

type QueueTypeMessageType struct {
	QueueTypeCode   *int    `json:"queue_type_code,omitempty"`
	MessageTypeCode *int    `json:"message_type_code,omitempty"`
	Name            *string `json:"name,omitempty"`
	IsReservation   *bool   `json:"is_reservation,omitempty"`
	IsProposal      *bool   `json:"is_proposal,omitempty"`
	IsTemplate      *bool   `json:"is_template,omitempty"`
	IsTicketNumber  *bool   `json:"is_ticket_number,omitempty"`
}
