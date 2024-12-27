package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ThemeLocationMenu struct
type ThemeLocationMenu struct {
	basic.Base
	basic.DataOwner
	ThemeLocationMenuAPI
}

// ThemeLocationMenuAPI ThemeLocationMenu API
type ThemeLocationMenuAPI struct {
	ThemeLocationID *uuid.UUID `json:"theme_location_id,omitempty" gorm:"type:varchar(36)"` // Theme Location ID
	MenuID          *uuid.UUID `json:"menu_id,omitempty" gorm:"type:varchar(36)"`           // Menu ID
}
