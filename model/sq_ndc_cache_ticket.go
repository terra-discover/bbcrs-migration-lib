package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SqNdcCacheTicket | super class
type SqNdcCacheTicket struct {
	basic.Base
	basic.DataOwner
	SessionID          *uuid.UUID                `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	AgentID            *string                   `json:"agent_id,omitempty" gorm:"type:varchar(60)" default:"CGKSIN"`
	AgentType          *string                   `json:"agent_type,omitempty" gorm:"type:varchar(60)" default:"CGKSIN"`
	PassengerReference *string                   `json:"passenger_reference,omitempty" gorm:"type:varchar(60)" default:"PAX1"`
	TicketDocument     *SqNdcCacheTicketDocument `json:"ticket_document,omitempty" gorm:"foreignKey:TicketID;references:ID"`
}

// SqNdcCacheTicketDocument | child class
type SqNdcCacheTicketDocument struct {
	basic.Base
	basic.DataOwner
	SessionID        *uuid.UUID                `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	TicketID         *uuid.UUID                `json:"ticket_id" gorm:"type:varchar(36)" format:"uuid"`  // TicketID
	TicketDocNbr     *string                   `json:"ticket_doc_nbr,omitempty" gorm:"type:varchar(60)" default:"6182434954653"`
	Type             *string                   `json:"type,omitempty" gorm:"type:varchar(60)" default:"T"`
	NumberofBooklets *int                      `json:"number_of_booklets,omitempty"  default:"1"`
	DateOfIssue      *strfmt.Date              `json:"date_of_issue,omitempty" gorm:"type:date" default:"2022-04-20"`
	ReportingType    *string                   `json:"reporting_type,omitempty" gorm:"type:varchar(60)" default:"I"`
	TicketCoupon     *[]SqNdcCacheTicketCoupon `json:"ticket_coupon,omitempty" gorm:"foreignKey:TicketDocumentID;references:ID"`
}

// SqNdcCacheTicketCoupon
type SqNdcCacheTicketCoupon struct {
	basic.Base
	basic.DataOwner
	SessionID             *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                         // session ID
	TicketDocumentID      *uuid.UUID `json:"ticket_document_id" gorm:"type:varchar(36)" format:"uuid"`                 // TicketDocumentID
	CouponNumber          *int       `json:"coupon_number,omitempty"  default:"1"`                                     // CouponNumber
	CouponReference       *string    `json:"coupon_reference,omitempty" gorm:"type:varchar(60)" default:"SEG1"`        // CouponReference
	ServiceReferences     *string    `json:"service_references,omitempty" gorm:"type:varchar(60)" default:"SEG1"`      // ServiceReferences
	InConnectionDocNbr    *int       `json:"in_connection_doc_nbr,omitempty"  default:"6182437865764"`                 // InConnectionDocNbr
	InConnectonCpnNbr     *int       `json:"in_connecton_cpn_nbr,omitempty"  default:"1"`                              // InConnectonCpnNbr
	ReasonForIssuanceRFIC *string    `json:"reason_for_issuance_rfic,omitempty" gorm:"type:varchar(10)" default:"A"`   // ReasonForIssuanceRFIC
	ReasonForIssuanceCode *string    `json:"reason_for_issuance_code,omitempty" gorm:"type:varchar(10)" default:"0B5"` // ReasonForIssuanceCode
	FareBasisCode         *string    `json:"fare_basis_code,omitempty" gorm:"type:varchar(60)" default:"T"`
	Status                *string    `json:"status,omitempty" gorm:"type:varchar(60)" default:"I"`
	SettlementAuthCode    *string    `json:"settlement_auth_code,omitempty" gorm:"varchar(50)"`          // SettlementAuthCode when success void of an order
	ExpectedRefundAmount  *float64   `json:"expected_refund_amount,omitempty" gorm:"type:decimal(22,2)"` // ExpectedRefundAmount perTicket perTraveler when success void of an order
}
