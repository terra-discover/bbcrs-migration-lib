package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BenefitAsset struct is not writable by request body
type BenefitAsset struct {
	basic.Base
	basic.DataOwner
	BenefitAssetAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                    // multimedia description
	BenefitID             *uuid.UUID             `json:"benefit_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // benefit id
}

// BenefitAssetAPI struct is writable by request body
type BenefitAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // multimedia description id
}
