package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ContentFormatTranslation struct
type ContentFormatTranslation struct {
	basic.Base
	basic.DataOwner
	ContentFormatName *string    `json:"content_format_name,omitempty" gorm:"type:varchar(256)"`                                                                                             // Content Format Name
	LanguageCode      *string    `json:"language_code,omitempty" gorm:"type:varchar(2);index:content_format_unique,unique,where:deleted_at is null" example:"en"`                            // 2 characters language code
	ContentFormatID   *uuid.UUID `json:"content_format_id,omitempty" gorm:"type:varchar(36);index:content_format_unique,unique,where:deleted_at is null" format:"uuid" swaggertype:"string"` // Content Format ID
}
