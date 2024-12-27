package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCity Integration Partner City
type IntegrationPartnerCity struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCityAPI
	IntegrationPartnerCityTranslation *IntegrationPartnerCityTranslation `json:"integration_partner_city_translation,omitempty"`                                     // Integration Partner City Translation
	City                              *City                              `json:"city,omitempty" gorm:"foreignKey:CityID;references:ID"`                              // City
	IntegrationPartner                *IntegrationPartner                `json:"integration_partner,omitempty" gorm:"foreignKey:IntegrationPartnerID;references:ID"` // Integration Partner
}

// IntegrationPartnerCityAPI Integration Partner City API
type IntegrationPartnerCityAPI struct {
	IntegrationPartnerID *uuid.UUID       `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;index:idx_idx_integration_partner_city__integration_partner_city_id,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"` // Integration Partner ID
	CityID               *uuid.UUID       `json:"city_id,omitempty" gorm:"type:varchar(36);not null;index:idx_idx_integration_partner_city__integration_partner_city_id,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`                // City ID
	CityCode             *string          `json:"city_code,omitempty" gorm:"type:varchar(36)"`                                                                                                                                                              // City Code
	CityName             *string          `json:"city_name,omitempty" gorm:"type:varchar(64)"`                                                                                                                                                              // City Name
	Latitude             *float64         `json:"latitude,omitempty"`                                                                                                                                                                                       // Latitude
	Longitude            *float64         `json:"longitude,omitempty"`                                                                                                                                                                                      // Longitude
	ForecastAt           *strfmt.DateTime `json:"forecast_at,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`                                                                                                                    // Forecast At
	ForecastExpiration   *strfmt.DateTime `json:"forecast_expiration,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`                                                                                                            // Forecast Expiration
	WeatherForecast      *string          `json:"weather_forecast,omitempty" gorm:"type:text"`                                                                                                                                                              // Weather Forecast
}

// IntegrationPartnerCityDataGet Integration Partner City Data Get
type IntegrationPartnerCityDataGet struct {
	IntegrationPartnerCityDataGetAPI
	ID                       *uuid.UUID       `json:"id,omitempty"`
	IntegrationPartnerID     *uuid.UUID       `json:"-"`
	IntegrationPartnerCityID *uuid.UUID       `json:"-"`
	CityID                   *uuid.UUID       `json:"-"`
	CreatedAt                *strfmt.DateTime `json:"created_at,omitempty" swaggertype:"string" format:"date-time"` // created at
	UpdatedAt                *strfmt.DateTime `json:"updated_at,omitempty" swaggertype:"string" format:"date-time"` // updated at
	Sort                     *int64           `json:"sort,omitempty"`                                               // sort
	Status                   *int64           `json:"status,omitempty"`
}

// IntegrationPartnerCityDataGetAPI Integration Partner City Data Get API
type IntegrationPartnerCityDataGetAPI struct {
	City                   *City                   `json:"city" gorm:"foreignKey:CityID;references:ID"`
	IntegrationPartner     *IntegrationPartner     `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"`
	IntegrationPartnerCity *IntegrationPartnerCity `json:"integration_partner_city" gorm:"foreignKey:IntegrationPartnerCityID;references:ID"`
}
