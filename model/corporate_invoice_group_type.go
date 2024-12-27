package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// CorporateInvoiceGroupType Corporate Invoice Group Type
type CorporateInvoiceGroupType struct {
	basic.Base
	basic.DataOwner
	CorporateInvoiceGroupTypeAPI
}

// CorporateInvoiceGroupTypeAPI Corporate Invoice Group Type API
type CorporateInvoiceGroupTypeAPI struct {
	InvoiceGroupTypeCode *int    `json:"invoice_group_type_code,omitempty" gorm:"type:smallint;not null"`    // invoice_group_type_code
	InvoiceGroupTypeName *string `json:"invoice_group_type_name,omitempty" gorm:"type:varchar(64);not null"` // invoice_group_type_name
}
