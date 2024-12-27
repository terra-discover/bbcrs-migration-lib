package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UnitOfMeasureTranslation Unit Of Measure Translation
type UnitOfMeasureTranslation struct {
	basic.Base
	basic.DataOwner
	UnitOfMeasureTranslationAPI
	UnitOfMeasureID *uuid.UUID `json:"unit_of_measure_id,omitempty" gorm:"type:varchar(36);uniqueIndex:unit_of_measure_translation_unique;not null" swaggertype:"string" format:"uuid"` // Unit Of Measure id
}

// UnitOfMeasureTranslationAPI Unit Of Measure Translation API
type UnitOfMeasureTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:unit_of_measure_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	UnitOfMeasureName   *string `json:"unit_of_measure_name,omitempty" gorm:"type:varchar(256)"`                                                             // Unit Of Measure Name
	UnitOfMeasureSymbol *string `json:"unit_of_measure_symbol,omitempty" gorm:"type:varchar(8)"`                                                             // Unit Of Measure Symbol
	QuantityName        *string `json:"quantity_name,omitempty" gorm:"type:varchar(256)"`                                                                    // Quantity Name
}
