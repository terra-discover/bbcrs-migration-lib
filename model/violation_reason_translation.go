package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ViolationReasonTranslation Violation Reason Translation
type ViolationReasonTranslation struct {
	basic.Base
	basic.DataOwner
	ViolationReasonTranslationAPI
	ViolationReason *ViolationReason `json:"violation_reason,omitempty" gorm:"foreignKey:ViolationReasonID;references:ID"` // violation reason
}

// ViolationReasonTranslationAPI Violation Reason Translation API
type ViolationReasonTranslationAPI struct {
	ViolationReasonID   *uuid.UUID `json:"violation_reason_id,omitempty" gorm:"type:varchar(36);uniqueIndex:violation_reason_translation_unique;not null" swaggertype:"string" swaggerignore:"true" format:"uuid"` // Violation Reason id
	LanguageCode        *string    `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:violation_reason_translation_unique;not null" example:"en"`                                                   // Language code example: en, id, cn etc...
	ViolationReasonName *string    `json:"violation_reason_name,omitempty" gorm:"type:varchar(256)"`                                                                                                               // Violation Reason Name
	Description         *string    `json:"description,omitempty" gorm:"type:text" example:"deskripsi"`                                                                                                             // Description
}
