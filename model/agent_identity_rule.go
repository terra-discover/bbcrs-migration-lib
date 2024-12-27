package model

import (
	"strings"

	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentIdentityRule Agent Identity Rule
type AgentIdentityRule struct {
	basic.Base
	basic.DataOwner
	AgentIdentityRuleAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentIdentityRuleAPI Agent Identity Rule Api
type AgentIdentityRuleAPI struct {
	AgentID         *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"`
	IdentityCode    *string    `json:"identity_code,omitempty" gorm:"type:varchar(16);index:,unique,where:deleted_at is null;not null"`
	IdentityName    *string    `json:"identity_name,omitempty" gorm:"type:varchar(256);not null"`
	Prefix          *string    `json:"prefix,omitempty" gorm:"type:varchar(16)"`
	DynamicPrefix   *string    `json:"dynamic_prefix,omitempty" gorm:"varchar(36)"`
	NextNumber      *string    `json:"next_number,omitempty" gorm:"varchar(16)"`
	IsReset         *bool      `json:"is_reset,omitempty" gorm:"type:bool"`
	ResetFrequency  *string    `json:"reset_frequency,omitempty" gorm:"type:char(1)"`
	LastNumber      *string    `json:"last_number,omitempty" gorm:"varchar(16)"`
	LastUsedCounter *int64     `json:"last_used_counter,omitempty" gorm:"type:bigint"`
	LastYear        *int       `json:"last_year,omitempty" gorm:"type:smallint"`
	LastMonth       *int       `json:"last_month,omitempty" gorm:"type:smallint"`
	Description     *string    `json:"description,omitempty" gorm:"type:text"`
}

// Seed Seed data
func (f *AgentIdentityRule) Seed(agentID uuid.UUID) *[]AgentIdentityRule {
	data := []AgentIdentityRule{}

	items := []string{
		"1|Reservation Code",
		"2|Cancellation Code",
		"3|Proposal Code",
		"4|Room Stay Code",
		"5|Reservation Request Code",
		"6|Over Credit Authorization Reference",
		"7|Corporate Balance Transaction Number",
		"8|Payment Transaction Number",
		"9|Reward Balance Transaction Number",
		"10|Invoice Code",
		"11|Product Code",
		"12|Voucher Type Code",
		"13|Voucher Code",
		"14|Voucher Transaction Number",
	}

	for i := range items {
		if lib.IsEmptyUUID(agentID) {
			agentID = uuid.New()
		}
		content := strings.Split(items[i], "|")
		code := content[0]
		name := content[1]
		desc := content[0] + " - " + content[1]
		data = append(data, AgentIdentityRule{
			AgentIdentityRuleAPI: AgentIdentityRuleAPI{
				AgentID:      &agentID,
				IdentityCode: &code,
				IdentityName: &name,
				Description:  &desc,
			},
		})
	}

	return &data
}
