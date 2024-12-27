package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CacheAirTravelerCriteria Cache Air Traveler Criteria
type CacheAirTravelerCriteria struct {
	basic.Base
	basic.DataOwner
	CacheAirTravelerCriteriaAPI
	CacheCriteria *CacheCriteria `json:"cache_criteria,omitempty" gorm:"foreignKey:CacheCriteriaID;references:ID"` // cache criteria
}

// CacheAirTravelerCriteriaAPI Cache Air Traveler Criteria API
type CacheAirTravelerCriteriaAPI struct {
	CacheCriteriaID  *uuid.UUID `json:"cache_criteria_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggerignore:"true"` // cache criteria id
	SeatsRequested   *int       `json:"seats_requested,omitempty" gorm:"type:smallint;" example:"1"`                                     // seats requested
	NumberOfAdults   *int       `json:"number_of_adults,omitempty" gorm:"type:smallint;" example:"1"`                                    // number of adults
	NumberOfChildren *int       `json:"number_of_children,omitempty" gorm:"type:smallint;" example:"1"`                                  // number of children
	NumberOfInfants  *int       `json:"number_of_infants,omitempty" gorm:"type:smallint;" example:"1"`                                   // number of infants
}
