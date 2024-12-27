package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelSupplierTranslation Hotel Supplier Translation
type HotelSupplierTranslation struct {
	basic.Base
	basic.DataOwner
	HotelSupplierTranslationAPI
	HotelSupplierID *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);uniqueIndex:hotel_supplier_translation_unique;not null" swaggertype:"string" format:"uuid"` // Hotel Supplier id
}

// HotelSupplierTranslationAPI Hotel Supplier Translation API
type HotelSupplierTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:hotel_supplier_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelSupplierName *string `json:"hotel_supplier_name,omitempty" gorm:"type:varchar(256)" example:"Booking.com"`                                       // Hotel Supplier Name
}
