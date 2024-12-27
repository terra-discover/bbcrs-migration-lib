package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Widget Widget
type Widget struct {
	basic.Base
	basic.DataOwner
	WidgetAPI
	WidgetTranslation *WidgetTranslation `json:"widget_translation,omitempty"`
	WidgetType        *WidgetType        `json:"widget_type" gorm:"foreignKey:WidgetTypeID;references:ID"`
}

// WidgetAPI Widget  API
type WidgetAPI struct {
	WidgetCode   *string    `json:"widget_code,omitempty" gorm:"type:varchar(256);not null"` // Widget  Code
	WidgetName   *string    `json:"widget_name,omitempty" gorm:"type:varchar(256);not null"` // Widget  Name
	WidgetTypeID *uuid.UUID `json:"widget_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	Description  *string    `json:"description,omitempty" gorm:"type:text"`
}
