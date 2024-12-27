package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// VoucherTypeAsset Model
type VoucherTypeAsset struct {
	basic.Base
	basic.DataOwner
	VoucherTypeID         *uuid.UUID             `json:"voucher_type_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"` // Voucher Type ID
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                              // multimedia description
	VoucherTypeAssetAPI
}

// VoucherTypeAssetAPI Model
type VoucherTypeAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"` // Multimedia Description ID
}
