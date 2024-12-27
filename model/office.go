package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Office Model
type Office struct {
	basic.Base
	basic.DataOwner
	OfficeAPI
	OfficeTranslation       *OfficeTranslation       `json:"office_translation,omitempty"`                                  // Office Translation
	AgentOffice             *AgentOffice             `json:"-" gorm:"foreignKey:OfficeID" swaggerignore:"true"`             // Agent Office
	OfficeOperationSchedule *OfficeOperationSchedule `json:"-" swaggerignore:"true"`                                        // Office Operation Schedule
	BusinessEntityID        *uuid.UUID               `json:"-" gorm:"type:varchar(36);" format:"uuid" swaggerignore:"true"` // Business Entity ID
	BusinessEntity          *BusinessEntity          `json:"-" swaggerignore:"true"`                                        // Business Entity
	CorporateGroup          []CorporateGroup         `json:"corporate_group"`
}

// OfficeAPI Office API
type OfficeAPI struct {
	AgentID             *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid"`                                                // Agent ID
	OfficeName          *string    `json:"office_name,omitempty" gorm:"type:varchar(256);not null" validate:"required" example:"PT. Bayu Buana"` // Office Name / Company Name / Branch Name
	Description         *string    `json:"description,omitempty" gorm:"type:text" example:"Bayu Buana Office"`                                   // Description
	BuildingRoom        *string    `json:"building_room,omitempty" gorm:"-" example:"LT170"`                                                     // Building Room
	AddressLine         *string    `json:"address_line,omitempty" gorm:"-" example:"BCA Tower"`                                                  // Address
	AddressTypeID       *uuid.UUID `json:"address_type_id,omitempty" gorm:"-" swaggertype:"string" format:"uuid"`                                // Address Type ID
	CountryID           *uuid.UUID `json:"country_id,omitempty" gorm:"-" swaggertype:"string" format:"uuid"`                                     // Country ID
	StateProvinceID     *uuid.UUID `json:"state_province_id,omitempty" gorm:"-" swaggertype:"string" format:"uuid"`                              // Steate Province ID
	CityID              *uuid.UUID `json:"city_id,omitempty" gorm:"-" swaggertype:"string" format:"uuid"`                                        // State Province ID
	PostalCode          *string    `json:"postal_code,omitempty" gorm:"-" example:"12345"`                                                       // Postal Code
	PhoneNumber         *string    `json:"phone_number,omitempty" gorm:"-" example:"021911"`                                                     // Phone Number
	PhoneLocationTypeID *uuid.UUID `json:"phone_location_type_id,omitempty" gorm:"-" swaggertype:"string" format:"uuid"`                         // Phone Location Type ID
	FaxNumber           *string    `json:"fax_number,omitempty" gorm:"-" example:"021911"`                                                       // Fax Number
	Latitude            *float64   `json:"latitude,omitempty" gorm:"-" example:"-1.109209"`                                                      // Latitude
	Longitude           *float64   `json:"longitude,omitempty" gorm:"-" example:"0.012930192"`                                                   // Longitude`
	Email               *string    `json:"email,omitempty" gorm:"-" example:"anonymous@nomail.com" `                                             // Email
	OperationHours      *string    `json:"operation_hours,omitempty" gorm:"-" example:"Monday - Friday (08:00AM - 17:00PM)"`                     // Operation Hours
}

// // BeforeCreate office
// func (a *Office) BeforeCreate(tx *gorm.DB) error {
// 	// calling super class method
// 	err := a.Base.BeforeCreate(tx)

// 	businessEntity := BusinessEntity{}
// 	// create business entity to get an id
// 	tx.Create(&businessEntity)

// 	address := Address{}
// 	address.PostalCode = a.PostalCode
// 	address.BuildingRoom = a.BuildingRoom
// 	address.Latitude = a.Latitude
// 	address.Longitude = a.Longitude
// 	address.AddressLine = a.AddressLine
// 	address.CountryID = a.CountryID
// 	address.CityID = a.CityID
// 	address.StateProvinceID = a.StateProvinceID
// 	address.AddressLine = a.AddressLine

// 	// create address to get an id
// 	tx.Create(&address)

// 	addressType := AddressType{}
// 	r1 := tx.First(&addressType, `address_type_code = ?`, 6)

// 	businessEntityAddress := BusinessEntityAddress{}
// 	businessEntityAddress.AddressID = address.ID
// 	if r1.RowsAffected == 1 {
// 		businessEntityAddress.AddressTypeID = addressType.ID
// 	}
// 	businessEntityAddress.BusinessEntityID = businessEntity.ID
// 	// create business entity address to get an id
// 	tx.Create(&businessEntityAddress)

// 	emailAddressType := EmailAddressType{}
// 	r := tx.First(&emailAddressType, `email_address_type_code = ?`, 2)

// 	businessEntityEmail := BusinessEntityEmail{}
// 	businessEntityEmail.BusinessEntityID = businessEntity.ID
// 	businessEntityEmail.Email = a.Email
// 	if r.RowsAffected == 1 {
// 		businessEntityEmail.EmailAddressTypeID = emailAddressType.ID
// 	}
// 	tx.Create(&businessEntityEmail)

