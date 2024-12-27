package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Attraction model
type Attraction struct {
	basic.Base
	basic.DataOwner
	AttractionAPI
	AttractionTranslation *AttractionTranslation `json:"attraction_translation,omitempty"`                                                        // translation
	BusinessEntityID      *uuid.UUID             `json:"business_entity_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // Business Entity ID
	// BusinessEntity               *BusinessEntity                `json:"-"`                                                                                       // Business entity
	// AttractionAssetDesktop       *AttractionAsset               `json:"attraction_asset_desktop,omitempty"`                                                      // desktop asset
	// AttractionAssetTablet        *AttractionAsset               `json:"attraction_asset_tablet,omitempty"`                                                       // tablet asset
	// AttractionAssetMobile        *AttractionAsset               `json:"attraction_asset_mobile,omitempty"`                                                       // mobile asset
	// AttractionCategoryAttraction []AttractionCategoryAttraction `json:"attraction_category_attraction,omitempty"`                                                // Hotel Amenity Category Hotel Amenity Type
}

// AttractionAPI Attraction API
type AttractionAPI struct {
	AttractionName *string `json:"attraction_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" validate:"required" example:"Roller Coaster"` // Attraction Name
	Description    *string `json:"description,omitempty" gorm:"type:text" example:"Fantasy World"`                                                                                  // Description
}

// AttractionDetail model
type AttractionDetail struct {
	Attraction
	AttractionAssetDesktop  *MultimediaDescription `json:"attraction_asset_desktop,omitempty" gorm:"foreignKey:ID"`
	AttractionAssetTablet   *MultimediaDescription `json:"attraction_asset_tablet,omitempty" gorm:"foreignKey:ID"`
	AttractionAssetMobile   *MultimediaDescription `json:"attraction_asset_mobile,omitempty" gorm:"foreignKey:ID"`
	AttractionCategoryNames *string                `json:"attraction_category_names,omitempty" example:"play ground & roller coaster"` // Attraction Category Names joined by ampersand `&`
	ListAttractionCategory  *[]AttractionCategory  `json:"list_attraction_category,omitempty" gorm:"foreignKey:ID;"`                   // only required to fill on get detail
	PostalCode              *string                `json:"postal_code,omitempty"`                                                      // zip code / postal code
	Latitude                *float64               `json:"latitude,omitempty"`                                                         // Latitude
	Longitude               *float64               `json:"longitude,omitempty"`                                                        // Longitude
	AddressLine             *string                `json:"address_line,omitempty"`                                                     // Address Line
	Email                   *string                `json:"email,omitempty"`                                                            // Email address
	PhoneNumber             *string                `json:"phone_number,omitempty"`                                                     // Phone Number number
	FaxNumber               *string                `json:"fax_number,omitempty"`                                                       // Fax Number number
	ZoneID                  *uuid.UUID             `json:"zone_id,omitempty" swaggertype:"string" format:"uuid"`                       // zone ID
	ZoneName                *string                `json:"zone_name,omitempty"`                                                        // Zone
	DestinationID           *uuid.UUID             `json:"destination_id,omitempty" swaggertype:"string" format:"uuid"`                // destination ID
	DestinationName         *string                `json:"destination_name,omitempty"`                                                 // Destination
	CityID                  *uuid.UUID             `json:"city_id,omitempty" swaggertype:"string" format:"uuid"`                       // City ID
	CityName                *string                `json:"city_name,omitempty"`                                                        // City
	CountryID               *uuid.UUID             `json:"country_id,omitempty" swaggertype:"string" format:"uuid"`                    // Country ID
	CountryName             *string                `json:"country_name,omitempty"`                                                     // Country
	StateProvinceID         *uuid.UUID             `json:"state_province_id,omitempty" swaggertype:"string" format:"uuid"`             // State Province ID
	StateProvinceName       *string                `json:"state_province_name,omitempty"`                                              // State Province

	AttractionInternalData
}

// AttractionInternalData - only used for passing data
type AttractionInternalData struct {
	InternalAddressID                  *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
	InternalBusinessEntityEmailID      *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
	InternalPhoneBusinessEntityPhoneID *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
	InternalFaxBusinessEntityPhoneID   *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
	InternalAttractionAssetDesktopID   *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
	InternalAttractionAssetTabletID    *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
	InternalAttractionAssetMobileID    *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
}

type AttractionRequest struct {
	AttractionAPI
	ListAttractionCategoryID *[]uuid.UUID `json:"list_attraction_category_id,omitempty"`                                       // Hotel Amenity Category Hotel Amenity Type
	AttractionAssetDesktopID *uuid.UUID   `json:"attraction_asset_desktop_id,omitempty" validate:"required"`                   // desktop asset
	AttractionAssetTabletID  *uuid.UUID   `json:"attraction_asset_tablet_id,omitempty" validate:"required"`                    // tablet asset
	AttractionAssetMobileID  *uuid.UUID   `json:"attraction_asset_mobile_id,omitempty" validate:"required"`                    // mobile asset
	PostalCode               *string      `json:"postal_code,omitempty" example:"15414"`                                       // zip code / postal code
	Latitude                 *float64     `json:"latitude,omitempty" example:"-0.12831928"`                                    // Latitude
	Longitude                *float64     `json:"longitude,omitempty" example:"128.31928"`                                     // Longitude
	AddressLine              *string      `json:"address_line,omitempty" example:"Jl. TB Simatupang"`                          // Address Line
	DestinationID            *uuid.UUID   `json:"destination_id,omitempty" swaggertype:"string" format:"uuid"`                 // destination ID
	CityID                   *uuid.UUID   `json:"city_id,omitempty" swaggertype:"string" format:"uuid" validate:"required"`    // City ID
	CountryID                *uuid.UUID   `json:"country_id,omitempty" swaggertype:"string" format:"uuid" validate:"required"` // Country ID
	StateProvinceID          *uuid.UUID   `json:"state_province_id,omitempty" swaggertype:"string" format:"uuid"`              // State Province Id
	ZoneID                   *uuid.UUID   `json:"zone_id,omitempty" swaggertype:"string" format:"uuid"`                        // zone ID
	Email                    *string      `json:"email,omitempty" validate:"omitempty,lte=0|email" example:"example@mail.com"`
	PhoneNumber              *string      `json:"phone_number,omitempty" validate:"omitempty,lte=0|customphonenumber" example:"+6221911"` // Phone Number number
	FaxNumber                *string      `json:"fax_number,omitempty" example:"+6221911"`                                                // Fax Number number
}

// // AttractionList model
// type AttractionList struct {
// 	Attraction
// 	PostalCode              *string          `json:"postal_code,omitempty" example:"15414"`                                      // zip code / postal code
// 	Latitude                *float64         `json:"latitude,omitempty" example:"-0.12831928"`                                   // Latitude
// 	Longitude               *float64         `json:"longitude,omitempty" example:"128.31928"`                                    // Longitude
// 	AddressLine             *string          `json:"address_line,omitempty" example:"Jl. TB Simatupang"`                         // Address Line
// 	Email                   *string          `json:"email,omitempty" example:"example@mail.com" validate:"omitempty,email"`      // Email address
// 	PhoneNumber             *string          `json:"phone_number,omitempty" example:"+6221911"`                                  // Phone Number number
// 	FaxNumber               *string          `json:"fax_number,omitempty" example:"+6221911"`                                    // Fax Number number
// 	AttractionCategoryNames *string          `json:"attraction_category_names,omitempty" example:"play ground & roller coaster"` // Attraction Category Names joined by ampersand `&`
// 	AttractionAssetDesktop  *AttractionAsset `json:"-" gorm:"-" swaggerignore:"true"`                                            // Attraction Asset Desktop
// 	AttractionAssetTablet   *AttractionAsset `json:"-" gorm:"-" swaggerignore:"true"`                                            // Attraction Asset Tablet
// 	AttractionAssetMobile   *AttractionAsset `json:"-" gorm:"-" swaggerignore:"true"`                                            // Attraction Asset Mobile
// 	ZoneID                  *uuid.UUID       `json:"zone_id,omitempty" swaggertype:"string" format:"uuid"`                       // zone ID
// 	Zone                    *Zone            `json:"zone,omitempty"`                                                             // Zone
// 	DestinationID           *uuid.UUID       `json:"destination_id,omitempty" swaggertype:"string" format:"uuid"`                // destination ID
// 	Destination             *Destination     `json:"destination,omitempty"`                                                      // Destination
// 	CityID                  *uuid.UUID       `json:"city_id,omitempty" swaggertype:"string" format:"uuid"`                       // City ID
// 	City                    *City            `json:"city,omitempty"`                                                             // City
// 	CountryID               *uuid.UUID       `json:"country_id,omitempty" swaggertype:"string" format:"uuid"`                    // Country ID
// 	Country                 *Country         `json:"country,omitempty"`                                                          // Country
// 	StateProvinceID         *uuid.UUID       `json:"state_province_id,omitempty" swaggertype:"string" format:"uuid"`             // State Province ID
// 	StateProvince           *StateProvince   `json:"state_province,omitempty"`                                                   // State Province
// }

// // BeforeCreate attraction
// func (a *Attraction) BeforeCreate(tx *gorm.DB) error {
// 	// calling super class method
// 	err := a.Base.BeforeCreate(tx)

// 	businessEntity := BusinessEntity{}
// 	// create business entity to get an id
// 	tx.Create(&businessEntity)

// 	address := Address{}
// 	address.PostalCode = a.PostalCode
// 	address.Latitude = a.Latitude
// 	address.Longitude = a.Longitude
// 	address.AddressLine = a.AddressLine
// 	address.ZoneID = a.ZoneID
// 	address.DestinationID = a.DestinationID
// 	address.CountryID = a.CountryID
// 	address.CityID = a.CityID
// 	address.StateProvinceID = a.StateProvinceID
// 	// create address to get an id
// 	tx.Create(&address)

// 	businessEntityAddress := BusinessEntityAddress{}
// 	addressType := AddressType{}
// 	b := tx.First(&addressType, `address_type_name = ?`, `Contact`)
// 	if b.RowsAffected == 1 {
// 		businessEntityAddress.AddressTypeID = addressType.ID
// 	}
// 	businessEntityAddress.AddressID = address.ID
// 	businessEntityAddress.BusinessEntityID = businessEntity.ID
// 	// create business entity address to get an id
// 	tx.Create(&businessEntityAddress)

// 	businessEntityEmail := BusinessEntityEmail{}
// 	emailAddressType := EmailAddressType{}
// 	b2 := tx.First(&emailAddressType, `email_address_type_name = ?`, `Business`)
// 	if b2.RowsAffected == 1 {
// 		businessEntityEmail.EmailAddressTypeID = emailAddressType.ID
// 	}
// 	businessEntityEmail.BusinessEntityID = businessEntity.ID
// 	businessEntityEmail.Email = a.Email
// 	tx.Create(&businessEntityEmail)

// 	businessEntityPhone := BusinessEntityPhone{}
// 	phoneType := PhoneTechnologyType{}
// 	r := tx.First(&phoneType, `phone_technology_type_name = ?`, `Voice`)
// 	if r.RowsAffected == 1 {
// 		businessEntityPhone.PhoneTechnologyTypeID = phoneType.ID
// 	}
// 	phoneLocationType := PhoneLocationType{}
// 	b3 := tx.First(&phoneLocationType, `phone_location_type_name = ?`, `Office`)
// 	if b3.RowsAffected == 1 {
// 		businessEntityPhone.PhoneLocationTypeID = phoneLocationType.ID
// 	}
// 	businessEntityPhone.BusinessEntityID = businessEntity.ID
// 	businessEntityPhone.PhoneNumber = a.PhoneNumber
// 	tx.Create(&businessEntityPhone)

// 	businessEntityFax := BusinessEntityPhone{}
// 	faxType := PhoneTechnologyType{}
// 	r2 := tx.First(&faxType, `phone_technology_type_name = ?`, `Fax`)
// 	if r2.RowsAffected == 1 {
// 		businessEntityFax.PhoneTechnologyTypeID = faxType.ID
// 	}
// 	faxLocationType := PhoneLocationType{}
// 	b4 := tx.First(&faxLocationType, `phone_location_type_name = ?`, `Office`)
// 	if b4.RowsAffected == 1 {
// 		businessEntityFax.PhoneLocationTypeID = faxLocationType.ID
// 	}
// 	businessEntityFax.BusinessEntityID = businessEntity.ID
// 	businessEntityFax.PhoneNumber = a.FaxNumber
// 	tx.Create(&businessEntityFax)

// 	a.BusinessEntity = &businessEntity
// 	a.BusinessEntityID = businessEntity.ID

// 	return err
// }

// // BeforeUpdate attraction
// func (a *Attraction) BeforeUpdate(tx *gorm.DB) error {
// 	// calling super class method
// 	err := a.Base.BeforeUpdate(tx)
// 	if nil == a.BusinessEntityID {
// 		return a.BeforeCreate(tx)
// 	}

// 	r0 := tx.First(&BusinessEntity{}, `id = ?`, a.BusinessEntityID)
// 	if r0.RowsAffected < 1 {
// 		return a.BeforeCreate(tx)
// 	}

// 	businessEntityAddress := BusinessEntityAddress{}
// 	r := tx.Joins(`Address`).First(&businessEntityAddress, `business_entity_id = ?`, a.BusinessEntityID)
// 	if r.RowsAffected == 1 {
// 		if nil != businessEntityAddress.Address {
// 			address := *businessEntityAddress.Address
// 			address.PostalCode = a.PostalCode
// 			address.Latitude = a.Latitude
// 			address.Longitude = a.Longitude
// 			address.AddressLine = a.AddressLine
// 			address.ZoneID = a.ZoneID
// 			address.DestinationID = a.DestinationID
// 			address.CountryID = a.CountryID
// 			address.CityID = a.CityID
// 			address.StateProvinceID = a.StateProvinceID
// 			tx.Updates(&address)
// 		}
// 	}

// 	businessEntityEmail := BusinessEntityEmail{}
// 	r2 := tx.First(&businessEntityEmail, `business_entity_id = ?`, a.BusinessEntityID)
// 	if r2.RowsAffected == 1 {
// 		businessEntityEmail.Email = a.Email
// 		tx.Updates(&businessEntityEmail)
// 	}

// 	phoneType := PhoneTechnologyType{}
// 	r3 := tx.First(&phoneType, `phone_technology_type_name = ?`, `Voice`)
// 	if r3.RowsAffected == 1 {
// 		businessEntityPhone := BusinessEntityPhone{}
// 		r4 := tx.First(&businessEntityPhone,
// 			`business_entity_id = ? AND phone_technology_type_id = ?`, a.BusinessEntityID, phoneType.ID)
// 		if r4.RowsAffected == 1 {
// 			businessEntityPhone.PhoneNumber = a.PhoneNumber
// 			tx.Updates(&businessEntityPhone)
// 		}
// 	}

// 	faxType := PhoneTechnologyType{}
// 	r5 := tx.First(&faxType, `phone_technology_type_name = ?`, `Fax`)
// 	if r5.RowsAffected == 1 {
// 		businessEntityFax := BusinessEntityPhone{}
// 		r6 := tx.First(&businessEntityFax,
// 			`business_entity_id = ? AND phone_technology_type_id = ?`, a.BusinessEntityID, faxType.ID)
// 		if r6.RowsAffected == 1 {
// 			businessEntityFax.PhoneNumber = a.FaxNumber
// 			tx.Updates(&businessEntityFax)
// 		}
// 	}

// 	return err
// }
