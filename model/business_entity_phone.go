package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessEntityPhone struct
type BusinessEntityPhone struct {
	basic.Base
	basic.DataOwner
	BusinessEntityPhoneAPI
	BusinessEntity      *BusinessEntity      `json:"business_entity,omitempty"`       // Business Entity
	PhoneLocationType   *PhoneLocationType   `json:"phone_location_type,omitempty"`   // Phone Location Type
	PhoneTechnologyType *PhoneTechnologyType `json:"phone_technology_type,omitempty"` // Phone Technology Type
	PhoneUseType        *PhoneUseType        `json:"phone_use_type,omitempty"`        // Phone Use Type
}

// BusinessEntityPhoneAPI for secure request body API
type BusinessEntityPhoneAPI struct {
	BusinessEntityID      *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36);not null;index:idx_business_entity_phone__business_entity_id,unique,where:deleted_at is null;"` // Business Entity Id
	PhoneLocationTypeID   *uuid.UUID `json:"phone_location_type_id,omitempty" gorm:"type:varchar(36);index:idx_business_entity_phone__business_entity_id,unique,where:deleted_at is null;"`      // Phone Location Type Id
	PhoneTechnologyTypeID *uuid.UUID `json:"phone_technology_type_id,omitempty" gorm:"type:varchar(36);index:idx_business_entity_phone__business_entity_id,unique,where:deleted_at is null;"`    // Phone Technology Type Id
	PhoneUseTypeID        *uuid.UUID `json:"phone_use_type_id,omitempty" gorm:"type:varchar(36);index:idx_business_entity_phone__business_entity_id,unique,where:deleted_at is null;"`           // Phone Use Type Id
	CountryAccessCode     *string    `json:"country_access_code,omitempty" gorm:"type:varchar(3)"`                                                                                               // Country Access Code
	AreaCityCode          *string    `json:"area_city_code,omitempty" gorm:"type:varchar(8)"`                                                                                                    // Area City Code
	PhoneNumber           *string    `json:"phone_number,omitempty" gorm:"type:varchar(32)"`                                                                                                     // Phone Number
	Extension             *string    `json:"extension,omitempty" gorm:"type:varchar(5)"`                                                                                                         // Extension
	IsFormatted           *bool      `json:"is_formatted,omitempty"`                                                                                                                             // Is Formatted
	IsPrimary             *bool      `json:"is_primary,omitempty" gorm:"type:bool;index:idx_business_entity_phone__business_entity_id,unique,where:deleted_at is null;"`                         // Is Primary
}
