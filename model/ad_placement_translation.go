package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdPlacementTranslation Ad Placement Translation
type AdPlacementTranslation struct {
	basic.Base
	basic.DataOwner
	AdPlacementTranslationAPI
	AdPlacementID *uuid.UUID `json:"ad_placement_id,omitempty" gorm:"type:varchar(36);uniqueIndex:ad_placement_translation_unique;not null" swaggertype:"string" format:"uuid"` // Ad Placement id
}

// AdPlacementTranslationAPI Ad Placement Translation API
type AdPlacementTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:ad_placement_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AdPlacementName *string `json:"ad_placement_name,omitempty" gorm:"type:varchar(64)"`                                                              // Ad Placement Name
	Comment         *string `json:"comment,omitempty" gorm:"type:varchar(4000)"`                                                                      // Comment
	Description     *string `json:"description,omitempty" gorm:"type:varchar(4000)"`                                                                  // Description
}
