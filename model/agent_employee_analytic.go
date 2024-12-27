package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentEmployeeAnalytic Model
type AgentEmployeeAnalytic struct {
	basic.Base
	basic.DataOwner
	AgentEmployeeAnalyticAPI
	EmployeeID       *Employee         `json:"employee" gorm:"foreignKey:EmployeeID;references:ID"`
	DestinationGroup *DestinationGroup `json:"destination_group" gorm:"foreignKey:DestinationGroupID;references:ID"`
	Currency         *Currency         `json:"currency" gorm:"foreignKey:CurrencyID;references:ID"`
}

// AgentEmployeeAnalyticAPI Model
type AgentEmployeeAnalyticAPI struct {
	EmployeeID         *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	MetricsID          *uuid.UUID `json:"metrics_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	MarketID           *uuid.UUID `json:"market_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	ProductTypeID      *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	CurrencyID         *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	Value              *uuid.UUID `json:"value,omitempty" gorm:"type:varchar(64);"`
	Hours              *uuid.UUID `json:"hours,omitempty" gorm:"type:varchar(64);"`
	DayOfTheMonth      *uuid.UUID `json:"day_of_the_month,omitempty" gorm:"type:varchar(64);"`
	Months             *uuid.UUID `json:"months,omitempty" gorm:"type:varchar(64);"`
	Years              *uuid.UUID `json:"years,omitempty" gorm:"type:varchar(64);"`
	IsTravelDate       *bool      `json:"is_travel_date,omitempty"`
	IsSummary          *bool      `json:"is_summary,omitempty"`
}
