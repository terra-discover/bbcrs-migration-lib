package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViewershipAccessCode Viewership Access Code
type ViewershipAccessCode struct {
	basic.Base
	basic.DataOwner
	ViewershipAccessCodeAPI
	Viewership *Viewership `json:"viewership" gorm:"foreignKey:ViewershipID;references:ID"`
}

// ViewershipAccessCodeAPI Viewership Access Code Api
type ViewershipAccessCodeAPI struct {
	ViewershipID *uuid.UUID `json:"viewership_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	AccessCode   *string    `json:"reference_code,omitempty" gorm:"type:varchar(36)"`
}
