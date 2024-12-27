package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SegmentCategoryTranslation Segment Category Translation
type SegmentCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	SegmentCategoryTranslationAPI
	SegmentCategoryID *uuid.UUID `json:"segment_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:segment_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Segment Category id
}

// SegmentCategoryTranslationAPI Segment Category Translation API
type SegmentCategoryTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:segment_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	SegmentCategoryName *string `json:"segment_category_name,omitempty" gorm:"type:varchar(256)" example:"All suite"`                                         // Segment Category Name
}
