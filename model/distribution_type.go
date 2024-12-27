package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// DistributionType Distribution Type
type DistributionType struct {
	basic.Base
	basic.DataOwner
	DistributionTypeAPI
	DistributionTypeTranslation *DistributionTypeTranslation `json:"distribution_type_translation,omitempty"` // Distribution Type Translation
}

// DistributionTypeAPI Distribution Type API
type DistributionTypeAPI struct {
	DistributionTypeCode *int    `json:"distribution_type_code,omitempty" gorm:"type:int;index:,unique,where:deleted_at is null;not null" example:"1"`            // Distribution Type Code
	DistributionTypeName *string `json:"distribution_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Fax"` // Distribution Type Name
}

// Seed data
func (s DistributionType) Seed() *[]DistributionType {
	data := []DistributionType{}
	items := []string{
		"Fax",
		"E-mail",
		"Mail",
		"Courier",
		"Airport collection",
		"City office",
		"Hotel desk",
		"Will call",
		"Express mail",
		"Telephone",
		"XML",
		"FTP",
		"Website",
		"HTTP",
		"Non-XML",
		"Any",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, DistributionType{
			DistributionTypeAPI: DistributionTypeAPI{
				DistributionTypeCode: &code,
				DistributionTypeName: &name,
			},
		})
	}

	return &data
}
