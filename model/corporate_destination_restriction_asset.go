package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateDestinationRestrictionAsset Corporate Destination Restriction Asset
type CorporateDestinationRestrictionAsset struct {
	basic.Base
	basic.DataOwner
	CorporateDestinationRestrictionID *uuid.UUID `json:"corporate_destination_restriction_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_destination_restriction_asset,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	CorporateDestinationRestrictionAssetAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
}

// CorporateDestinationRestrictionAssetAPI Corporate Destination Restriction Asset Api
type CorporateDestinationRestrictionAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_destination_restriction_asset,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
}
