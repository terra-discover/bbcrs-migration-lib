package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ThemeLocationWidget struct
type ThemeLocationWidget struct {
	basic.Base
	basic.DataOwner
	ThemeLocationWidgetAPI
	ThemeLocation *ThemeLocation `json:"theme_location,omitempty"`
	Widget        *Widget        `json:"widget,omitempty"`
}

// ThemeLocationWidgetAPI struct
type ThemeLocationWidgetAPI struct {
	ThemeLocationID *uuid.UUID `json:"theme_location_id,omitempty" gorm:"type:varchar(36);not null"`
	WidgetID        *uuid.UUID `json:"widget_id,omitempty" gorm:"type:varchar(36);not null"`
}
