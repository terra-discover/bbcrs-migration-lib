package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Menu Menu
type Menu struct {
	basic.Base
	basic.DataOwner
	MenuAPI
	MenuTranslation *MenuTranslation `json:"menu_translation,omitempty"`
}

// MenuAPI Menu API
type MenuAPI struct {
	MenuCode    *string `json:"menu_code,omitempty" gorm:"type:varchar(36);not null;index:idx_menu_code_deleted_at,unique,where:deleted_at is null" example:"agent.menu"`  // Menu Code
	MenuName    *string `json:"menu_name,omitempty" gorm:"type:varchar(256);not null;index:idx_menu_name_deleted_at,unique,where:deleted_at is null" example:"Agent Menu"` // Menu Name
	Description *string `json:"description,omitempty" gorm:"type:text"`                                                                                                    // Description
}

// Seed Menu
func (menu *Menu) Seed() *Menu {
	seed := Menu{
		MenuAPI: MenuAPI{
			MenuCode: lib.Strptr("agent.menu"),
			MenuName: lib.Strptr("Agent Menu"),
		},
	}
	_ = lib.Merge(seed, &menu)
	return menu
}

// Menus struct
type Menus []Menu

// Seed for seeder menu in db
func (menus *Menus) Seed() *Menus {
	code := []string{"agent.menu", "agent.user-profile.menu", "corporate.menu", "corporate.user-profile.menu"}
	name := []string{"Agent Menu", "Agent User Profile Menu", "Corporate Menu", "Corporate User Profile Menu"}
	for i := 0; i < len(code); i++ {
		*menus = append(*(menus), Menu{
			MenuAPI: MenuAPI{
				MenuCode: lib.Strptr(code[i]),
				MenuName: lib.Strptr(name[i]),
			},
		})
	}
	return menus
}
