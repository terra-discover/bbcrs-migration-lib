package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// JobTitle model
type JobTitle struct {
	basic.Base
	basic.DataOwner
	JobTitleAPI
	JobTitleTranslation *JobTitleTranslation `json:"job_title_translation,omitempty"`
	TravelPolicyClasses []TravelPolicyClass  `json:",omitempty" gorm:"many2many:travel_policy_class_job_title;"`
	CorporateJobTitle   *CorporateJobTitle   `json:"-" swaggerignore:"true"`
}

// JobTitleAPI model
type JobTitleAPI struct {
	JobTitleCode *string `json:"job_title_code,omitempty" gorm:"type:varchar(36);not null"`
	JobTitleName *string `json:"job_title_name,omitempty" gorm:"type:varchar(256);not null"`
}
