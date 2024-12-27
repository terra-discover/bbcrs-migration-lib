package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Subdistrict Subdistrict
type Subdistrict struct {
	basic.Base
	basic.DataOwner
	SubdistrictAPI
	SubdistrictTranslation *SubdistrictTranslation `json:"subdistrict_translation,omitempty"` // Subdistrict Translation
	District               *District               `json:"district,omitempty"`                // District
}

// SubdistrictAPI Subdistrict API
type SubdistrictAPI struct {
	SubdistrictCode *string    `json:"subdistrict_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;index:idx_subdistrict__subdistrict_name,where:deleted_at is null;not null" example:"Gambir"` // Gambir
	SubdistrictName *string    `json:"subdistrict_name,omitempty" gorm:"type:varchar(256);index:idx_subdistrict__subdistrict_name,where:deleted_at is null;not null" example:"Gambir"`                                       // Subdistrict Name
	DistrictID      *uuid.UUID `json:"district_id,omitempty" swaggertype:"string" format:"uuid"`                                                                                                                             // District ID
}
