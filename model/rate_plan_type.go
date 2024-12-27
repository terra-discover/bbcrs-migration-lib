package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// RatePlanType Rate Plan Type
type RatePlanType struct {
	basic.Base
	basic.DataOwner
	RatePlanTypeAPI
	RatePlanTypeTranslation *RatePlanTypeTranslation `json:"rate_plan_type_translation,omitempty"` // Rate Plan Type Translation
}

// RatePlanTypeAPI Rate Plan Type API
type RatePlanTypeAPI struct {
	RatePlanTypeCode *int    `json:"rate_plan_type_code,omitempty" gorm:"type:int;index:,unique,where:deleted_at is null;not null" example:"1"`                    // Rate Plan Type Code
	RatePlanTypeName *string `json:"rate_plan_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Association"` // Rate Plan Type Name
}

// Seed data
func (s RatePlanType) Seed() *[]RatePlanType {
	data := []RatePlanType{}
	items := []string{
		"Association",
		"Club/concierge",
		"Convention",
		"Corporate",
		"Day rate",
		"Distressed inventory",
		"Family plan",
		"Government",
		"Military",
		"Negotiated",
		"Package",
		"Promotional",
		"Regular/rack",
		"Senior citizen",
		"Tour/wholesale",
		"Travel industry",
		"Weekend ",
		"Last room available",
		"Non-last room available",
		"Consortia ",
		"Group",
		"Contract",
		"Promotional package/tour",
		"Published",
		"Net",
		"Multi-day package",
		"Weekly",
		"Monthly",
		"Standard redemption rate",
		"Discount redemption rate",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, RatePlanType{
			RatePlanTypeAPI: RatePlanTypeAPI{
				RatePlanTypeCode: &code,
				RatePlanTypeName: &name,
			},
		})
	}

	return &data
}
