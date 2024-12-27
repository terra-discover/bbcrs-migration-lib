package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentGateway struct {
	basic.Base
	basic.DataOwner
	PaymentGatewayAPI
	PaymentGatewayTranslation *PaymentGatewayTranslation `json:"payment_gateway_translation,omitempty"` // Payment Gateway Translation
	PaymentGatewayAsset       *PaymentGatewayAsset       `json:"operating_system_asset,omitempty"`      // PaymentGatewayAsset
}

type PaymentGatewayAPI struct {
	PaymentGatewayCode  *string                 `json:"payment_gateway_code,omitempty" gorm:"type:varchar(36);index:,unique,where:length(payment_gateway_code) > 0 and deleted_at is null;not null;" swaggertype:"string"`  // The code of the payment gateway. Payment gateway code is derived from payment gateway name in lower case and underscore separated, e.g. "klik_bca".
	PaymentGatewayName  *string                 `json:"payment_gateway_name,omitempty" gorm:"type:varchar(256);index:,unique,where:length(payment_gateway_name) > 0 and deleted_at is null;not null;" swaggertype:"string"` // The name or title of the page, e.g. "Klik BCA"
	PaymentTypeID       *uuid.UUID              `json:"payment_type_id,omitempty" format:"uuid" gorm:"type:varchar(36);not null"`                                                                                           // The reference to the payment type
	Description         *string                 `json:"description,omitempty" gorm:"type:text"`                                                                                                                             // Description
	PaymentGatewayAsset *PaymentGatewayAssetAPI `json:"operating_system_asset,omitempty" gorm:"-"`                                                                                                                          // PaymentGatewayAsset
}

// GetPaymentGateway by query filter
func (p *PaymentGateway) GetPaymentGateway(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).First(&p)
}

func (p PaymentGateway) Seed() *[]PaymentGateway {
	type paymentTypeID string
	const (
		CC               paymentTypeID = "cb22b1e5-36de-4b0a-86b7-4db89b912499"
		BankTransfer     paymentTypeID = "774cf533-94e6-49df-a618-af6d33e17a37"
		InternetBanking  paymentTypeID = "2012eb52-4b6a-49c9-8bfc-4e48807df959"
		ConvenienceStore paymentTypeID = "29fd9885-44a4-4588-a4ae-e8ec62202423"
		EWallet          paymentTypeID = "0116b8db-5e53-499f-8502-b2158f435b67"
		CardlessCredit   paymentTypeID = "4c6d302a-9345-41d2-b99e-e360d84c8ede"
	)

	// Payment Gateway Code is Midtrans payemnt type
	initialDatas := [][4]string{
		// CC
		{
			"credit_card", "Credit Card", string(CC),
		},
		// BankTransfer
		{
			"bank_transfer", "Bank Transfer", string(BankTransfer),
		},
		{
			"echannel", "e-Channel", string(BankTransfer),
		},
		// InternetBanking
		{
			"bca_klikpay", "BCA KlikPay", string(InternetBanking),
		},
		{
			"bca_klikbca", "KlikBCA", string(InternetBanking),
		},
		{
			"bri_epay", "BRImo", string(InternetBanking),
		},
		{
			"cimb_clicks", "CIMBClicks (OctoClicks)", string(InternetBanking),
		},
		{
			"danamon_online", "Danamon Online Banking (DOB)", string(InternetBanking),
		},
		{
			"uob_ezpay", "UOB Ezpay", string(InternetBanking),
		},
		// EWallet
		{
			"qris", "QRIS", string(EWallet),
		},
		{
			"gopay", "GoPay", string(EWallet),
		},
		{
			"shopeepay", "ShopeePay", string(EWallet),
		},
		// ConvenienceStore
		{
			"cstore", "Convenience Store", string(ConvenienceStore),
		},
		// CardlessCredit
		{
			"akulaku", "Akulaku Paylater", string(CardlessCredit),
		},
		{
			"kredivo", "Kredivo", string(CardlessCredit),
		},
	}

	seed := []PaymentGateway{}

	for _, item := range initialDatas {
		code := item[0]
		name := item[1]
		typeID := item[2]

		paymentGateway := PaymentGateway{}
		paymentGateway.PaymentGatewayCode = &code
		paymentGateway.PaymentGatewayName = &name
		paymentGateway.PaymentTypeID = lib.UUIDPtr(uuid.MustParse(typeID))

		seed = append(seed, paymentGateway)
	}

	return &seed
}
