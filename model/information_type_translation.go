package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// InformationTypeTranslation struct
type InformationTypeTranslation struct {
	basic.Base
	basic.DataOwner
	InformationTypeName *string    `json:"information_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                               // Information Type Name
	LanguageCode        *string    `json:"language_code,omitempty" gorm:"type:varchar(2);index:information_type_unique,unique,where:deleted_at is null" example:"en"`                              // 2 characters language code
	InformationTypeID   *uuid.UUID `json:"information_type_id,omitempty" gorm:"type:varchar(36);index:information_type_unique,unique,where:deleted_at is null" format:"uuid" swaggertype:"string"` // Information Type ID
}
