package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirItineraryPayment Air Itinerary Payment
type AirItineraryPayment struct {
	basic.Base
	basic.DataOwner
	AirItineraryPaymentAPI
	PaymentStatus *Status `json:"payment_status,omitempty" gorm:"foreignKey:PaymentStatus;references:StatusCode"`
}

// AirItineraryPaymentAPI Air Itinerary Payment Api
type AirItineraryPaymentAPI struct {
	AirItineraryID                *uuid.UUID `json:"air_itinerary_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"` //
	PaymentStatus                 *int       `json:"itinerary_status,omitempty" gorm:"type:smallint"`                                           // Indicates the payment status of the air itinerary, e.g. guaranteed, hold, pre-paid, etc.
	CorporateBalanceTransactionID *uuid.UUID `json:"corporate_balance_transaction_id,omitempty" swaggertype:"string" format:"uuid"`             //
	PaymentCardID                 *uuid.UUID `json:"payment_card_id,omitempty" swaggertype:"string" format:"uuid"`                              //
	PaymentTransactionID          *uuid.UUID `json:"payment_transaction_id,omitempty" swaggertype:"string" format:"uuid"`                       //
	BankAccountID                 *uuid.UUID `json:"bank_account_id,omitempty" swaggertype:"string" format:"uuid"`                              //
	RewardBalanceTransactionID    *uuid.UUID `json:"reward_balance_transaction_id,omitempty" swaggertype:"string" format:"uuid"`                //
	VoucherTransactionID          *uuid.UUID `json:"voucher_transaction_id,omitempty" swaggertype:"string" format:"uuid"`                       //
	Comment                       *string    `json:"comment,omitempty" gorm:"type:text"`                                                        //
	Description                   *string    `json:"description,omitempty" gorm:"type:text"`                                                    //
}
