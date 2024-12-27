package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// AccessType Agent User Type
type AccessType struct {
	basic.Base
	basic.DataOwner
	AccessTypeAPI
	Agent                 *Agent                 `json:"agent,omitempty"`
	UserType              *UserType              `json:"user_type,omitempty"`               // City
	AccessTypeTranslation *AccessTypeTranslation `json:"access_type_translation,omitempty"` // Agent User Type Translation
}

// AccessTypeAPI Agent User Type Api
type AccessTypeAPI struct {
	Code                        *string    `json:"code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" validate:"required,gt=0,lte=36"`
	Name                        *string    `json:"name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" validate:"required,gt=0,lte=256"`
	AgentID                     *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	UserTypeID                  *uuid.UUID `json:"user_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AssignAllDestinationSystems *bool      `json:"assign_all_destination_systems,omitempty" gorm:"type:bool"`
	AssignAllHotels             *bool      `json:"assign_all_hotels,omitempty" gorm:"type:bool"`
	AssignAllVendors            *bool      `json:"assign_all_vendors,omitempty" gorm:"type:bool"`
	AssignAllCorporates         *bool      `json:"assign_all_corporates,omitempty" gorm:"type:bool"`
	AssignAllSubagents          *bool      `json:"assign_all_subagents,omitempty" gorm:"type:bool"`
	AssignAllMembers            *bool      `json:"assign_all_members,omitempty" gorm:"type:bool"`
	IsDefault                   *bool      `json:"is_default,omitempty" gorm:"type:bool"`
}

// AccessTypeDetail Access Type Detail
type AccessTypeDetail struct {
	AccessType
	NumberOfAssignedModules *int `json:"number_of_assigned_modules,omitempty"`
	NumberOfAssignedUsers   *int `json:"number_of_assigned_users,omitempty"`
}

// Seed Agent User Type
func (agentUserType *AccessType) Seed() *AccessType {
	return &AccessType{
		AccessTypeAPI: AccessTypeAPI{
			Code:                        lib.Strptr("code"),
			Name:                        lib.Strptr("name"),
			AgentID:                     lib.UUIDPtr(uuid.New()),
			UserTypeID:                  lib.UUIDPtr(uuid.New()),
			AssignAllDestinationSystems: lib.Boolptr(false),
			AssignAllHotels:             lib.Boolptr(false),
			AssignAllVendors:            lib.Boolptr(false),
			AssignAllCorporates:         lib.Boolptr(false),
			AssignAllSubagents:          lib.Boolptr(false),
			AssignAllMembers:            lib.Boolptr(false),
			IsDefault:                   lib.Boolptr(false),
		},
	}
}
