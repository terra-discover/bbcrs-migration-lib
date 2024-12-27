package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type PaymentTransaction struct {
	basic.Base
	basic.DataOwner
	PaymentTransactionAPI
	// Reservation *Reservation `json:"reservation,omitempty" gorm:"foreignKey:ReservationID;references:ID"`
	PaymentGateway *PaymentGateway `json:"payment_gateway,omitempty" gorm:"foreignKey:PaymentGatewayID;references:ID"`
}

type PaymentTransactionAPI struct {
	TransactionNumber         *string          `json:"transaction_number,omitempty" gorm:"type:varchar(36);not null"`                       // A unique transaction number.
	ReferenceNumber           *string          `json:"reference_number,omitempty" gorm:"type:varchar(36)"`                                  //  A reference number, e.g. reservation code, invoice number, credit note number, etc.
	ReservationID             *uuid.UUID       `json:"reservation_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // The reference to a reservation.
	TransactionType           *string          `json:"transaction_type,omitempty" gorm:"type:varchar(36)"`                                  // The type of transaction, e.g. adjustment (debit or credit), credit note, applied to invoice.
	TransactionStatus         *int             `json:"transaction_status,omitempty" gorm:"type:smallint"`                                   // Indicates the status of the transaction, e.g. created, authorized, captured, denied, voided, etc.
	TransactionTimestamp      *strfmt.DateTime `json:"transaction_timestamp,omitempty" gorm:"type:timestamptz" format:"date-time"`          // Timestamp of the creation of the transaction (in UTC).
	TransactionLocalTimestamp *strfmt.DateTime `json:"transaction_local_timestamp,omitempty" gorm:"type:timestamptz" format:"date-time"`    // Timestamp of the creation of the transaction (in Agent's Time Zone).
	CurrencyID                *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`    // The reference to the currency of the money amount.
	Amount                    *float64         `json:"amount,omitempty"`                                                                    // The amount of the transaction.
	SessionID                 *uuid.UUID       `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	PaymentGatewayID          *uuid.UUID       `json:"payment_gateway_id,omitempty" gorm:"type:varchar(36)"`
	MerchantID                *string          `json:"merchant_id,omitempty" gorm:"type:varchar(256)"`
	ProcessorID               *string          `json:"processor_id,omitempty" gorm:"type:varchar(256)"`
	PartnerTransactionNumber  *string          `json:"partner_transaction_number,omitempty" gorm:"type:varchar(36)"`
	PartnerTransactionStatus  *string          `json:"partner_transaction_status,omitempty" gorm:"type:varchar(36)"`
	ApprovalCode              *string          `json:"approval_code,omitempty" gorm:"type:varchar(36)"`
	ApprovalTimestamp         *strfmt.DateTime `json:"approval_timestamp,omitempty" gorm:"type:timestamptz" format:"date-time"` // Timestamp of the creation of the transaction (in UTC).
	ApprovalLocalTimestamp    *strfmt.DateTime `json:"approval_local_timestamp,omitempty" gorm:"type:timestamptz" format:"date-time"`
	Comment                   *string          `json:"comment,omitempty" gorm:"type:text"`
	Description               *string          `json:"description,omitempty" gorm:"type:text"` // The internal comment of the transaction.
}
