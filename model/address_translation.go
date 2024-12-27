package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AddressTranslation struct
type AddressTranslation struct {
	basic.Base
	basic.DataOwner
	AddressTranslationAPI
	AddressID *uuid.UUID `json:"address_id,omitempty" gorm:"type:varchar(36);uniqueIndex:address_id_translation_unique;not null" swaggertype:"string" format:"uuid"`
}

// AddressTranslationAPI struct
type AddressTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:address_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AddressLine        *string `json:"address_line,omitempty" gorm:"type:varchar(512)"`
	BuildingRoom       *string `json:"building_room,omitempty" gorm:"type:varchar(256)"`
	CityOther          *string `json:"city_other,omitempty" gorm:"type:varchar(64)"`
	County             *string `json:"county,omitempty" gorm:"type:varchar(32)"`
	PostalCode         *string `json:"postal_code,omitempty" gorm:"type:varchar(16)"`
	StateProvinceOther *string `json:"state_province_other,omitempty" gorm:"type:varchar(128)"`
	StreetNumber       *string `json:"street_number,omitempty" gorm:"type:varchar(64)"`
}
