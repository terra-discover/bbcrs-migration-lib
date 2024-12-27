package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentAnalytic AgentAnalytic
type AgentAnalytic struct {
	basic.Base
	basic.DataOwner
	AgentAnalyticAPI
	Agent            *Agent            `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	ProductTypeID    *ProductType      `json:"product_type" gorm:"foreignKey:ProductTypeID;references:ID"`
	AirlineID        *Airline          `json:"airline" gorm:"foreignKey:AgentID;references:ID"`
	Hotel            *Hotel            `json:"hotel" gorm:"foreignKey:HotelID;references:ID"`
	DestinationGroup *DestinationGroup `json:"destination_group" gorm:"foreignKey:DestinationGroupID;references:ID"`
	City             *City             `json:"city" gorm:"foreignKey:CityID;references:ID"`
	Country          *Country          `json:"country" gorm:"foreignKey:CountryID;references:ID"`
	StateProvince    *StateProvince    `json:"state_province" gorm:"foreignKey:StateProvinceID;references:ID"`
	Destination      *Destination      `json:"destination" gorm:"foreignKey:DestinationID;references:ID"`
	CabinType        *CabinType        `json:"cabin_type" gorm:"foreignKey:CabinTypeID;references:ID"`
	Division         *Division         `json:"division" gorm:"foreignKey:DivisionID;references:ID"`
	Office           *Office           `json:"office" gorm:"foreignKey:OfficeID;references:ID"`
	HotelSupplier    *HotelSupplier    `json:"hotel_supplier" gorm:"foreignKey:HotelSupplierID;references:ID"`
	Currency         *Currency         `json:"currency" gorm:"foreignKey:CurrencyID;references:ID"`
}

// AgentAnalyticAPI AgentAnalytic API
type AgentAnalyticAPI struct {
	AgentID                      *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	MetricsID                    *uuid.UUID `json:"metrics_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	MarketID                     *uuid.UUID `json:"market_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	ProductTypeID                *uuid.UUID `json:"product_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	AirlineID                    *uuid.UUID `json:"airline_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	HotelID                      *uuid.UUID `json:"hotel_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	DestinationGroupID           *uuid.UUID `json:"destination_group_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	DepartureAirportLocationCode *string    `json:"departure_airport_location_code,omitempty" gorm:"type:varchar(8);"`
	DepartureCityCode            *string    `json:"departure_city_code,omitempty" swaggertype:"string" gorm:"type:varchar(8);"`
	ArrivalAirportLocationCode   *string    `json:"arrival_airport_location_code,omitempty" swaggertype:"string" gorm:"type:varchar(8);"`
	ArrivalCityCode              *string    `json:"arrival_city_code,omitempty" swaggertype:"string" gorm:"type:varchar(8);"`
	CityID                       *uuid.UUID `json:"city_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	CountryID                    *uuid.UUID `json:"country_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	StateProvinceID              *uuid.UUID `json:"state_province_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	DestinationID                *uuid.UUID `json:"destination_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	CabinTypeID                  *uuid.UUID `json:"cabin_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	DivisionID                   *uuid.UUID `json:"division_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	OfficeID                     *uuid.UUID `json:"office_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	ProjectID                    *uuid.UUID `json:"project_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	CostCenterID                 *uuid.UUID `json:"cost_center_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	HotelSupplierID              *uuid.UUID `json:"hotel_supplier_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	CurrencyID                   *uuid.UUID `json:"currency_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	Value                        *string    `json:"value,omitempty" swaggertype:"string" gorm:"type:varchar(64);"`
	DayOfTheMonth                *string    `json:"day_of_the_month,omitempty" swaggertype:"string" gorm:"type:varchar(64);"`
	Months                       *string    `json:"months,omitempty" swaggertype:"string" gorm:"type:varchar(64);"`
	Years                        *string    `json:"years,omitempty" swaggertype:"string" gorm:"type:varchar(64);"`
	IsTravelDate                 *bool      `json:"is_travel_date,omitempty"`
	IsSummary                    *bool      `json:"is_summary,omitempty"`
}

// TableName ...
func (a *AgentAnalytic) TableName() string {
	return "agent_analytics"
}
