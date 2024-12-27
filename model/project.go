package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Project Project
type Project struct {
	basic.Base
	basic.DataOwner
	ProjectAPI
	ProjectTranslation *ProjectTranslation `json:"project_translation,omitempty"` // Project Translation
}

// ProjectAPI Project API
type ProjectAPI struct {
	ProjectCode *string `json:"project_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`  // Project Code
	ProjectName *string `json:"project_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Project Name
}
