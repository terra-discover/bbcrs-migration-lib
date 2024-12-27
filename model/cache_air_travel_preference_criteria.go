package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CacheAirTravelPreferenceCriteria Cache Air Travel Preference Criteria
type CacheAirTravelPreferenceCriteria struct {
	basic.Base
	basic.DataOwner
	CacheAirTravelPreferenceCriteriaAPI
	CacheCriteria *CacheCriteria `json:"cache_criteria,omitempty" gorm:"foreignKey:CacheCriteriaID;references:ID" swaggerignore:"true"` // cache criteria
	CabinType     *CabinType     `json:"cabin_type,omitempty" gorm:"foreignKey:CabinTypeID;references:ID"`                              // cabin type
}

// CacheAirTravelPreferenceCriteriaAPI Cache Air Travel Preference Criteria API
type CacheAirTravelPreferenceCriteriaAPI struct {
	CacheCriteriaID  *uuid.UUID `json:"cache_criteria_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggerignore:"true"` // cache criteria id
	CabinTypeID      *uuid.UUID `json:"cabin_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                   // cabin type id
	SeatsRequested   *int       `json:"seats_requested,omitempty" gorm:"type:smallint;" example:"1"`                                     // seats requested
	NumberOfAdults   *int       `json:"number_of_adults,omitempty" gorm:"type:smallint;" example:"1"`                                    // number of adults
	NumberOfChildren *int       `json:"number_of_children,omitempty" gorm:"type:smallint;" example:"1"`                                  // number of children
	NumberOfInfants  *int       `json:"number_of_infants,omitempty" gorm:"type:smallint;" example:"1"`                                   // number of infants
}
