package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PassengerType Passenger Type
type PassengerType struct {
	basic.Base
	basic.DataOwner
	PassengerTypeAPI
	PassengerTypeTranslation *PassengerTypeTranslation `json:"passenger_type_translation,omitempty"`
}

// PassengerTypeAPI Passenger Type API
type PassengerTypeAPI struct {
	PassengerTypeCode   *int    `json:"passenger_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null" example:"1"`                                                                                                     // Passenger Type Code
	PassengerTypeName   *string `json:"passenger_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"ADULT"`                                                                                             // Passenger Type Name
	PassengerAlpha3Code *string `json:"passenger_alpha_3_code,omitempty" gorm:"column:passenger_alpha_3_code;type:varchar(3);index:,unique,where:deleted_at is null AND passenger_alpha_3_code IS NOT NULL AND passenger_alpha_3_code != ''" example:"ADT"` // Passenger Alpha 3 Code
}

// Seed data
func (p PassengerType) Seed() *[]PassengerType {
	data := []PassengerType{}
	items := []string{
		"Any passenger type|",
		"Additional occupant with adult|",
		"Additional occupant without adult|",
		"Adult|ADT",
		"Child|CHD",
		"Employee|",
		"Farmer|",
		"Female|",
		"Free adult|",
		"Free child|",
		"Government|",
		"Handicapped|",
		"Infant|INF",
		"Male|",
		"Military|",
		"Over 21|",
		"Over 65|",
		"Person aged between 4 and 11 years|",
		"Person between 11 and 99 years|",
		"Person between 12 and 99 years|",
		"Person between 13 and 99 years|",
		"Person between 14 and 99 years|",
		"Person between 15 and 99 years.|",
		"Person between 16 and 99 years.|",
		"Person between 17 and 99 years.|",
		"Person between 26 and 99 years.|",
		"Person between 4 and 12 years|",
		"Person between 4 and 13 years|",
		"Person between 4 and 16 years|",
		"Person between 5 and 12 years|",
		"Person between 5 and 16 years|",
		"Person between 6 and 12 years|",
		"Person between 6 and 14 years|",
		"Person between 6 and 15 years|",
		"Person between 6 and 16 years|",
		"Person between 6 and 17 years|",
		"Person between 7 and 12 years|",
		"Person between 7 and 15 years|",
		"Senior|",
		"Student|",
		"Teenager|",
		"Under 110 cm|",
		"Under 12|",
		"Under 150 cm|",
		"Under 17|",
		"Under 2|",
		"Under 21|",
		"Unknown|",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		var code int = i + 1
		var name string = contents[0]
		var code3 string = contents[1]
		data = append(data, PassengerType{
			PassengerTypeAPI: PassengerTypeAPI{
				PassengerTypeCode:   &code,
				PassengerAlpha3Code: &code3,
				PassengerTypeName:   &name,
			},
		})
	}

	return &data
}

// SeedOne PassengerType
func (p *PassengerType) SeedOne() *PassengerType {
	seed := PassengerType{
		PassengerTypeAPI: PassengerTypeAPI{
			PassengerTypeCode:   lib.Intptr(1),
			PassengerTypeName:   lib.Strptr("ADULT"),
			PassengerAlpha3Code: lib.Strptr("ADT"),
		},
	}
	_ = lib.Merge(seed, &p)
	return p
}
