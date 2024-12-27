package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerPaymentGateway Integration Partner Payment Gateway
type IntegrationPartnerPaymentGateway struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerPaymentGatewayAPI
	PaymentGateway     *PaymentGateway     `json:"payment_gateway" gorm:"foreignKey:PaymentGatewayID;references:ID"`         // Payment Gateway
	IntegrationPartner *IntegrationPartner `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"` // Integration Partner
	Currency           *Currency           `json:"currency" gorm:"foreignKey:CurrencyID;references:ID"`                      // Currency
	Bank               *Bank               `json:"bank" gorm:"foreignKey:BankID;references:ID"`                              // Bank
}

// IntegrationPartnerPaymentGatewayAPI Integration Partner Payment Gateway API
type IntegrationPartnerPaymentGatewayAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"` // Integration Partner ID
	PaymentGatewayID     *uuid.UUID `json:"payment_gateway_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"`     // Payment Gateway ID
	MerchantID           *string    `json:"merchant_id,omitempty" gorm:"type:varchar(256)"`                                                       // Merchant ID
	TerminalID           *string    `json:"terminal_id,omitempty" gorm:"type:varchar(256)"`                                                       // Terminal ID
	ChannelCode          *string    `json:"channel_code,omitempty" gorm:"type:varchar(256)"`                                                      // Channel Code
	CurrencyID           *uuid.UUID `json:"currency_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`                     // Currency ID
	TransactionUrl       *string    `json:"transaction_url,omitempty" gorm:"type:varchar(256)"`                                                   // Transaction Url
	NotificationUrl      *string    `json:"notification_url,omitempty" gorm:"type:varchar(256)"`                                                  // Notification Url
	ClientKey            *string    `json:"client_key,omitempty" gorm:"type:varchar(256)"`                                                        // Client Key
	ServerKey            *string    `json:"server_key,omitempty" gorm:"type:varchar(256)"`                                                        // Server Key
	BankID               *uuid.UUID `json:"bank_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`                         // Bank ID
	VirtualAccountNumber *string    `json:"virtual_account_number,omitempty" gorm:"type:varchar(256)"`                                            // Virtual Account Number
	ConvenienceStoreCode *string    `json:"convenience_store_code,omitempty" gorm:"type:varchar(256)"`                                            // Convenience Store Code
}

// IntegrationPartnerPaymentGatewayRequest Integration Partner Payment Gateway Request
type IntegrationPartnerPaymentGatewayRequest struct {
	PaymentGatewayCode   *string    `json:"payment_gateway_code,omitempty" example:"credit_card" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"` // Payment Gateway Code
	PaymentGatewayName   *string    `json:"payment_gateway_name,omitempty" example:"Credit Card" gorm:"type:varchar(256)"`                                                // Payment Gateway Name
	MerchantID           *string    `json:"merchant_id,omitempty" gorm:"type:varchar(256)"`                                                                               // Merchant ID
	TerminalID           *string    `json:"terminal_id,omitempty" gorm:"type:varchar(256)"`                                                                               // Terminal ID
	ChannelCode          *string    `json:"channel_code,omitempty" gorm:"type:varchar(256)"`                                                                              // Channel Code
	CurrencyID           *uuid.UUID `json:"currency_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`                                             // Currency ID
	TransactionUrl       *string    `json:"transaction_url,omitempty" gorm:"type:varchar(256)"`                                                                           // Transaction Url
	NotificationUrl      *string    `json:"notification_url,omitempty" gorm:"type:varchar(256)"`                                                                          // Notification Url
	ClientKey            *string    `json:"client_key,omitempty" gorm:"type:varchar(256)"`                                                                                // Client Key
	ServerKey            *string    `json:"server_key,omitempty" gorm:"type:varchar(256)"`                                                                                // Server Key
	BankID               *uuid.UUID `json:"bank_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`                                                 // Bank ID
	VirtualAccountNumber *string    `json:"virtual_account_number,omitempty" gorm:"type:varchar(256)"`                                                                    // Virtual Account Number
	ConvenienceStoreCode *string    `json:"convenience_store_code,omitempty" gorm:"type:varchar(256)"`                                                                    // Convenience Store Code
}
