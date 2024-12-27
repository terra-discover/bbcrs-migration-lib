package model

import (
	"log"
	"sync"
	"time"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// FlightCaching model
type FlightCaching struct {
	basic.Base
	basic.DataOwner
	GroupID                   *int                                  `json:"group_id,omitempty"`
	SessionID                 *uuid.UUID                            `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_session_id;not null"`
	ExpiredIn                 *int64                                `json:"expired_in,omitempty"`
	SourceType                *string                               `json:"source_type,omitempty"`
	IsLowestFare              *bool                                 `json:"is_lowest_fare"`
	IsPreferredTravelPolicy   *bool                                 `json:"is_preferred_travel_policy" gorm:"default:false"`
	IsNotEligibleTravelPolicy *bool                                 `json:"is_not_eligible_travel_policy"`
	IsTooCloseToDeparture     *bool                                 `json:"is_too_close_to_departure" gorm:"default:false"`
	TotalItemsDuration        *int64                                `json:"total_items_duration" gorm:"-"`
	Violations                *[]FlightCachingTravelPolicyViolation `json:"violations"`
	FlightCachingAirlines     *[]FlightCachingAirlines              `json:"airlines,omitempty" gorm:"foreignKey:FlightCachingID;references:ID"`
	FlightCachingCabins       *[]FlightCachingCabins                `json:"cabins,omitempty" gorm:"foreignKey:FlightCachingID;references:ID"`
}

// GetFlightSegment by query filter
func (data *FlightCaching) GetFlightCaching(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).
		Preload(`FlightCachingAirlines.FlightCachingRoutes`).
		Preload(`FlightCachingCabins.FlightCachingFares`).
		Preload(clause.Associations).Take(&data)
}

// SetCachingTrx
func (d *FlightCaching) SetCachingTrx(db *gorm.DB, datas []interface{}) {
	start := time.Now().UnixNano()

	wg := sync.WaitGroup{}
	wg.Add(len(datas))
	for i := range datas {
		go func(batchObjects interface{}) {
			tx := db.Begin()

			tx = tx.Session(&gorm.Session{
				SkipDefaultTransaction: true,
			})

			defer func() {
				if r := recover(); r != nil {
					log.Printf("ERR SetCachingTrx 1 : %+v", r)
					tx.Rollback()
				}
			}()

			if err := tx.CreateInBatches(batchObjects, 10).Error; nil != err {
				log.Printf("ERR SetCachingTrx 2 : %+v", err)
				tx.Rollback()
			}

			if err := tx.Commit().Error; nil != err {
				log.Printf("ERR SetCachingTrx 3 : %+v", err)
				tx.Rollback()
			}

			wg.Done()
		}(datas[i])
	}

	wg.Wait()
	end := time.Now().UnixNano()
	log.Printf("SetCachingTrx (LEN=%d) DURATION %.2f", len(datas), float64(end-start)/1000000000)
}
