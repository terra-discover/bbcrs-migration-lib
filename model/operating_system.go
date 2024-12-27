package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OperatingSystem Operating System
type OperatingSystem struct {
	basic.Base
	basic.DataOwner
	OperatingSystemAPI
	OperatingSystemTranslation *OperatingSystemTranslation `json:"operating_system_translation,omitempty"`
	OperatingSystemAsset       *OperatingSystemAsset       `json:"operating_system_asset,omitempty"` // OperatingSystemAsset
}

// OperatingSystemAPI Operating System API
type OperatingSystemAPI struct {
	OperatingSystemCode  *string                  `json:"operating_system_code,omitempty" gorm:"type:varchar(36);not null;" example:"windows"`                                                  // Operating System Code
	OperatingSystemName  *string                  `json:"operating_system_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Microsoft Windows"` // Operating System Name
	OperatingSystemAsset *OperatingSystemAssetAPI `json:"operating_system_asset,omitempty" gorm:"-"`                                                                                            // OperatingSystemAsset
}

// Seed OperatingSystem
func (operatingSystem *OperatingSystem) Seed() *OperatingSystem {
	seed := OperatingSystem{
		OperatingSystemAPI: OperatingSystemAPI{
			OperatingSystemCode: lib.Strptr("windows"),
			OperatingSystemName: lib.Strptr("Microsoft Windows"),
		},
	}
	_ = lib.Merge(seed, &operatingSystem)
	return operatingSystem
}

// OperatingSystems migrator
type OperatingSystems []OperatingSystem

// Seed Data
func (operatingSystems *OperatingSystems) Seed() *OperatingSystems {
	code := []string{"windows", "mac", "linux", "android", "ios"}
	name := []string{"Microsoft Windows", "Apple MacOS", "Linux", "Android", "Apple iOS"}
	for i := 0; i < len(name); i++ {
		seed := OperatingSystem{
			OperatingSystemAPI: OperatingSystemAPI{
				OperatingSystemCode: lib.Strptr(code[i]),
				OperatingSystemName: lib.Strptr(name[i]),
			},
		}
		*operatingSystems = append(*operatingSystems, seed)
	}
	return operatingSystems
}
