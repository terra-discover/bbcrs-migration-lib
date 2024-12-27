package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// ChargeType Charge Type
type ChargeType struct {
	basic.Base
	basic.DataOwner
	ChargeTypeAPI
	ChargeTypeTranslation *ChargeTypeTranslation `json:"charge_type_translation,omitempty"` // Charge Type Translation
}

// ChargeTypeAPI Charge Type API
type ChargeTypeAPI struct {
	ChargeTypeCode *int    `json:"charge_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`         // Charge Type Code
	ChargeTypeName *string `json:"charge_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Daily"` // Charge Type Name
}

// Seed data
func (s ChargeType) Seed() *[]ChargeType {
	data := []ChargeType{}
	items := []string{
		"Daily",
		"Hourly",
		"Half day",
		"Additions per stay",
		"Per occurrence",
		"Per event",
		"Per person",
		"First use",
		"One time use",
		"Per minute",
		"Per function",
		"Per stay",
		"Complimentary",
		"Other",
		"Maximum charge",
		"Over-minute charge",
		"Weekly",
		"Per room per stay",
		"Per room per night",
		"Per person per stay",
		"Per person per night",
		"Minimum charge",
		"Per rental",
		"Per item",
		"Per room",
		"Per reservation/booking",
		"Per gallon",
		"Per dozen",
		"Per tray",
		"Per order",
		"Per unit",
		"One way",
		"Round trip",
		"Per ticket",
		"Per one way",
		"Per segment",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]

	switchName:
		switch name {
		case "Per ticket":
			{
				code = 1001
				break switchName
			}
		case "Per one way":
			{
				code = 1002
				break switchName
			}
		case "Per segment":
			{
				code = 1003
				break switchName
			}
		}

		data = append(data, ChargeType{
			ChargeTypeAPI: ChargeTypeAPI{
				ChargeTypeCode: &code,
				ChargeTypeName: &name,
			},
		})
	}

	return &data
}
