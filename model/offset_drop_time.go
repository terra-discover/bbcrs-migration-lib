package model

import (
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OffsetDropTime Offset Drop Time
type OffsetDropTime struct {
	basic.Base
	basic.DataOwner
	OffsetDropTimeAPI
	OffsetDropTimeTranslation *OffsetDropTimeTranslation `json:"offset_drop_time_translation,omitempty"` // Offset Drop Time Translation
}

// OffsetDropTimeAPI Offset Drop Time API
type OffsetDropTimeAPI struct {
	OffsetDropTimeCode *string `json:"offset_drop_time_code,omitempty" gorm:"type:varchar(16);index:,unique,where:deleted_at is null;not null" example:"BeforeArrival"`   // Offset Drop Time Code
	OffsetDropTimeName *string `json:"offset_drop_time_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Before Arrival"` // Offset Drop Time Name
	Offset             *string `json:"offset,omitempty" gorm:"type:varchar(8)"`                                                                                           // Offset
}

// Seed data
func (s OffsetDropTime) Seed() *[]OffsetDropTime {
	data := []OffsetDropTime{}
	items := []string{
		"BeforeArrival|Before Arrival",
		"BeforeDeparture|Before Departure",
		"AfterBooking|After Booking",
		"AfterConfirmatio|After Confirmation",
		"AfterArrival|After Arrival",
		"AfterDeparture|After Departure",
		"BeforeDueDate|Before Due Date",
		"AfterDueDate|After Due Date",
		"AfterPurchase|After Purchase",
		"AfterRequest|After Request",
		"BeforeExpiry|Before Expiry",
		"AfterExpiry|After Expiry",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		var code string = contents[0]
		var name string = contents[1]
		data = append(data, OffsetDropTime{
			OffsetDropTimeAPI: OffsetDropTimeAPI{
				OffsetDropTimeCode: &code,
				OffsetDropTimeName: &name,
			},
		})
	}

	return &data
}
