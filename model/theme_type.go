package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ThemeType Theme Type
type ThemeType struct {
	basic.Base
	basic.DataOwner
	ThemeTypeAPI
	ThemeTypeTranslation *ThemeTypeTranslation `json:"theme_type_translation,omitempty"` // Theme Type Translation
	SiteType             *SiteType             `json:"site_type" gorm:"foreignKey:SiteTypeID;references:ID"`
}

// ThemeTypeAPI Theme Type API
type ThemeTypeAPI struct {
	ThemeTypeCode *string    `json:"theme_type_code,omitempty" gorm:"type:varchar(36);uniqueIndex:idx_theme_type_code_deleted_at,where:deleted_at is null;not null" example:"be"` // Theme Type Code
	ThemeTypeName *string    `json:"theme_type_name,omitempty" gorm:"type:varchar(256)" example:"backend"`                                                                        // Theme Type Name
	BuiltInTypeID *uuid.UUID `json:"built_in_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                       // Built In Type ID
	SiteTypeID    *uuid.UUID `json:"site_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                           // Site Type ID
	Description   *string    `json:"description,omitempty" gorm:"type:text" example:"description"`                                                                                // Description
}
