package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// WidgetType Widget Type
type WidgetType struct {
	basic.Base
	basic.DataOwner
	WidgetTypeAPI
	WidgetTypeTranslation *WidgetTypeTranslation `json:"widget_type_translation,omitempty"`
	SiteType              *SiteType              `json:"site_type" gorm:"foreignKey:SiteTypeID;references:ID"`
}

// WidgetTypeAPI Widget Type API
type WidgetTypeAPI struct {
	WidgetTypeCode *string    `json:"widget_type_code,omitempty" gorm:"type:varchar(256);not null"` // Widget Type Code
	WidgetTypeName *string    `json:"widget_type_name,omitempty" gorm:"type:varchar(256);not null"` // Widget Type Name
	BuildInTypeID  *string    `json:"build_in_type_id,omitempty" gorm:"type:varchar(36);"`          // Build In Type Id
	SiteTypeID     *uuid.UUID `json:"site_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // Site Type Id
	Media          *string    `json:"media,omitempty" gorm:"type:varchar(10);"`                     // Media
	DummyValue     *string    `json:"dummy_value,omitempty" gorm:"type:text"`                       // Dummy Value
	Description    *string    `json:"description,omitempty" gorm:"type:text"`                       // Description
}
