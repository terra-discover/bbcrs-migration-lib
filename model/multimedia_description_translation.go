package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MultimediaDescriptionTranslation struct
type MultimediaDescriptionTranslation struct {
	basic.Base
	basic.DataOwner
	Title                   *string        `json:"title,omitempty" gorm:"type:varchar(256)"`                                                                                                                           // Title
	FileName                *string        `json:"file_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                       // File Name
	FileSize                *float64       `json:"file_size,omitempty" format:"float" swaggertype:"number"`                                                                                                            // File Size
	DimensionWidth          *float64       `json:"dimension_width,omitempty" format:"float" swaggertype:"number"`                                                                                                      // Dimension Width
	DimensionHeight         *float64       `json:"dimension_height,omitempty" format:"float" swaggertype:"number"`                                                                                                     // Dimension Height
	OriginalFileName        *string        `json:"original_file_name,omitempty" gorm:"type:varchar(256)"`                                                                                                              // Original File Name
	URL                     *string        `json:"url,omitempty" gorm:"type:varchar(256)"`                                                                                                                             // Url
	Description             *string        `json:"description,omitempty" gorm:"type:text"`                                                                                                                             // Description
	ContentFormatID         *uuid.UUID     `json:"content_format_id,omitempty" format:"uuid" swaggertype:"string" gorm:"type:varchar(36)"`                                                                             // Content Format ID
	ContentFormat           *ContentFormat `json:"content_format,omitempty"`                                                                                                                                           // Content Format
	LanguageCode            *string        `json:"language_code,omitempty" gorm:"type:varchar(2);index:multimedia_description_unique,unique,where:deleted_at is null" example:"en"`                                    // 2 characters language code
	MultimediaDescriptionID *uuid.UUID     `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);index:multimedia_description_unique,unique,where:deleted_at is null" format:"uuid" swaggertype:"string"` // Multimedia Description Id
}
