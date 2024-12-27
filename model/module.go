package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// Module Module read only model
type Module struct {
	basic.Base
	basic.DataOwner
	ModuleAPI
	CallOut           *CallOut           `json:"call_out,omitempty" gorm:"foreignKey:CallOutID;references:ID"`
	ModulePackage     *ModulePackage     `json:"module_package,omitempty" gorm:"foreignKey:ModulePackageID;references:ID"`
	ModuleTranslation *ModuleTranslation `json:"module_translation,omitempty"`
}

// ModuleAPI Module API
type ModuleAPI struct {
	ModuleCode       *string          `json:"module_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`
	ModuleName       *string          `json:"module_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"`
	ModulePackageID  *uuid.UUID       `json:"module_package_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	Path             *string          `json:"path,omitempty" gorm:"type:varchar(256)"`
	IsCMS            *bool            `json:"is_cms,omitempty"`
	IsHelp           *bool            `json:"is_help,omitempty"`
	IsSystem         *bool            `json:"is_system,omitempty"`
	CallOutID        *uuid.UUID       `json:"call_out_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	CallOutStartDate *strfmt.DateTime `json:"call_out_start_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	CallOutEndDate   *strfmt.DateTime `json:"call_out_end_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	Description      *string          `json:"description,omitempty" gorm:"type:text"`
}

// Seed modules
func (m *Module) Seed() *[]Module {
	data := []Module{}
	items := []string{
		"help_center_agent_extranet|Bayu Buana Extranet - Help Center",
		"help_center_corporate_extranet|Corporate Client Extranet - Help Center",
		"help_center_member_extranet|Loyalty Member Extranet - Help Center",
		"help_center_nonmember_extranet|B2C/ Retail - Help Center",
		"dashboard_agent_extranet|Bayu Buana Extranet - Dashboard",
		"dashboard_corporate_extranet|Corporate Client Extranet - Dashboard",
		"dashboard_member_extranet|Loyalty Member Extranet - Dashboard",
		"notification_agent_extranet|Bayu Buana Extranet - Alert & Notification",
	}
	id, _ := uuid.Parse("a2ad2826-2041-4f2b-a2a0-4dcff747cd8b")
	for l := range items {
		item := strings.Split(items[l], "|")
		data = append(data, Module{
			ModuleAPI: ModuleAPI{
				ModuleCode:      lib.Strptr(item[0]),
				ModuleName:      lib.Strptr(item[1]),
				ModulePackageID: &id,
			},
		})
	}

	return &data
}

// ModuleAccessCapability combine module and capability
type ModuleAccessCapability struct {
	Module
	AllowedCapabilities int64 `json:"allowed_capabilities"` // allowed capabilities
	Capabilities        struct {
		ID              *uuid.UUID `json:"-" swaggerignore:"true"`
		AllowView       bool       `json:"allow_view"`        // Is view allowed
		AllowCreate     bool       `json:"allow_create"`      // Is create allowed
		AllowEdit       bool       `json:"allow_edit"`        // Is edit allowed
		AllowDelete     bool       `json:"allow_delete"`      // Is delete allowed
		AllowBulkUpdate bool       `json:"allow_bulk_update"` // Is bulk_update allowed
		AllowExport     bool       `json:"allow_export"`      // Is export allowed
		AllowPublish    bool       `json:"allow_publish"`     // Is publish allowed
		AllowActivate   bool       `json:"allow_activate"`    // Is activate allowed
	} `json:"capabilities" gorm:"references:ID;foreignKey:ID"`
}
