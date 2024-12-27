package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// ModulePackage Module Package
type ModulePackage struct {
	basic.Base
	basic.DataOwner
	ModulePackageAPI
	ModulePackageTranslation *ModulePackageTranslation `json:"module_package_translation,omitempty"` // Module Package Translation
}

// ModulePackageAPI Module Package API
type ModulePackageAPI struct {
	ModulePackageCode *string `json:"module_package_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" example:"help_center"`  // Module package code
	ModulePackageName *string `json:"module_package_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Help Center"` // Module package name
	Description       *string `json:"description,omitempty" gorm:"type:text" example:"Help Center"`                                                                 // Description
}

// Seed module packages
func (p *ModulePackage) Seed() *[]ModulePackage {
	data := []ModulePackage{}

	id, _ := uuid.Parse("a2ad2826-2041-4f2b-a2a0-4dcff747cd8b")
	data = append(data, ModulePackage{
		Base: basic.Base{ID: &id},
		ModulePackageAPI: ModulePackageAPI{
			ModulePackageCode: lib.Strptr("help_center"),
			ModulePackageName: lib.Strptr("Help Center"),
		},
	})

	return &data
}
