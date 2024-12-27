package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PageSetting struct
type PageSetting struct {
	basic.Base
	basic.DataOwner
	PageSettingAPI
	Page *Page `json:"page,omitempty"` // Page
}

// PageSettingAPI struct
type PageSettingAPI struct {
	PageID   *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Page ID
	Name     *string    `json:"name,omitempty"`                                                                        // Name
	Value    *string    `json:"value,omitempty"`                                                                       // Value
	Location *string    `json:"location,omitempty"`                                                                    // Location
}
