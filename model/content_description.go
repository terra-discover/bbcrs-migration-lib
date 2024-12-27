package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ContentDescription Content Description
type ContentDescription struct {
	basic.Base
	basic.DataOwner
	ContentDescriptionAPI
	ContentDescriptionTranslation *ContentDescriptionTranslation `json:"content_description_translation,omitempty"` // Content Description Translation
	InformationType               *InformationType               `json:"information_type,omitempty"`                // information type
	AdditionalDetailType          *AdditionalDetailType          `json:"additional_detail_type,omitempty"`          // Additional Detail Type
}

// ContentDescriptionAPI Content Description API
type ContentDescriptionAPI struct {
	InformationTypeID      *uuid.UUID `json:"information_type_id,omitempty" swaggertype:"string" format:"uuid"`       // Information Type ID
	AdditionalDetailTypeID *uuid.UUID `json:"additional_detail_type_id,omitempty" swaggertype:"string" format:"uuid"` // Additional Detail Type ID
	Title                  *string    `json:"title,omitempty" gorm:"type:varchar(256)"`                               // Title
	URL                    *string    `json:"url,omitempty" gorm:"type:varchar(256)"`                                 // Url
	Description            *string    `json:"description,omitempty" gorm:"type:text"`                                 // Description
}
