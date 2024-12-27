package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Airline Airline
type Airline struct {
	basic.Base
	basic.DataOwner
	AirlineAPI
	AirlineTranslation *AirlineTranslation `json:"airline_translation,omitempty"` // Airline Translation
	AirlineAsset       *AirlineAsset       `json:"airline_asset,omitempty"`       // Airline Asset
	Company            *Company            `json:"company,omitempty"`             // Company Details
}

// AirlineAPI Airline API
type AirlineAPI struct {
	AirlineCode *string    `json:"airline_code,omitempty" gorm:"type:varchar(2);not null;index:idx_airline_code_deleted_at,unique,where:deleted_at is null" example:"GA"`                // Airline Code
	AirlineName *string    `json:"airline_name,omitempty" gorm:"type:varchar(64);not null;index:idx_airline_name_deleted_at,unique,where:deleted_at is null" example:"Garuda Indonesia"` // Airline Name
	NumericCode *string    `json:"numeric_code,omitempty" gorm:"type:varchar(3)" example:"GIA"`                                                                                          // Numeric Code
	CompanyID   *uuid.UUID `json:"company_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                                                                           // Company ID
	// TODO: must review in master-service
	// AirlineAsset *AirlineAssetAPI `json:"airline_asset,omitempty" gorm:"-"`
}

// Seed Airline
func (airline *Airline) Seed() *Airline {
	seed := Airline{
		AirlineAPI: AirlineAPI{
			AirlineCode: lib.Strptr("GA"),
			AirlineName: lib.Strptr("Garuda Indonesia"),
			NumericCode: lib.Strptr("GIA"),
			CompanyID:   lib.UUIDPtr(uuid.New()),
		},
	}
	_ = lib.Merge(seed, &airline)
	return airline
}

// GetAirline by query filter
func (data *Airline) GetAirline(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Preload("AirlineAsset.MultimediaDescription").Take(&data)
}
