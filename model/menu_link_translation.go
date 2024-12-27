package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MenuLinkTranslation Menu Link Translation
type MenuLinkTranslation struct {
	basic.Base
	basic.DataOwner
	MenuLinkTranslationAPI
	MenuLinkID *uuid.UUID `json:"menu_link_id,omitempty" gorm:"type:varchar(36);uniqueIndex:menu_link_translation_unique;not null" swaggertype:"string" format:"uuid"` // Menu id
}

// MenuLinkTranslationAPI Menu Link Translation API
type MenuLinkTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:menu_link_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MenuLinkName *string `json:"menu_link_name,omitempty" gorm:"type:varchar(256)" example:"Agent Menu"`                                        // Menu Link Name
	URL          *string `json:"url,omitempty" gorm:"type:varchar(256)" example:"example.com"`                                                  // Menu URL
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                        // Description
}