// 	businessEntityPhone := BusinessEntityPhone{}
// 	phoneType := PhoneTechnologyType{}
// 	r2 := tx.First(&phoneType, `phone_technology_type_name = ?`, `Voice`)
// 	if r2.RowsAffected == 1 {
// 		businessEntityPhone.PhoneTechnologyTypeID = phoneType.ID
// 	}
// 	businessEntityPhone.BusinessEntityID = businessEntity.ID
// 	businessEntityPhone.PhoneNumber = a.PhoneNumber
// 	businessEntityPhone.PhoneLocationTypeID = a.PhoneLocationTypeID
// 	tx.Create(&businessEntityPhone)

// 	businessEntityFax := BusinessEntityPhone{}
// 	faxType := PhoneTechnologyType{}
// 	r3 := tx.First(&faxType, `phone_technology_type_name = ?`, `Fax`)
// 	if r3.RowsAffected == 1 {
// 		businessEntityFax.PhoneTechnologyTypeID = faxType.ID
// 	}
// 	businessEntityFax.BusinessEntityID = businessEntity.ID
// 	businessEntityFax.PhoneNumber = a.FaxNumber
// 	tx.Create(&businessEntityFax)

// 	a.BusinessEntity = &businessEntity
// 	a.BusinessEntityID = businessEntity.ID

// 	return err
// }

// // AfterCreate office
// func (a *Office) AfterCreate(tx *gorm.DB) error {

// 	operationSchedule := OperationSchedule{}
// 	operationSchedule.OperationScheduleName = lib.Strptr(uuid.New().String())
// 	operationSchedule.Description = a.OperationHours
// 	tx.Create(&operationSchedule)

// 	officeOperationSchedule := OfficeOperationSchedule{}
// 	officeOperationSchedule.OperationScheduleID = operationSchedule.ID
// 	officeOperationSchedule.OfficeID = a.ID
// 	tx.Create(&officeOperationSchedule)

// 	agentOffice := AgentOffice{}
// 	agentOffice.AgentID = a.AgentID
// 	agentOffice.OfficeID = a.ID
// 	tx.Create(&agentOffice)

// 	// officeTranslation := OfficeTranslation{}
// 	// officeTranslation.OfficeID = a.ID
// 	// officeTranslation.LanguageCode = a.BusinessEntity.Language.LanguageCode
// 	// officeTranslation.OfficeName = a.OfficeName
// 	// tx.Create(&officeTranslation)

// 	return nil
// }

// // BeforeUpdate office
// func (a *Office) BeforeUpdate(tx *gorm.DB) error {
// 	// calling super class method
// 	err := a.Base.BeforeUpdate(tx)
// 	if nil == a.BusinessEntityID {
// 		a.BeforeCreate(tx)
// 		return a.AfterCreate(tx)
// 	}

// 	r0 := tx.First(&BusinessEntity{}, `id = ?`, a.BusinessEntityID)
// 	if r0.RowsAffected < 1 {
// 		a.BeforeCreate(tx)
// 		return a.AfterCreate(tx)
// 	}

// 	businessEntityAddress := BusinessEntityAddress{}
// 	r := tx.Joins(`Address`).First(&businessEntityAddress, `business_entity_id = ?`, a.BusinessEntityID)
// 	if r.RowsAffected == 1 {
// 		if nil != businessEntityAddress.Address {
// 			address := *businessEntityAddress.Address
// 			address.PostalCode = a.PostalCode
// 			address.BuildingRoom = a.BuildingRoom
// 			address.Latitude = a.Latitude
// 			address.Longitude = a.Longitude
// 			address.AddressLine = a.AddressLine
// 			address.CountryID = a.OfficeAPI.CountryID
// 			address.CityID = a.CityID
// 			if a.OfficeAPI.CityID == nil || *a.OfficeAPI.CityID == uuid.Nil {
// 				tx.Model(&address).Updates(map[string]interface{}{
// 					"city_id": nil,
// 				})
// 			}
// 			address.StateProvinceID = a.StateProvinceID
// 			if a.OfficeAPI.StateProvinceID == nil || *a.OfficeAPI.StateProvinceID == uuid.Nil {
// 				tx.Model(&address).Updates(map[string]interface{}{
// 					"state_province_id": nil,
// 				})
// 			}
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

// 	officeOperationSchedule := OfficeOperationSchedule{}
// 	tx.Joins(`OperationSchedule`).First(&officeOperationSchedule, `office_id = ?`, a.ID)

// 	if nil != officeOperationSchedule.ID && nil != officeOperationSchedule.OperationSchedule {
// 		officeOperationSchedule.OperationSchedule.Description = a.OperationHours
// 		osc := officeOperationSchedule.OperationSchedule
// 		tx.Updates(&osc)
// 	} else {
// 		a.AfterCreate(tx)
// 	}

// 	return err
// }
