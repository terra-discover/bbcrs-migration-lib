package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Metrics Metrics
type Metrics struct {
	basic.Base
	basic.DataOwner
	MetricsAPI
	MetricsTranslation *MetricsTranslation `json:"metrics_translation,omitempty"` // Metrics Translation
}

// MetricsAPI Metrics API
type MetricsAPI struct {
	MetricsCode *int    `json:"metrics_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null" example:"1"`                             // Metrics Code
	MetricsName *string `json:"metrics_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null" example:"Average transaction value"` // Metrics Name
	Description *string `json:"description,omitempty" gorm:"type:text"`                                                                                     // Description
}

// Seed data
func (s Metrics) Seed() *[]Metrics {
	data := []Metrics{}
	items := []string{
		"Average transaction value",
		"Average transaction volume",
		"Average payment punctuality",
		"Longest payment overdue",
		"Total transaction value",
		"Total transaction volume",
		"Response time (in minutes)",
		"Resolution time (in minutes)",
		"Total number of trips",
		"Total number of travelers",
		"Total invoice overdue",
		"Closest payment overdue",
		"Total number of tickets",
		"Total number of room nights",
		"Less than 30 days payment overdue",
		"31 - 60 days payment overdue",
		"61 - 90 days payment overdue",
		"More than 90 days payment overdue",
		"Average booking lead time",
		"Average transaction value before tax",
		"Total transaction value before tax",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, Metrics{
			MetricsAPI: MetricsAPI{
				MetricsCode: &code,
				MetricsName: &name,
			},
		})
	}

	return &data
}
