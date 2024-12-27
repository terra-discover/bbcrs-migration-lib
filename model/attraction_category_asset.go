package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AttractionCategoryAsset struct
type AttractionCategoryAsset struct {
	basic.Base
	basic.DataOwner
	AttractionCategoryID  *uuid.UUID             `json:"attraction_category_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Attraction Category
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                                                            // multimedia description
	AttractionCategoryAssetAPI
}

// AttractionCategoryAssetAPI for handle request body
type AttractionCategoryAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
