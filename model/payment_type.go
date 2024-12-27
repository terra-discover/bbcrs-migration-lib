package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// PaymentType Payment Type
type PaymentType struct {
	basic.Base
	basic.DataOwner
	PaymentTypeAPI
	PaymentTypeTranslation *PaymentTypeTranslation `json:"payment_type_translation,omitempty"`
}

// PaymentTypeAPI Payment Type API
type PaymentTypeAPI struct {
	PaymentTypeCode *string `json:"payment_type_code,omitempty" gorm:"type:varchar(36);index:,unique,where:length(payment_type_code) > 0 and deleted_at is null;not null;" swaggertype:"string"`  // Payment Type Code
	PaymentTypeName *string `json:"payment_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:length(payment_type_name) > 0 and deleted_at is null;not null;" swaggertype:"string"` // Payment Type Name
}

// Seed PaymentType
func (paymentType *PaymentType) Seed() *[]PaymentType {
	initialDatas := [][3]string{
		{
			"cb22b1e5-36de-4b0a-86b7-4db89b912499", "CC", "Credit Card",
		},
		{
			"774cf533-94e6-49df-a618-af6d33e17a37", "BankTransfer", "Bank Transfer/ Virtual Account (ATM)",
		},
		{
			"2012eb52-4b6a-49c9-8bfc-4e48807df959", "InternetBanking", "Internet Banking",
		},
		{
			"29fd9885-44a4-4588-a4ae-e8ec62202423", "ConvenienceStore", "Convenience Store",
		},
		{
			"0116b8db-5e53-499f-8502-b2158f435b67", "EWallet", "e-Wallet",
		},
		{
			"4c6d302a-9345-41d2-b99e-e360d84c8ede", "CardlessCredit", "Cardless Credit",
		},
	}

	seed := []PaymentType{}

	for _, item := range initialDatas {
		id := item[0]
		code := item[1]
		name := item[2]

		paymentType := PaymentType{}
		paymentType.ID = lib.UUIDPtr(uuid.MustParse(id))
		paymentType.PaymentTypeCode = &code
		paymentType.PaymentTypeName = &name

		seed = append(seed, paymentType)
	}

	return &seed
}
