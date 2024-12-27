package model

import (
	"time"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// Agent Agent
type Agent struct {
	basic.Base
	basic.DataOwner
	AgentAPI
	BusinessEntity   *BusinessEntity   `json:"business_entity,omitempty"` // City
	AgentTranslation *AgentTranslation `json:"agent_translation,omitempty"`
	AgentAsset       *AgentAsset       `json:"agent_asset,omitempty"`
}

// AgentAPI Agent API
type AgentAPI struct {
	AgentCode        *string        `json:"agent_code,omitempty" gorm:"type:varchar(36);not null;index:idx_agent_code_deleted_at,unique,where:deleted_at is null" validate:"required,gt=0,lte=36"`   // Agent Code
	AgentName        *string        `json:"agent_name,omitempty" gorm:"type:varchar(256);not null;index:idx_agent_name_deleted_at,unique,where:deleted_at is null" validate:"required,gt=0,lte=256"` // Agent Name
	ProfileStatus    *int           `json:"profile_status,omitempty" gorm:"type:smallint;" example:"1"`                                                                                              // Profile Status                                                                                                                                              // Is Default
	BusinessEntityID *uuid.UUID     `json:"business_entity_id,omitempty" swaggertype:"string" format:"uuid"`
	Comment          *string        `json:"comment,omitempty" gorm:"type:text"`
	Description      *string        `json:"description,omitempty" gorm:"type:text"`
	AgentAsset       *AgentAssetAPI `json:"agent_asset,omitempty" gorm:"-"`
}

// Seed Agent
func (a *Agent) Seed(agentID uuid.UUID) *[]Agent {
	now := strfmt.DateTime(time.Now().UTC())

	data := []Agent{}

	agentName := "Bayu Buana"
	if lib.IsEmptyUUID(agentID) {
		agentID = uuid.New()
	}
	agentIDString := agentID.String()
	initial := Agent{}
	initial.ID = &agentID
	initial.AgentCode = &agentIDString
	initial.AgentName = &agentName
	initial.ProfileStatus = lib.Intptr(1)
	initial.BusinessEntityID = &agentID
	initial.CreatedAt = &now
	data = append(data, initial)

	return &data
}
