package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomStayPayment Model RoomStayPayment
type RoomStayPayment struct {
	basic.Base
	basic.DataOwner
	RoomStayPaymentAPI
	PaymentStatus *Status `json:"payment_status,omitempty" gorm:"foreignKey:PaymentStatus;references:StatusCode"`
}

// RoomStayPaymentAPI Model RoomStayPaymentAPI
type RoomStayPaymentAPI struct {
	RoomStayID                    *uuid.UUID `json:"hotel_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string"  format:"uuid" example:"fcee9740-30ee-4d46-af2f-d58cf49f8642"`                // Room Stay ID
	PaymentStatus                 *int       `json:"room_stay_status,omitempty" gorm:"type:smallint" swaggertype:"number" format:"int32" example:"10"`                                                      // PaymentStatus
	CorporateBalanceTransactionID *uuid.UUID `json:"corporate_balance_transaction_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"  format:"uuid" example:"fcee9740-30ee-4d46-af2f-d58cf49f8642"` // CorporateBalanceTransactionID
	PaymentCardID                 *uuid.UUID `json:"payment_card_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"  format:"uuid" example:"fcee9740-30ee-4d46-af2f-d58cf49f8642"`                  // PaymentCardID
	PaymentTransactionID          *uuid.UUID `json:"payment_transaction_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"  format:"uuid" example:"fcee9740-30ee-4d46-af2f-d58cf49f8642"`           // PaymentTransactionID
	BankAccountID                 *uuid.UUID `json:"bank_account_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"  format:"uuid" example:"fcee9740-30ee-4d46-af2f-d58cf49f8642"`                  // BankAccountID
	RewardBalanceTransactionID    *uuid.UUID `json:"reward_balance_transaction_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"  format:"uuid" example:"fcee9740-30ee-4d46-af2f-d58cf49f8642"`    // RewardBalanceTransactionID
	VoucherTransactionID          *uuid.UUID `json:"voucher_transaction_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"  format:"uuid" example:"fcee9740-30ee-4d46-af2f-d58cf49f8642"`           // VoucherTransactionID
	Comment                       *string    `json:"comment,omitempty" gorm:"type:text" swaggertype:"string" example:"this is very good"`                                                                   // Comment
	Description                   *string    `json:"description,omitempty" gorm:"type:text" swaggertype:"string" example:"this is very good"`                                                               // Description
}
