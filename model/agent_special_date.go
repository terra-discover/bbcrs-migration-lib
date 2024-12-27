package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// AgentSpecialDate Agent Special Date
type AgentSpecialDate struct {
	basic.Base
	basic.DataOwner
	AgentSpecialDateAPI
	Agent                       *Agent                       `json:"agent,omitempty"`
	AgentSpecialDateTranslation *AgentSpecialDateTranslation `json:"agent_special_date_translation,omitempty"`
}

// AgentSpecialDateAPI Agent Special Date Api
type AgentSpecialDateAPI struct {
	AgentID         *uuid.UUID       `json:"agent_id,omitempty" gorm:"type:varchar(36);not null"`
	SpecialDateName *string          `json:"special_date_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"`
	Fixed           *bool            `json:"fixed,omitempty" gorm:"type:bool"`
	StartDate       *strfmt.DateTime `json:"start_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	EndDate         *strfmt.DateTime `json:"end_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	DayOfTheMonth   *string          `json:"day_of_the_month,omitempty" gorm:"type:varchar(64)"`
	Months          *string          `json:"months,omitempty" gorm:"type:varchar(64)"`
	Description     *string          `json:"description,omitempty" gorm:"type:text"`
}

// AgentSpecialDateDataGet Agent Special Date Data Get
type AgentSpecialDateDataGet struct {
	basic.Base
	basic.DataOwner
	AgentID         *uuid.UUID `json:"agent_id,omitempty"`
	SpecialDateName *string    `json:"special_date_name,omitempty"`
	Fixed           *bool      `json:"fixed,omitempty"`
	StartDate       *string    `json:"start_date,omitempty"`
	EndDate         *string    `json:"end_date,omitempty"`
	StartDateSorter *float64   `json:"start_date_sorter,omitempty"`
	EndDateSorter   *float64   `json:"end_date_sorter,omitempty"`
	DayOfTheMonth   *string    `json:"day_of_the_month,omitempty"`
	Months          *string    `json:"months,omitempty"`
	Description     *string    `json:"description,omitempty"`
}

// Seed AgentSpecialDate
func (agent *AgentSpecialDate) Seed() *AgentSpecialDate {
	seed := AgentSpecialDate{
		AgentSpecialDateAPI: AgentSpecialDateAPI{
			AgentID:         lib.UUIDPtr(uuid.New()),
			SpecialDateName: lib.Strptr("New Year"),
			Fixed:           lib.Boolptr(true),
			StartDate:       nil,
			EndDate:         nil,
			DayOfTheMonth:   nil,
			Months:          nil,
			Description:     nil,
		}}
	_ = lib.Merge(seed, &agent)
	return agent
}
