package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CacheGuestCriteria Cache Guest Criteria
type CacheGuestCriteria struct {
	basic.Base
	basic.DataOwner
	CacheGuestCriteriaAPI
	CacheCriteria  *CacheCriteria `json:"cache_criteria,omitempty" gorm:"foreignKey:CacheCriteriaID;references:ID" swaggerignore:"true"` // cache criteria
	CitizenCountry *Country       `json:"citizen_country,omitempty" gorm:"foreignKey:CitizenCountryID;references:ID"`                    // citizen country
}

// CacheGuestCriteriaAPI Cache Guest Criteria API
type CacheGuestCriteriaAPI struct {
	CacheCriteriaID  *uuid.UUID `json:"cache_criteria_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggerignore:"true"` // cache criteria id
	CitizenCountryID *uuid.UUID `json:"citizen_country_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                              // citizen country id
	RoomsRequested   *int       `json:"rooms_requested,omitempty" gorm:"type:smallint;" example:"1"`                                     // rooms requested
	NumberOfAdults   *int       `json:"number_of_adults,omitempty" gorm:"type:smallint;" example:"1"`                                    // number of adults
	NumberOfChildren *int       `json:"number_of_children,omitempty" gorm:"type:smallint;" example:"1"`                                  // number of children
}
