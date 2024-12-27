package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PictureCategoryTranslation struct
type PictureCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	PictureCategoryName *int64     `json:"picture_category_name,omitempty" gorm:"type:varchar(256)"`                                                                                               // Picture Category Name
	LanguageCode        *string    `json:"language_code,omitempty" gorm:"type:varchar(2);index:picture_category_unique,unique,where:deleted_at is null" example:"en"`                              // 2 characters language code
	PictureCategoryID   *uuid.UUID `json:"picture_category_id,omitempty" gorm:"type:varchar(36);index:picture_category_unique,unique,where:deleted_at is null" format:"uuid" swaggertype:"string"` // Picture Category ID
}
