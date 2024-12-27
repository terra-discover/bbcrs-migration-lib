package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateAdditionalInfo Model
type CorporateAdditionalInfo struct {
	basic.Base
	basic.DataOwner
	CorporateAdditionalInfoAPI
	Corporate      *Corporate      `json:"corporate,omitempty"`
	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty"`
}

// CorporateAdditionalInfoAPI Model
type CorporateAdditionalInfoAPI struct {
	CorporateID      *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null;index:idx_corporate_additional_info__corporate_id,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`       // Corporate ID
	AdditionalInfoID *uuid.UUID `json:"additional_info_id,omitempty" gorm:"type:varchar(36);not null;index:idx_corporate_additional_info__corporate_id,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"` // AdditionalInfo ID
}
