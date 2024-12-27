package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateJobTitle model
type CorporateJobTitle struct {
	basic.Base
	basic.DataOwner
	CorporateJobTitleAPI
	JobTitle *JobTitle `json:"-" swaggerignore:"true"`
}

// CorporateJobTitleAPI model
type CorporateJobTitleAPI struct {
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_job_title,unique,where:deleted_at is null;not null;" swaggertype:"string" format:"uuid"`
	JobTitleID  *uuid.UUID `json:"job_title_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_job_title,unique,where:deleted_at is null;not null;" swaggertype:"string" format:"uuid"`
	IsDefault   *bool      `json:"is_default,omitempty" gorm:"type:bool;default:false" swaggertype:"boolean"`
}
