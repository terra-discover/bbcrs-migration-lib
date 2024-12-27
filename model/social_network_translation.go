package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SocialNetworkTranslation Social Network Translation
type SocialNetworkTranslation struct {
	basic.Base
	basic.DataOwner
	SocialNetworkTranslationAPI
	SocialNetworkID *uuid.UUID `json:"social_network_id,omitempty" gorm:"type:varchar(36);uniqueIndex:social_network_translation_unique;not null" swaggertype:"string" format:"uuid"` // Social Network id
}

// SocialNetworkTranslationAPI Social Network Translation API
type SocialNetworkTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:social_network_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	SocialNetworkName *string `json:"social_network_name,omitempty" gorm:"type:varchar(256)" example:"facebook"`                                          // Social Network Name
}
