package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// SocialNetwork Social Network
type SocialNetwork struct {
	basic.Base
	basic.DataOwner
	SocialNetworkAPI
	SocialNetworkTranslation *SocialNetworkTranslation `json:"social_network_translation,omitempty"`
	SocialNetworkAsset       *SocialNetworkAsset       `json:"social_network_asset,omitempty"` // Social Network Asset
}

// SocialNetworkAPI Social Network API
type SocialNetworkAPI struct {
	SocialNetworkCode  *string                `json:"social_network_code,omitempty" gorm:"type:varchar(2);index:,unique,where:deleted_at is null;not null" example:"fb"`         // Social Network Code
	SocialNetworkName  *string                `json:"social_network_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"facebook"` // Social Network Name
	SocialNetworkAsset *SocialNetworkAssetAPI `json:"social_network_asset,omitempty" gorm:"-"`                                                                                   // Social Network Asset
}
