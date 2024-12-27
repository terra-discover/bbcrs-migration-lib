package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Country struct is not writable by request body
type Country struct {
	basic.Base
	basic.DataOwner
	CountryAPI
	CountryTranslation *CountryTranslation `json:"country_translation,omitempty"`                                 // country translation
	Region             *Region             `json:"region,omitempty" gorm:"foreignKey:RegionID;references:ID"`     // region
	Timezone           *TimeZone           `json:"timezone,omitempty" gorm:"foreignKey:TimezoneID;references:ID"` // timezone
	Language           *Language           `json:"language,omitempty" gorm:"foreignKey:LanguageID;references:ID"` // language
	Currency           *Currency           `json:"currency,omitempty" gorm:"foreignKey:CurrencyID;references:ID"` // currency
}

// CountryAPI struct is writable by request body
type CountryAPI struct {
	CountryCode       *string    `json:"country_code,omitempty" gorm:"type:varchar(2);index:,unique,where:deleted_at is null;not null" example:"id"`                                                                                                 // Country code
	CountryName       *string    `json:"country_name,omitempty" gorm:"type:varchar(64);index:,unique,where:deleted_at is null;not null" example:"Indonesia"`                                                                                         // Country name
	CountryNativeName *string    `json:"country_native_name,omitempty" gorm:"type:varchar(256)" example:"Indonesia"`                                                                                                                                 // Country native name
	CountryAlpha3Code *string    `json:"country_alpha_3_code,omitempty" gorm:"column:country_alpha_3_code;type:varchar(3);index:,unique,where:deleted_at is null AND country_alpha_3_code is not null AND country_alpha_3_code != ''" example:"IDN"` // Country alpha 3 codes
	NumericCode       *string    `json:"numeric_code,omitempty" gorm:"type:varchar(3)" example:"360"`                                                                                                                                                // Country numeric code: [https://en.wikipedia.org/wiki/ISO_3166-1_numeric](https://en.wikipedia.org/wiki/ISO_3166-1_numeric)
	CountryAccessCode *string    `json:"country_access_code,omitempty" gorm:"type:varchar(36)" example:"62"`                                                                                                                                         // Country access code
	TimezoneID        *uuid.UUID `json:"timezone_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                                                                                                                                // timezone ID
	CurrencyID        *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                                                                                                                                // currency ID
	LanguageID        *uuid.UUID `json:"language_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                                                                                                                                // language ID
	RegionID          *uuid.UUID `json:"region_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                                                                                                                                  // region ID
	Nationality       *string    `json:"nationality,omitempty" gorm:"type:varchar(64)"`                                                                                                                                                              // nationality
}
