package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentMessage Agent Message
type AgentMessage struct {
	basic.Base
	basic.DataOwner
	AgentMessageAPI
	Agent       *Agent       `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	Message     *Message     `json:"message" gorm:"foreignKey:MessageID;references:ID"`
	Corporate   *Corporate   `json:"corporate" gorm:"foreignKey:CorporateID;references:ID"`
	Proposal    *Proposal    `json:"proposal" gorm:"foreignKey:ProposalID;references:ID"`
	Reservation *Reservation `json:"reservation" gorm:"foreignKey:ReservationID;references:ID"`
	Form        *Form        `json:"form" gorm:"foreignKey:FormID;references:ID"`
	Module      *Module      `json:"module" gorm:"foreignKey:ModuleID;references:ID"`
	Term        *Term        `json:"term" gorm:"foreignKey:TermID;references:ID"`
	Pages       *Page        `json:"pages" gorm:"foreignKey:PageID;references:ID"`
	Approver    *UserAccount `json:"approver" gorm:"foreignKey:ApproverID;references:ID"`
}

// AgentMessageAPI Agent Message Api
type AgentMessageAPI struct {
	AgentID                         *uuid.UUID       `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	MessageID                       *uuid.UUID       `json:"message_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ReferenceCode                   *string          `json:"reference_code,omitempty" gorm:"type:varchar(256)"`
	EmailTemplateID                 *uuid.UUID       `json:"email_template_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	CorporateID                     *uuid.UUID       `json:"corporate_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	BusinessPartnerID               *uuid.UUID       `json:"business_partner_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ProposalID                      *uuid.UUID       `json:"proposal_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ReservationID                   *uuid.UUID       `json:"reservation_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	TicketNumber                    *uuid.UUID       `json:"ticket_number,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	InvoiceID                       *uuid.UUID       `json:"invoice_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	InvoiceGroupID                  *uuid.UUID       `json:"invoice_group_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	InquiryID                       *uuid.UUID       `json:"inquiry_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	FormID                          *uuid.UUID       `json:"form_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ModuleID                        *uuid.UUID       `json:"module_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	TermID                          *uuid.UUID       `json:"term_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	PageID                          *uuid.UUID       `json:"page_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	NewsID                          *uuid.UUID       `json:"news_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ContactPersonID                 *uuid.UUID       `json:"contact_person_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ContactEmployeeID               *uuid.UUID       `json:"contact_employee_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ContactUserAccountID            *uuid.UUID       `json:"contact_user_account_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ContactName                     *string          `json:"contact_name,omitempty" gorm:"type:varchar(128)"`
	ContactPhoneNumber              *string          `json:"contact_phone_number,omitempty" gorm:"type:varchar(32)"`
	ContactEmail                    *string          `json:"contact_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`
	Topic                           *string          `json:"topic,omitempty" gorm:"type:varchar(256)"`
	Message                         *string          `json:"message,omitempty" gorm:"type:text"`
	CreditStatus                    *int             `json:"credit_status,omitempty" gorm:"type:smallint"`
	Reason                          *string          `json:"reason,omitempty" gorm:"type:text"`
	CreditApprovalTimeLimit         *strfmt.DateTime `json:"credit_approval_time_limit,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	CorporateCreditLimitToleranceID *uuid.UUID       `json:"corporate_credit_limit_tolerance_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ApproverID                      *uuid.UUID       `json:"approver_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ParentAgentMessageID            *uuid.UUID       `json:"parent_agent_message_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	EmailSignInCode                 *string          `json:"email_sign_in_code,omitempty" gorm:"type:varchar(256)"`
	EmailSignInExpiration           *strfmt.DateTime `json:"email_sign_in_expiration,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	EmailUnsubscribeCode            *string          `json:"email_unsubscribe_code,omitempty" gorm:"type:varchar(256)"`
	EmailUnsubscribeExpiration      *strfmt.DateTime `json:"email_unsubscribe_expiration,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
}

type AgentMessageResponse struct {
	AgentMessage
	Message     *Message     `json:"message,omitempty"`
	MessageType *MessageType `json:"message_type,omitempty" gorm:"foreignKey:MessageID"`
}
