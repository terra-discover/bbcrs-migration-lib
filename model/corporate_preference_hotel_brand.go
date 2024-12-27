package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporatePreferenceHotelBrand Corporate Preference Hotel Brand
type CorporatePreferenceHotelBrand struct {
	basic.Base
	basic.DataOwner
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_preference_hotel_brand,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	CorporatePreferenceHotelBrandAPI
}

// CorporatePreferenceHotelBrandAPI Corporate Preference Hotel Brand Api
type CorporatePreferenceHotelBrandAPI struct {
	HotelBrandID *uuid.UUID `json:"hotel_brand_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_preference_hotel_brand,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
}

// CorporatePreferenceHotelBrandPayload Corporate Preference Hotel Brand Payload
type CorporatePreferenceHotelBrandPayload struct {
	ListHotelBrandID []*uuid.UUID `json:"list_hotel_brand_id,omitempty" gorm:"-:migration"`
}
