package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Campaign Ad Placement
type Campaign struct {
	basic.Base
	basic.DataOwner
	CampaignAPI
}

// CampaignAPI Ad Placement API
type CampaignAPI struct {
	CampaignCode          *string `json:"campaign_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`
	CampaignName          *string `json:"campaign_name,omitempty" gorm:"type:varchar(64);not null;index:,unique,where:deleted_at is null"`
	ApplicationOrder      *int    `json:"application_order,omitempty" gorm:"type:int"`
	AssignAllProductTypes *bool   `json:"assign_all_product_types,omitempty"`
	AssignAllAirlines     *bool   `json:"assign_all_airlines,omitempty"`
	AssignAllCabinTypes   *bool   `json:"assign_all_cabin_types,omitempty"`
	AssignAllHotels       *bool   `json:"assign_all_hotels,omitempty"`
	Comment               *string `json:"comment,omitempty" gorm:"type:text"`
	Description           *string `json:"description,omitempty" gorm:"type:text"`
}
