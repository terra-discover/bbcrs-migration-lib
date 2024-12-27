package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateProcessingFeeCategory Model
type CorporateProcessingFeeCategory struct {
	basic.Base
	basic.DataOwner
	CorporateProcessingFeeCategoryAPI
	Corporate                    *Corporate             `json:"corporate,omitempty"`               // Corporate
	ProcessingFeeCategory        *ProcessingFeeCategory `json:"processing_fee_category,omitempty"` // ProcessingFee Category
	AgentProcessingFeeCategoryID *uuid.UUID             `json:"agent_processing_fee_category_id"`
}

// CorporateProcessingFeeCategoryAPI Model
type CorporateProcessingFeeCategoryAPI struct {
	CorporateID             *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`               // Corporate ID
	ProcessingFeeCategoryID *uuid.UUID `json:"processing_fee_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // ProcessingFee Category ID
}
