package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Person Person
type Person struct {
	basic.Base
	basic.DataOwner
	PersonAPI
	NamePrefix     *NamePrefix     `json:"name_prefix,omitempty" gorm:"foreignKey:NamePrefixID;references:ID"`         // Name Prefix
	CitizenCountry *Country        `json:"citizen_country,omitempty" gorm:"foreignKey:CitizenCountryID;references:ID"` // Citizen Country
	MaritalStatus  *MaritalStatus  `json:"marital_status,omitempty" gorm:"foreignKey:MaritalStatusID;references:ID"`   // Marital Status
	Gender         *Gender         `json:"gender,omitempty" gorm:"foreignKey:GenderID;references:ID"`                  // Gender
	Religion       *Religion       `json:"religion,omitempty" gorm:"foreignKey:ReligionID;references:ID"`              // Religion
	BusinessEntity *BusinessEntity `json:"business_entity,omitempty" gorm:"foreignKey:BusinessEntityID;references:ID"` // Business Entity
	// TODO: prepare FlightCachingTraveller into 1 package
	Traveler *FlightCachingTraveller `json:"traveler,omitempty" gorm:"foreignKey:ID;references:PersonID"` // Traveler
	Infant   *FlightCachingTraveller `json:"infant,omitempty" gorm:"-"`
}

// PersonAPI Person API
type PersonAPI struct {
	GivenName        *string      `json:"given_name,omitempty" gorm:"type:varchar(128)" example:"Bayu"`             // Given Name
	MiddleName       *string      `json:"middle_name,omitempty" gorm:"type:varchar(128)" example:"Buana"`           // Middle Name
	NamePrefixID     *uuid.UUID   `json:"name_prefix_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`           // Name Prefix Id
	NameSuffix       *string      `json:"name_suffix,omitempty" gorm:"type:varchar(16)"`                            // Name Suffix
	NameTitle        *string      `json:"name_title,omitempty" gorm:"type:varchar(16)"`                             // Name Title
	Surname          *string      `json:"surname,omitempty" gorm:"type:varchar(128)" example:"Travel"`              // Surname
	SurnamePrefix    *string      `json:"surname_prefix,omitempty" gorm:"type:varchar(16)"`                         // Surname Prefix
	CitizenCountryID *uuid.UUID   `json:"citizen_country_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`       // Citizen Country Id
	BusinessEntityID *uuid.UUID   `json:"business_entity_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`       // Business Entity Id
	BirthDate        *strfmt.Date `json:"birth_date,omitempty" gorm:"type:date" format:"date" swaggertype:"string"` // Birth Date
	MaritalStatusID  *uuid.UUID   `json:"marital_status_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`        // Marital Status Id
	GenderID         *uuid.UUID   `json:"gender_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                // Gender Id
	ReligionID       *uuid.UUID   `json:"religion_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`              // Religion Id
}

// Seed data
func (d *Person) Seed(userID uuid.UUID) *[]Person {
	if lib.IsEmptyUUID(userID) {
		userID = uuid.New()
	}

	BusinessEntityID := uuid.New()
	data := []Person{}
	data = append(data, Person{
		Base: basic.Base{ID: &userID},
		PersonAPI: PersonAPI{
			GivenName:        lib.Strptr("Bayu"),
			MiddleName:       lib.Strptr("buana"),
			Surname:          lib.Strptr("buana"),
			BusinessEntityID: &BusinessEntityID,
		},
	})

	return &data
}
