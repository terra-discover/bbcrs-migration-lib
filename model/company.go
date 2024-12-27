package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// Company struct
type Company struct {
	basic.Base
	basic.DataOwner
	CompanyAPI
	CompanyTranslation *CompanyTranslation `json:"company_translation,omitempty"`
	BusinessEntity     *BusinessEntity     `json:"business_entity,omitempty" swaggerignore:"true"`
}

// CompanyAPI Company API
type CompanyAPI struct {
	CompanyCode      *string    `json:"company_code,omitempty" gorm:"type:varchar(36);not null;index:idx_company_code_deleted_at,unique,where:deleted_at is null"`  // Company Code
	CompanyName      *string    `json:"company_name,omitempty" gorm:"type:varchar(256);not null;index:idx_company_name_deleted_at,unique,where:deleted_at is null"` // Company Name
	BusinessEntityID *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                    // Business Entity ID
	Description      *string    `json:"description,omitempty" gorm:"type:text"`                                                                                     // Description
}

// Seed Company
func (company *Company) Seed() *Company {
	seed := Company{
		CompanyAPI: CompanyAPI{
			CompanyCode:      lib.Strptr("GIAA"),
			CompanyName:      lib.Strptr("Garuda Indonesia (Persero) Tbk"),
			BusinessEntityID: lib.UUIDPtr(uuid.New()),
			Description:      nil,
		},
	}
	_ = lib.Merge(seed, &company)
	return company
}
