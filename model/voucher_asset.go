package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// VoucherAsset Model
type VoucherAsset struct {
	basic.Base
	basic.DataOwner
	VoucherID             *uuid.UUID             `json:"voucher_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"` // Voucher Type ID
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                         // multimedia description
	VoucherAssetAPI
}

// VoucherAssetAPI Model
type VoucherAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"` // Multimedia Description ID
}
