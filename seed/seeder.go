package migration

import (
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SeedAll data
// params trySimulateSeederFail only TRUE on testing
func SeedAll(db *gorm.DB, req DataSeedsRequest, trySimulateSeederFail ...bool) error {
	isTryFail := len(trySimulateSeederFail) > 0 &&
		trySimulateSeederFail[0]

	return db.Transaction(func(tx *gorm.DB) error {
		seeds := DataSeeds(req)
		for i := range seeds {
			err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(seeds[i]).Error
			if isTryFail {
				err = errors.New("simulate seeder fail")
			}
			if nil != err {
				name := reflect.TypeOf(seeds[i]).String()
				errorMessage := err.Error()
				return fmt.Errorf("%s seeder fail with %s", name, errorMessage)
			}
		}
		return nil
	})
}
