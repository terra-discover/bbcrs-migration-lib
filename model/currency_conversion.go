package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CurrencyConversion Currency Conversion
type CurrencyConversion struct {
	basic.Base
	basic.DataOwner
	CurrencyConversionAPI
	FromCurrency       *Currency        `json:"from_currency" gorm:"foreignKey:FromCurrencyID;references:ID"`
	ToCurrency         *Currency        `json:"to_currency" gorm:"foreignKey:ToCurrencyID;references:ID"`
	ConversionRateType *string          `json:"conversion_rate_type,omitempty" example:"C" gorm:"type:varchar(1)"`                    // Conversion Rate Type
	ValidFrom          *strfmt.DateTime `json:"valid_from,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // Valid From
	ValidTo            *strfmt.DateTime `json:"valid_to,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`   // Valid To
}

// CurrencyConversionAPI Currency Conversion API
type CurrencyConversionAPI struct {
	FromCurrencyID *uuid.UUID `json:"from_currency_id,omitempty" swaggertype:"string" gorm:"type:varchar(36);not null" format:"uuid"` // From Currency ID
	ToCurrencyID   *uuid.UUID `json:"to_currency_id,omitempty" swaggertype:"string" gorm:"type:varchar(36);not null" format:"uuid"`   // To Currency ID
	MultiplyRate   *float64   `json:"multiply_rate,omitempty" gorm:"type:decimal(19,8)" example:"1.08"`                               // Multiply Rate
	IsAutomatic    *bool      `json:"is_automatic,omitempty" example:"false"`                                                         // Is Automatic
}

// BeforeCreate currency conversion
// func (a *CurrencyConversion) BeforeCreate(tx *gorm.DB) error {
// 	err := a.Base.BeforeCreate(tx)

// 	oldValue := fmt.Sprintf("%f", *a.MultiplyRate)
// 	newValue := fmt.Sprintf("%f", *a.MultiplyRate)

// 	var resultEventID string
// 	var dataEvent Event
// 	tx.Model(&dataEvent).
// 		Select("id").
// 		Where(`event_code = 17`).Row().Scan(&resultEventID)
// 	eventID, _ := uuid.Parse(resultEventID)

// 	eventLog := EventLog{}
// 	eventLog.EventID = &eventID
// 	eventLog.CurrencyConversionID = a.ID
// 	eventLog.OldValue = &oldValue
// 	eventLog.NewValue = &newValue
// 	eventLog.UserAccountID = a.CreatorID
// 	tx.Create(&eventLog)

// 	return err
// }

// BeforeUpdate currency conversion
// func (a *CurrencyConversion) BeforeUpdate(tx *gorm.DB) error {
// 	err := a.Base.BeforeUpdate(tx)

// 	eventLog := EventLog{}
// 	res := tx.First(&eventLog, `currency_conversion_id`, a.ID)
// 	if res.RowsAffected == 0 {
// 		return a.Base.BeforeCreate(tx)
// 	}

// 	if res.RowsAffected == 1 {
// 		oldValue := fmt.Sprintf("%f", *a.MultiplyRate)
// 		newValue := fmt.Sprintf("%f", *a.MultiplyRate)

// 		eventLog.OldValue = &oldValue
// 		eventLog.NewValue = &newValue
// 		tx.Updates(&eventLog)
// 	}

// 	return err
// }
