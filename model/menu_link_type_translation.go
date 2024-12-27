package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MenuLinkTypeTranslation Menu Link Type Translation
type MenuLinkTypeTranslation struct {
	basic.Base
	basic.DataOwner
	MenuLinkTypeTranslationAPI
	MenuLinkTypeID *uuid.UUID `json:"menu_link_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:menu_link_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Menu Link Type id
}

// MenuLinkTypeTranslationAPI Menu Link Type Translation API
type MenuLinkTypeTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:menu_link_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MenuLinkTypeName *string `json:"menu_link_type_name,omitempty" gorm:"type:varchar(256)" example:"Submenu"`                                           // Menu Link Type Name
}
