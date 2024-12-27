package model

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// CorporateAirlineFare Corporate Airline Fare
type CorporateAirlineFare struct {
	basic.Base
	basic.DataOwner
	CorporateAirlineFareAPI
	Corporate *Corporate `json:"corporate,omitempty" gorm:"foreignKey:CorporateID;references:ID"` // corporate
	Airline   *Airline   `json:"airline,omitempty" gorm:"foreignKey:AirlineID;references:ID"`     // airline
}

// CorporateAirlineFareAPI Corporate Airline Fare API
type CorporateAirlineFareAPI struct {
	CorporateID        *uuid.UUID       `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`                   // Corporate Id
	AirlineID          *uuid.UUID       `json:"airline_id,omitempty" validate:"required" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Airline Id
	AccountCode        *string          `json:"account_code,omitempty" validate:"required" gorm:"type:varchar(36)"`                                           // Account code
	NegotiatedFareCode *string          `json:"negotiated_fare_code,omitempty" validate:"required" gorm:"type:varchar(36)"`                                   // Negotiated fare code
	EffectiveDate      *strfmt.DateTime `json:"effective_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`                     // Effective Date
	ExpireDate         *strfmt.DateTime `json:"expire_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`                        // Expire Date
	Description        *string          `json:"description,omitempty" gorm:"type:text"`                                                                       // Description
}

func (c *CorporateAirlineFare) GetAllAvailableAirlineFare(tx *gorm.DB, corporateID *uuid.UUID) []CorporateAirlineFare {
	var corporateAirlineFares []CorporateAirlineFare
	today := time.Now().Format("2006-01-02")
	tx.Preload("Airline").Where("corporate_id = ?", corporateID).
		Where("date(effective_date) <= ?", today).
		Where("date(expire_date) >= ?", today).
		Find(&corporateAirlineFares)

	return corporateAirlineFares
}
