package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OperationScheduleTranslation Operation Schedule Translation
type OperationScheduleTranslation struct {
	basic.Base
	basic.DataOwner
	OperationScheduleTranslationAPI
	OperationScheduleID *uuid.UUID `json:"operation_schedule_id,omitempty" gorm:"type:varchar(36);uniqueIndex:operation_schedule_translation_unique;not null" swaggertype:"string" format:"uuid"` // Operation Schedule id
}

// OperationScheduleTranslationAPI Operation Schedule Translation API
type OperationScheduleTranslationAPI struct {
	LanguageCode          *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:operation_schedule_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	OperationScheduleName *string `json:"operation_schedule_name,omitempty" gorm:"type:varchar(256)" example:"cleaning"`                                          // Operation Schedule Name
	Description           *string `json:"description,omitempty" gorm:"type:text"`                                                                                 // Description
}
