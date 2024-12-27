package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateInvoiceBatch Corporate Invoice Batch
type CorporateInvoiceBatch struct {
	basic.Base
	basic.DataOwner
	CorporateInvoiceBatchAPI
}

// CorporateInvoiceBatchAPI Corporate Invoice Batch API
type CorporateInvoiceBatchAPI struct {
	CorporateID      *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`  // corporate_id
	InvoiceBatchName *string    `json:"invoice_batch_name,omitempty" gorm:"type:varchar(64)" example:"invoice"` // invoice_batch_name
	ScheduleID       *uuid.UUID `json:"schedule_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`   // schedule_id
}
