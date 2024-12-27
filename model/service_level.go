package model

import (
	"strconv"
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceLevel Service Level
type ServiceLevel struct {
	basic.Base
	basic.DataOwner
	ServiceLevelAPI
	ServiceLevelTranslation *ServiceLevelTranslation `json:"service_level_translation,omitempty"` // Service Level Translation
}

// ServiceLevelAPI Service Level API
type ServiceLevelAPI struct {
	ServiceLevelCode *int    `json:"service_level_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null and service_level_code is not null"` // Service Level Code
	ServiceLevelName *string `json:"service_level_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"`                                // Service Level Name
	IsPrimary        *bool   `json:"is_primary,omitempty"`                                                                                                                 // Is Primary
	Comment          *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                   // Comment
	Description      *string `json:"description,omitempty" gorm:"type:text"`                                                                                               // Description
}

// Seed Seed data
func (f *ServiceLevel) Seed() *[]ServiceLevel {
	data := []ServiceLevel{}

	items := []string{
		"1|Shared|Entry Level. Travel consultants are shared across multiple corporates.",
		"2|Dedicated Offsite|Dedicated travel consultants stationed in own office attend a single corporate.",
		"3|Dedicated Onsite|Dedicated travel consultants stationed in client's office attend a single corporate.",
	}

	for i := range items {
		content := strings.Split(items[i], "|")
		strCode := content[0]
		code, _ := strconv.Atoi(strCode)
		name := content[1]
		desc := content[2]
		data = append(data, ServiceLevel{
			ServiceLevelAPI: ServiceLevelAPI{
				ServiceLevelCode: &code,
				ServiceLevelName: &name,
				Description:      &desc,
			},
		})
	}

	return &data
}
