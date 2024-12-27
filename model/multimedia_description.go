package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MultimediaDescription struct
type MultimediaDescription struct {
	basic.Base
	basic.DataOwner
	Title               *string            `json:"title,omitempty" gorm:"type:varchar(256)"`                                                       // Title
	FileName            *string            `json:"file_name,omitempty" gorm:"type:varchar(256)"`                                                   // File Name
	FileSize            *float64           `json:"file_size,omitempty" format:"float" swaggertype:"number"`                                        // File Size
	DimensionWidth      *float64           `json:"dimension_width,omitempty" format:"float" swaggertype:"number"`                                  // Dimension Width
	DimensionHeight     *float64           `json:"dimension_height,omitempty" format:"float" swaggertype:"number"`                                 // Dimension Height
	OriginalFileName    *string            `json:"original_file_name,omitempty" gorm:"type:varchar(256)"`                                          // Original File Name
	URL                 *string            `json:"url,omitempty" gorm:"type:varchar(256)" example:"protocol://domain.ltd/path/to/image.extension"` // Url
	Description         *string            `json:"description,omitempty" gorm:"type:text"`                                                         // Description
	InformationTypeID   *uuid.UUID         `json:"-" swaggerignore:"true" gorm:"type:varchar(36)"`                                                 // Information Type ID
	PictureCategoryID   *uuid.UUID         `json:"-" swaggerignore:"true" gorm:"type:varchar(36)"`                                                 // Picture Category ID
	ContentFormatID     *uuid.UUID         `json:"-" swaggerignore:"true" gorm:"type:varchar(36)"`                                                 // Content Format ID
	DimensionCategoryID *uuid.UUID         `json:"-" swaggerignore:"true" gorm:"type:varchar(36)"`                                                 // Dimension Category ID
	UnitOfMeasureID     *uuid.UUID         `json:"-" swaggerignore:"true" gorm:"type:varchar(36)"`                                                 // Unit Of Measure ID
	InformationType     *InformationType   `json:"information_type,omitempty" swaggerignore:"true"`                                                // Information Type
	PictureCategory     *PictureCategory   `json:"picture_category,omitempty" swaggerignore:"true"`                                                // Picture Category
	ContentFormat       *ContentFormat     `json:"content_format,omitempty" swaggerignore:"true"`                                                  // Content Format
	DimensionCategory   *DimensionCategory `json:"dimension_category,omitempty" swaggerignore:"true"`                                              // Dimension Category
	UnitOfMeasure       *UnitOfMeasure     `json:"unit_of_measure,omitempty" swaggerignore:"true"`                                                 // Unit Of Measure
	CorporateAsset      *CorporateAsset    `json:"corporate_asset,omitempty" swaggerignore:"true"`
}
