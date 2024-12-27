package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViewershipTranslation Viewership Translation
type ViewershipTranslation struct {
	basic.Base
	basic.DataOwner
	ViewershipTranslationAPI
	Viewership *Viewership `json:"viewership" gorm:"foreignKey:ViewershipID;references:ID"`
	Language   *Language   `json:"language" gorm:"foreignKey:LanguageID;references:ID"`
}

// ViewershipTranslationAPI Viewership Translation Api
type ViewershipTranslationAPI struct {
	ViewershipID   *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	LanguageID     *uuid.UUID `json:"language_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	ViewershipName *string    `json:"viewership_name,omitempty" gorm:"type:varchar(64);"`
	Comment        *string    `json:"comment,omitempty" gorm:"type:text;"`     // Comment
	Description    *string    `json:"description,omitempty" gorm:"type:text;"` // Description
}
