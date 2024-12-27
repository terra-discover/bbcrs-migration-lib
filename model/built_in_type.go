package model

import (
	"strings"

	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BuiltInType Built In Type
type BuiltInType struct {
	basic.Base
	basic.DataOwner
	BuiltInTypeAPI
	BuiltInTypeTranslation *BuiltInTypeTranslation `json:"built_in_type_translation,omitempty"` // Built In Type Translation
}

// BuiltInTypeAPI Built In Type API
type BuiltInTypeAPI struct {
	BuiltInTypeCode *string `json:"built_in_type_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" example:"agent"`     // Built In Type Code
	BuiltInTypeName *string `json:"built_in_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Internal"` // Built In Type Name
}

// Seed data
func (s BuiltInType) Seed() *[]BuiltInType {
	data := []BuiltInType{}
	items := []string{
		"847f6b27-7191-499a-8455-d26fecd0e912|agent|Internal",
		"dd3a7ce8-2b5a-468b-8ed7-701df32ed3f2|vendor|Vendor",
		"ee14f1e2-9fba-4dcf-b093-4001bfca1afa|subagent|Sub Agent",
		"a29e91e4-a1b6-4612-9ce6-25c4bc96a745|corporate|Corporate Client",
		"0345ad5e-ab02-4d2b-b069-a05ae69ab9ae|retail|Retail",
		"5cddf574-b6db-492b-8288-8a04386436f2|member|Member",
		"bedb5f42-6d8a-4026-a1ac-ca8a6704f469|nonmember|Non-Member Customer",
	}

	for i := range items {
		values := strings.Split(items[i], "|")
		var id uuid.UUID
		id, _ = uuid.Parse(values[0])
		var code string = values[1]
		var name string = values[2]
		data = append(data, BuiltInType{
			Base: basic.Base{
				ID: &id,
			},
			BuiltInTypeAPI: BuiltInTypeAPI{
				BuiltInTypeCode: &code,
				BuiltInTypeName: &name,
			},
		})
	}

	return &data
}

type BuiltInTypeCode string

const (
	AgentBuiltInTypeCode     BuiltInTypeCode = "agent"
	VendorBuiltInTypeCode    BuiltInTypeCode = "vendor"
	SubAgentBuiltInTypeCode  BuiltInTypeCode = "subagent"
	CorporateBuiltInTypeCode BuiltInTypeCode = "corporate"
	RetailBuiltInTypeCode    BuiltInTypeCode = "retail"
	MemberBuiltInTypeCode    BuiltInTypeCode = "member"
	NonMemberBuiltInTypeCode BuiltInTypeCode = "nonmember"
)
