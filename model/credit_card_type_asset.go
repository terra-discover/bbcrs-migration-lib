package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CreditCardTypeAsset Model
type CreditCardTypeAsset struct {
	basic.Base
	basic.DataOwner
	CreditCardTypeID      *uuid.UUID             `json:"credit_card_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Credit Card Type Id
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                  // multimedia description
	CreditCardTypeAssetAPI
}

// CreditCardTypeAssetAPI Model
type CreditCardTypeAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
