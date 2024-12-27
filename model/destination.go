package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Destination struct is not writable by request body
type Destination struct {
	basic.Base
	basic.DataOwner
	DestinationAPI
	DestinationTranslation  *DestinationTranslation `json:"destination_translation,omitempty"`
	DestinationAssetDesktop *DestinationAsset       `json:"destination_asset_desktop,omitempty"`
	DestinationAssetTablet  *DestinationAsset       `json:"destination_asset_tablet,omitempty"`
	DestinationAssetMobile  *DestinationAsset       `json:"destination_asset_mobile,omitempty"`
	Country                 *Country                `json:"country,omitempty"`
	DestinationCity         *City                   `json:"destination_city,omitempty" gorm:"foreignKey:DestinationCityID"`
}

// DestinationAPI Destination API
type DestinationAPI struct {
	DestinationCode         *string              `json:"destination_code,omitempty" gorm:"type:varchar(36);not null;index:unique_destination__destination_code,unique,where:deleted_at is null" example:"east-jakarta"`
	DestinationName         *string              `json:"destination_name,omitempty" gorm:"type:varchar(256);not null;index:unique_destination__destination_name,unique,where:deleted_at is null" example:"East Jakarta"`
	CountryID               *uuid.UUID           `json:"country_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	DestinationCityID       *uuid.UUID           `json:"destination_city_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	Slug                    *string              `json:"slug,omitempty" gorm:"type:varchar(256)"  example:"https://bayubuanatravel.com/east-jakarta"`
	Description             *string              `json:"description,omitempty" gorm:"type:text"`
	DestinationAssetDesktop *DestinationAssetAPI `json:"destination_asset_desktop,omitempty" gorm:"-"`
	DestinationAssetTablet  *DestinationAssetAPI `json:"destination_asset_tablet,omitempty" gorm:"-"`
	DestinationAssetMobile  *DestinationAssetAPI `json:"destination_asset_mobile,omitempty" gorm:"-"`
}
