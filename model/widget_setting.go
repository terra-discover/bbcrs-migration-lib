package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// WidgetSetting struct
type WidgetSetting struct {
	basic.Base
	basic.DataOwner
	WidgetSettingAPI
	Widget *Widget `json:"widget,omitempty"`
}

// WidgetSettingAPI struct
type WidgetSettingAPI struct {
	WidgetID *uuid.UUID `json:"widget_id,omitempty" gorm:"type:varchar(36);not null"`
	Name     *string    `json:"name,omitempty" gorm:"type:varchar(256);not null"`
	Value    *string    `json:"value,omitempty" gorm:"type:text"`
}
