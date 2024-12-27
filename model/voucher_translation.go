package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// VoucherTranslation Voucher Translation
type VoucherTranslation struct {
	basic.Base
	basic.DataOwner
	VoucherTranslationAPI
	VoucherID *uuid.UUID `json:"voucher_id,omitempty" gorm:"type:varchar(36);uniqueIndex:voucher_translation_unique;not null" swaggertype:"string" format:"uuid"` // Voucher id
}

// VoucherTranslationAPI Voucher Translation API
type VoucherTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:voucher_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	VoucherName  *string `json:"voucher_name,omitempty" example:"voucher code" gorm:"type:varchar(64)"`                                       // Voucher Name
	Comment      *string `json:"comment,omitempty" example:"comment" gorm:"type:text"`                                                        // Comment
	Description  *string `json:"description,omitempty" example:"descriptions" gorm:"type:text"`                                               // Description
}
