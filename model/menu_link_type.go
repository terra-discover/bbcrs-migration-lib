package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MenuLinkType Menu Link Type
type MenuLinkType struct {
	basic.Base
	basic.DataOwner
	MenuLinkTypeAPI
	MenuLinkTypeTranslation *MenuLinkTypeTranslation `json:"menu_link_type_translation,omitempty"`
}

// MenuLinkTypeAPI Menu Link Type API
type MenuLinkTypeAPI struct {
	MenuLinkTypeCode *string `json:"menu_link_type_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" example:"submenu"`  // Menu Link Type Code
	MenuLinkTypeName *string `json:"menu_link_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Submenu"` // Menu Link Type Name
	IsDefault        *bool   `json:"is_default,omitempty"`                                                                                                     // Is Default
}

// Seed MenuLinkType
func (menuLinkType *MenuLinkType) Seed() *MenuLinkType {
	seed := MenuLinkType{
		MenuLinkTypeAPI: MenuLinkTypeAPI{
			MenuLinkTypeCode: lib.Strptr("submenu"),
			MenuLinkTypeName: lib.Strptr("Submenu"),
		},
	}
	_ = lib.Merge(seed, &menuLinkType)
	return menuLinkType
}

// MenuLinkTypes migrator
type MenuLinkTypes []MenuLinkType

// Seed data
func (menuLinkTypes *MenuLinkTypes) Seed() *MenuLinkTypes {
	code := []string{"submenu", "menu-section", "menu-item"}
	name := []string{"Submenu", "Menu Section", "Menu Item"}
	for i := 0; i < len(name); i++ {
		seed := MenuLinkType{
			MenuLinkTypeAPI: MenuLinkTypeAPI{
				MenuLinkTypeCode: lib.Strptr(code[i]),
				MenuLinkTypeName: lib.Strptr(name[i]),
			},
		}
		*menuLinkTypes = append(*menuLinkTypes, seed)
	}
	return menuLinkTypes
}
