package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// SegmentCategory Segment Category
type SegmentCategory struct {
	basic.Base
	basic.DataOwner
	SegmentCategoryAPI
	SegmentCategoryTranslation *SegmentCategoryTranslation `json:"segment_category_translation,omitempty"` // Segment Category Translation
}

// SegmentCategoryAPI Segment Category API
type SegmentCategoryAPI struct {
	SegmentCategoryCode *int    `json:"segment_category_code,omitempty" gorm:"type:int;index:,unique,where:deleted_at is null;not null" example:"1"`                  // Segment Category Code
	SegmentCategoryName *string `json:"segment_category_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"All suite"` // Segment Category Name
}

// Seed data
func (r SegmentCategory) Seed() *[]SegmentCategory {
	data := []SegmentCategory{}
	items := []string{
		"All suite",
		"Budget",
		"Corporate business transient",
		"Deluxe",
		"Economy",
		"Extended stay",
		"First class",
		"Luxury",
		"Meeting/Convention",
		"Moderate",
		"Residential apartment",
		"Resort",
		"Tourist",
		"Upscale",
		"Efficiency",
		"Standard",
		"Midscale",
		"Moderate 2",
		"Quality",
		"Quality 2",
		"Unknown",
		"Midscale without F&B",
		"Upper upscale",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, SegmentCategory{
			SegmentCategoryAPI: SegmentCategoryAPI{
				SegmentCategoryCode: &code,
				SegmentCategoryName: &name,
			},
		})
	}

	return &data
}
