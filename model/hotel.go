package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// Hotel struct
type Hotel struct {
	basic.Base
	basic.DataOwner
	HotelAPI
	HotelTranslation *HotelTranslation `json:"hotel_translation,omitempty"` // Hotel Translation
	BusinessEntity   *BusinessEntity   `json:"business_entity,omitempty"`   // Business Entity
	HotelAsset       *HotelAsset       `json:"hotel_asset,omitempty"`       // Hotel Asset
	HotelRating      *HotelRating      `json:"hotel_rating,omitempty"`
}

// HotelAPI Hotel API
type HotelAPI struct {
	HotelCode        *string        `json:"hotel_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`  // Hotel Code
	HotelName        *string        `json:"hotel_name,omitempty" gorm:"type:varchar(128);not null;index:,unique,where:deleted_at is null"` // Hotel Name
	HotelStatus      *int           `json:"hotel_status,omitempty" gorm:"type:smallint"`                                                   // Hotel Status
	BusinessEntityID *uuid.UUID     `json:"business_entity_id,omitempty" gorm:"type:varchar(36)"`                                          // Business Entity ID
	WhenBuilt        *string        `json:"when_built,omitempty" gorm:"type:varchar(16)"`                                                  // When Built
	LastRenovation   *string        `json:"last_renovation,omitempty" gorm:"type:varchar(16)"`                                             // Last Renovation
	AreaWeather      *string        `json:"area_weather,omitempty" gorm:"type:text"`                                                       // Area Weather
	HotelBrandID     *uuid.UUID     `json:"hotel_brand_id,omitempty" gorm:"type:varchar(36)"`                                              // Hotel Brand ID
	Comment          *string        `json:"comment,omitempty" gorm:"type:text"`                                                            // Comment
	Description      *string        `json:"description,omitempty" gorm:"type:text"`                                                        // Description
	HotelAsset       *HotelAssetAPI `json:"hotel_asset,omitempty" gorm:"-"`                                                                // Hotel Asset
}

// Seed Hotel
func (hotel *Hotel) Seed() *Hotel {
	seed := Hotel{
		HotelAPI: HotelAPI{
			HotelCode:        lib.Strptr("Hotel Code"),
			HotelName:        lib.Strptr("Hotel Name"),
			HotelStatus:      lib.Intptr(1),
			BusinessEntityID: lib.UUIDPtr(uuid.New()),
			WhenBuilt:        nil,
			LastRenovation:   nil,
			AreaWeather:      nil,
			HotelBrandID:     nil,
			Comment:          nil,
			Description:      nil,
		},
	}
	_ = lib.Merge(seed, &hotel)
	return hotel
}

// HotelProfileDataGet struct
type HotelProfileDataGet struct {
	ID               *uuid.UUID       `json:"id,omitempty"`
	HotelBrandID     *uuid.UUID       `json:"-"`
	CityID           *uuid.UUID       `json:"-"`
	CountryID        *uuid.UUID       `json:"-"`
	BusinessEntityID *uuid.UUID       `json:"business_entity_id"`
	CreatedAt        *strfmt.DateTime `json:"created_at,omitempty" swaggertype:"string" format:"date-time"` // created at
	UpdatedAt        *strfmt.DateTime `json:"updated_at,omitempty" swaggertype:"string" format:"date-time"` // updated at
	Sort             *int64           `json:"sort,omitempty"`                                               // sort
	HotelCode        *string          `json:"hotel_code,omitempty"`                                         // Hotel Code
	HotelName        *string          `json:"hotel_name,omitempty"`                                         // Hotel Name
	HotelStatus      *int             `json:"hotel_status,omitempty" `                                      // Hotel Status                                // Business Entity ID
	WhenBuilt        *string          `json:"when_built,omitempty"`                                         // When Built
	LastRenovation   *string          `json:"last_renovation,omitempty"`                                    // Last Renovation
	AreaWeather      *string          `json:"area_weather,omitempty" `                                      // Area Weather
	Comment          *string          `json:"comment,omitempty" `                                           // Comment
	Description      *string          `json:"description,omitempty" `                                       // Description
	HotelProfileDataGetAPI
}

// HotelProfileDataGetAPI struct
type HotelProfileDataGetAPI struct {
	HotelBrand     *HotelBrand     `json:"hotel_brand,omitempty" gorm:"foreignKey:HotelBrandID;references:ID"`
	City           *City           `json:"city,omitempty" gorm:"foreignKey:CityID;references:ID"`
	Country        *Country        `json:"country,omitempty" gorm:"foreignKey:CountryID;references:ID"`
	BusinessEntity *BusinessEntity `json:"business_entity,omitempty" gorm:"foreignKey:BusinessEntityID;references:ID"`
	HotelAsset     *HotelAsset     `json:"hotel_asset,omitempty" gorm:"foreignKey:ID"`
}

// get hotel data for /master/hotels
type HotelDataGet struct {
	Hotel
	Address *string `json:"address,omitempty" ` // Hotel Address
}
