package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelTranslation Hotel Translation
type HotelTranslation struct {
	basic.Base
	basic.DataOwner
	HotelTranslationAPI
	HotelID *uuid.UUID `json:"hotel_id,omitempty" gorm:"type:varchar(36);uniqueIndex:hotel_translation_unique;not null" swaggertype:"string" format:"uuid"` // Hotel id
}

// HotelTranslationAPI Hotel Translation API
type HotelTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:hotel_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	HotelName      *string `json:"hotel_name,omitempty" gorm:"type:varchar(128)"`                                                             // Hotel Name
	WhenBuilt      *string `json:"when_built,omitempty" gorm:"type:varchar(16)"`                                                              // When Built
	LastRenovation *string `json:"last_renovation,omitempty" gorm:"type:varchar(16)"`                                                         // Last Renovation
	AreaWeather    *string `json:"area_weather,omitempty" gorm:"type:text"`                                                                   // Area Weather
	Comment        *string `json:"comment,omitempty" gorm:"type:text"`                                                                        // Comment
	Description    *string `json:"description,omitempty" gorm:"type:text"`                                                                    // Description
}
