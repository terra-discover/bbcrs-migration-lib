package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// VoucherTypeTranslation Voucher Type Translation
type VoucherTypeTranslation struct {
	basic.Base
	basic.DataOwner
	VoucherTypeTranslationAPI
	VoucherTypeID *uuid.UUID `json:"voucher_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:voucher_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Voucher Type id
}

// VoucherTypeTranslationAPI Voucher Type Translation API
type VoucherTypeTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:voucher_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	VoucherTypeName *string `json:"voucher_type_name,omitempty" example:"voucher code" gorm:"type:varchar(256)"`                                      // Voucher Type Name
	Comment         *string `json:"comment,omitempty" example:"comments" gorm:"type:text"`                                                            // Comment
	Description     *string `json:"description,omitempty" example:"descriptions" gorm:"type:text"`                                                    // Description
}
