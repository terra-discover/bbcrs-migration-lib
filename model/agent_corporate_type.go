package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentCorporateType Model
type AgentCorporateType struct {
	basic.Base
	basic.DataOwner
	AgentCorporateTypeAPI
	CorporateType *CorporateType `json:"corporate_type" gorm:"foreignKey:CorporateTypeID;references:ID"`
	Agent         *Agent         `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentCorporateTypeAPI API
type AgentCorporateTypeAPI struct {
	AgentID         *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;index:idx_agent_corporate_type__corporate_type_id,unique,where:deleted_at is null" format:"uuid"`
	CorporateTypeID *uuid.UUID `json:"corporate_type_id,omitempty" gorm:"type:varchar(36);not null;index:idx_agent_corporate_type__corporate_type_id,unique,where:deleted_at is null" format:"uuid"`
}

// Seed Agent Corporate Type
func (a AgentCorporateType) Seed(agentID uuid.UUID) *[]AgentCorporateType {
	data := []AgentCorporateType{}
	listCorporateType := []string{
		"1b7d2d4b-8934-447f-8812-19d579cdb7b2",
		"c5f84e5f-b75c-45d2-be87-1bdb7c160b76",
		"1cdc048e-3373-4736-81f3-a4d7e421139e",
		"b81d9769-705b-46f7-b625-777655c01394",
	}

	for i := range listCorporateType {
		corporateTypeID := uuid.MustParse(listCorporateType[i])

		data = append(data, AgentCorporateType{
			AgentCorporateTypeAPI: AgentCorporateTypeAPI{
				AgentID:         &agentID,
				CorporateTypeID: &corporateTypeID,
			},
		})
	}

	return &data
}
