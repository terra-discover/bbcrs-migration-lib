package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AccreditationType Accreditation Type
type AccreditationType struct {
	basic.Base
	basic.DataOwner
	AccreditationTypeAPI
	AccreditationTypeTranslation *AccreditationTypeTranslation `json:"accreditation_type_translation,omitempty"`
}

// AccreditationTypeAPI Accreditation Type API
type AccreditationTypeAPI struct {
	AccreditationTypeCode *string `json:"accreditation_type_code,omitempty" gorm:"type:varchar(36);not null;index:idx_accreditation_type_code_deleted_at,unique,where:deleted_at is null" example:"iata"`                                            // Accreditation Type Code
	AccreditationTypeName *string `json:"accreditation_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_accreditation_type_name_deleted_at,unique;where:deleted_at is null" example:"International Air Transport Association (IATA)"` // Accreditation Type Name
}

// Seed AccreditationType
func (accreditationType *AccreditationType) Seed() *AccreditationType {
	seed := AccreditationType{
		AccreditationTypeAPI: AccreditationTypeAPI{
			AccreditationTypeCode: lib.Strptr("iata"),
			AccreditationTypeName: lib.Strptr("International Air Transport Association (IATA)"),
		},
	}
	_ = lib.Merge(seed, &accreditationType)
	return accreditationType
}

// AccreditationTypes struct
type AccreditationTypes []AccreditationType

// Seed for seeder
func (accreditationTypes *AccreditationTypes) Seed() *AccreditationTypes {
	code := []string{"iata", "abta", "tids"}
	name := []string{"International Air Transport Association (IATA)", "Association of British Travel Agents (ABTA)", "Travel Industry Designator Service (TIDS)"}
	for i := 0; i < len(code); i++ {
		*accreditationTypes = append(*accreditationTypes, AccreditationType{
			AccreditationTypeAPI: AccreditationTypeAPI{
				AccreditationTypeCode: lib.Strptr(code[i]),
				AccreditationTypeName: lib.Strptr(name[i]),
			},
		})
	}
	return accreditationTypes
}
