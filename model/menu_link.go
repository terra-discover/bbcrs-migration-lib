package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// MenuLink struct
type MenuLink struct {
	basic.Base
	basic.DataOwner
	MenuLinkAPI
	MenuLinkTranslation *MenuLinkTranslation `json:"menu_link_translation,omitempty"`            // MenuLinkTranslation
	MenuLinkType        *MenuLinkType        `json:"menu_link_type,omitempty"`                   // MenuLinkType
	Menu                *Menu                `json:"menu,omitempty"`                             // Menu
	Module              *Module              `json:"module,omitempty"`                           // Module
	Page                *Page                `json:"page,omitempty"`                             // Page
	MenuLinkAsset       *MenuLinkAsset       `json:"menu_link_asset,omitempty"`                  // MenuLinkAsset
	ParentLink          *MenuLink            `json:"parent_link,omitempty" swaggerignore:"true"` // Parent Link
}

// MenuLinkAPI struct
type MenuLinkAPI struct {
	MenuLinkCode    *string           `json:"menu_link_code" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`  // Menu Link Code
	MenuLinkName    *string           `json:"menu_link_name" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Menu Link Name
	MenuLinkTypeID  *uuid.UUID        `json:"menu_link_type_id" gorm:"type:varchar(36);not null"`                                      // MenuLinkType ID
	MenuID          *uuid.UUID        `json:"menu_id,omitempty" gorm:"type:varchar(36);not null"`                                      // Menu ID
	ParentLinkID    *uuid.UUID        `json:"parent_link_id" gorm:"type:varchar(36)"`                                                  // ParentLink ID
	ModuleID        *uuid.UUID        `json:"module_id" gorm:"type:varchar(36)"`                                                       // Module ID
	PageID          *uuid.UUID        `json:"page_id" gorm:"type:varchar(36)"`                                                         // Page ID
	URL             *string           `json:"url" gorm:"type:varchar(256)"`                                                            // URL
	IsExpanded      *bool             `json:"is_expanded"`                                                                             // Is Expanded
	Target          *string           `json:"target" gorm:"type:varchar(36)"`                                                          // Target
	IsAuthenticated *bool             `json:"is_authenticated"`                                                                        // Is Authenticated
	IsLogOut        *bool             `json:"is_log_out"`                                                                              // Is Log Out
	Description     *string           `json:"description" gorm:"type:text"`                                                            // Description
	MenuLinkAsset   *MenuLinkAssetAPI `json:"menu_link_asset,omitempty" gorm:"-"`
}

// Seed data
func (menuLink *MenuLink) Seed() *MenuLink {
	menuLinkCode := "MNL"
	menuLinkName := "menuLinkName"
	url := "url"
	isExpanded := false
	target := "target"
	isAuthenticated := false
	isLogOut := false
	isDescription := "description"

	initial := MenuLink{
		MenuLinkAPI: MenuLinkAPI{
			MenuLinkCode:    &menuLinkCode,
			MenuLinkName:    &menuLinkName,
			MenuLinkTypeID:  lib.UUIDPtr(uuid.New()),
			MenuID:          lib.UUIDPtr(uuid.New()),
			ParentLinkID:    lib.UUIDPtr(uuid.New()),
			ModuleID:        lib.UUIDPtr(uuid.New()),
			PageID:          lib.UUIDPtr(uuid.New()),
			URL:             &url,
			IsExpanded:      &isExpanded,
			Target:          &target,
			IsAuthenticated: &isAuthenticated,
			IsLogOut:        &isLogOut,
			Description:     &isDescription,
		},
	}
	lib.Merge(initial, &menuLink)
	return menuLink
}
