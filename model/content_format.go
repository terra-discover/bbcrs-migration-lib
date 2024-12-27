package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// ContentFormat struct
type ContentFormat struct {
	basic.Base
	basic.DataOwner
	ContentFormatCode        *int64                    `json:"content_format_code,omitempty" gorm:"type:smallint;uniqueIndex:,where:deleted_at is null;not null;not null"` // Content Format Code
	ContentFormatName        *string                   `json:"content_format_name,omitempty" gorm:"type:varchar(256);not null"`                                            // Content Format Name
	ContentFormatTranslation *ContentFormatTranslation `json:"content_format_translation,omitempty"`                                                                       // Content Format Translation
}
