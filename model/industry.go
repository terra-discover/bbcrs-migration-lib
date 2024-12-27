package model

import (
	"strconv"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Industry struct
type Industry struct {
	basic.Base
	basic.DataOwner
	IndustryAPI
	IndustryTranslation *IndustryTranslation `json:"industry_translation,omitempty"`
}

// IndustryAPI Industry API
type IndustryAPI struct {
	IndustryCode *string `json:"industry_code,omitempty" gorm:"type:varchar(36);not null;index:idx_industry_code_deleted_at,unique,where:deleted_at is null"`  // Industry Code
	IndustryName *string `json:"industry_name,omitempty" gorm:"type:varchar(256);not null;index:idx_industry_name_deleted_at,unique,where:deleted_at is null"` // Industry Name
}

// Seed Industry
func (industry *Industry) Seed() *Industry {
	seed := Industry{
		IndustryAPI: IndustryAPI{
			IndustryCode: lib.Strptr("1"),
			IndustryName: lib.Strptr("Agriculture"),
		},
	}
	_ = lib.Merge(seed, &industry)
	return industry
}

// Industries struct
type Industries []Industry

// Seed for seeder industry in db
func (industries *Industries) Seed() *Industries {
	name := []string{"Agriculture", "Chemical", "Oil and Gas", "Manufacturing", "Technology"}
	for i := 0; i < len(name); i++ {
		*industries = append(*(industries), Industry{IndustryAPI: IndustryAPI{
			IndustryCode: lib.Strptr(strconv.Itoa(i + 1)),
			IndustryName: lib.Strptr(name[i]),
		}})
	}
	return industries
}
