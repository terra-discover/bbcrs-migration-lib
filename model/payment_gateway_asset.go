package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PaymentGatewayAsset Payment Gateway Asset
type PaymentGatewayAsset struct {
	basic.Base
	basic.DataOwner
	PaymentGatewayID      *uuid.UUID             `json:"payment_gateway_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"` // Payment Gateway ID
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                 // multimedia description
	PaymentGatewayAssetAPI
}

// PaymentGatewayAssetAPI Payment Gateway Asset API
type PaymentGatewayAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
