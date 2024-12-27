package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateInvoiceGroup Corporate Invoice Group
type CorporateInvoiceGroup struct {
	basic.Base
	basic.DataOwner
	CorporateInvoiceGroupAPI
	Corporate                 *Corporate                 `json:"corporate,omitempty"`                    // Corporate
	CorporateInvoiceGroupType *CorporateInvoiceGroupType `json:"corporate_invoice_group_type,omitempty"` // Corporate Invoice Group Type
}

// CorporateInvoiceGroupAPI Corporate Invoice Group  API
type CorporateInvoiceGroupAPI struct {
	CorporateID                 *uuid.UUID `json:"corporate_id" gorm:"type:varchar(36);not null;"`                    // Corporate ID
	CorporateInvoiceGroupTypeID *uuid.UUID `json:"corporate_invoice_group_type_id" gorm:"type:varchar(36);not null;"` // Corporate Invoice Group Type ID
}
