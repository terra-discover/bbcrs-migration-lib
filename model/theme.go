package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Theme Theme
type Theme struct {
	basic.Base
	basic.DataOwner
	ThemeAPI
	ThemeTranslation *ThemeTranslation `json:"theme_translation,omitempty"` // Theme Translation
	ThemeType        *ThemeType        `json:"theme_type,omitempty"`        // Theme Type
}

// ThemeAPI Theme API
type ThemeAPI struct {
	ThemeCode   *string    `json:"theme_code,omitempty" gorm:"type:varchar(36);not null" example:"bi1"`                 // Theme Code
	ThemeName   *string    `json:"theme_name,omitempty" gorm:"type:varchar(256);not null" example:"background image 1"` // Theme Name
	ThemeTypeID *uuid.UUID `json:"theme_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`  // Theme Type ID
	IsDefault   *bool      `json:"is_default,omitempty" example:"true"`                                                 // Is Default
	Description *string    `json:"description,omitempty" gorm:"type:text" example:"description"`                        // Description
}
