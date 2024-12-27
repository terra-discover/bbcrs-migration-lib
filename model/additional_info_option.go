package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdditionalInfoOption Additional Info Option
type AdditionalInfoOption struct {
	basic.Base
	basic.DataOwner
	AdditionalInfoOptionAPI
	AdditionalInfoOptionTranslation *AdditionalInfoOptionTranslation `json:"additional_info_option_translation,omitempty"` // Additional Info Option Translation
	AdditionalInfo                  *AdditionalInfo                  `json:"additional_info,omitempty"`                    // Additional Info
}

// AdditionalInfoOptionAPI Additional Info Option API
type AdditionalInfoOptionAPI struct {
	AdditionalInfoOptionName *string    `json:"additional_info_option_name,omitempty" gorm:"type:varchar(256);index:idx_additional_info_option__additional_info_id,unique,where:deleted_at is null;not null"`                          // Additional Info Option Name
	AdditionalInfoID         *uuid.UUID `json:"additional_info_id,omitempty" gorm:"type:varchar(36);index:idx_additional_info_option__additional_info_id,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Additional Info ID
	IsDefault                *bool      `json:"is_default,omitempty"`                                                                                                                                                                  // Is Default
	Description              *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                // Description
}
