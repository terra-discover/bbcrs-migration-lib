package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// PictureCategory struct
type PictureCategory struct {
	basic.Base
	basic.DataOwner
	PictureCategoryCode        *int64                      `json:"picture_category_code,omitempty" gorm:"type:smallint;uniqueIndex:,where:deleted_at is null;not null"` // Picture Category Code
	PictureCategoryName        *string                     `json:"picture_category_name,omitempty" gorm:"type:varchar(256)"`                                            // Picture Category Name
	PictureCategoryTranslation *PictureCategoryTranslation `json:"picture_category_translation,omitempty"`                                                              // Picture Category Translation
}
