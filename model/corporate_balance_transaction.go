package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type CorporateBalanceTransaction struct {
	basic.Base
	basic.DataOwner
	CorporateBalanceTransactionAPI
	Corporate *Corporate `json:"corporate,omitempty" gorm:"foreignKey:CorporateID;references:ID"`
	// Reservation *Reservation `json:"reservation,omitempty" gorm:"foreignKey:ReservationID;references:ID"`
}

type CorporateBalanceTransactionAPI struct {
	CorporateID               *uuid.UUID       `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`                                            // Corporate ID The reference to a corporate
	TransactionNumber         *string          `json:"transaction_number,omitempty" gorm:"type:varchar(36);not null;index:idx_transaction_number_deleted_at,unique,where:deleted_at is null"` // A unique transaction number.
	ReferenceNumber           *string          `json:"reference_number,omitempty" gorm:"type:varchar(36)"`                                                                                    //  A reference number, e.g. reservation code, invoice number, credit note number, etc.
	ReservationID             *uuid.UUID       `json:"reservation_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                   // The reference to a reservation.
	InvoiceID                 *string          `json:"invoice_id,omitempty" gorm:"type:varchar(36)"`                                                                                          // The reference to an invoice.
	CreditNoteID              *string          `json:"credit_note_id,omitempty" gorm:"type:varchar(36)"`                                                                                      // The reference to a credit note.
	TransactionType           *string          `json:"transaction_type,omitempty" gorm:"type:varchar(36)"`                                                                                    // The type of transaction, e.g. adjustment (debit or credit), credit note, applied to invoice.
	TransactionStatus         *int             `json:"transaction_status,omitempty" gorm:"type:smallint"`                                                                                     // Indicates the status of the transaction, e.g. created, authorized, captured, denied, voided, etc.
	TransactionTimestamp      *strfmt.DateTime `json:"transaction_timestamp,omitempty" gorm:"type:timestamptz" format:"date-time"`                                                            // Timestamp of the creation of the transaction (in UTC).
	TransactionLocalTimestamp *strfmt.DateTime `json:"transaction_local_timestamp,omitempty" gorm:"type:timestamptz" format:"date-time"`                                                      // Timestamp of the creation of the transaction (in Agent's Time Zone).
	CurrencyId                *string          `json:"currency_id,omitempty" gorm:"type:varchar(36)"`                                                                                         // The reference to the currency of the money amount.
	Amount                    *float64         `json:"amount,omitempty"`                                                                                                                      // The amount of the transaction.
	Balance                   *float64         `json:"balance,omitempty"`                                                                                                                     // The balance after the transaction was applied.
	Comment                   *string          `json:"comment,omitempty" gorm:"type:text"`                                                                                                    // The internal comment of the transaction.
	Description               *string          `json:"description,omitempty" gorm:"type:text"`
}
