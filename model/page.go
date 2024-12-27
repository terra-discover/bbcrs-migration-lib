package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Page struct
type Page struct {
	basic.Base
	basic.DataOwner
	PageAPI
	PageTranslation *PageTranslation `json:"page_translation,omitempty"` // PageTranslation
	PageType        *PageType        `json:"page_type,omitempty"`        // PageType
	CallOut         *CallOut         `json:"call_out,omitempty"`         // CallOut
	PageAsset       *PageAsset       `json:"page_asset,omitempty"`       // Page Asset
}

// PageAPI Page API
type PageAPI struct {
	PageCode              *string          `json:"page_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" example:"contact-us"`  // Page Code
	PageName              *string          `json:"page_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Contact Us"` // Page Name
	Slug                  *string          `json:"slug,omitempty" gorm:"type:varchar(256)" example:"contact-us"`                                                      // Slug
	HTMLTitle             *string          `json:"html_title,omitempty" gorm:"type:varchar(256)"`                                                                     // Html Title
	PageTypeID            *uuid.UUID       `json:"page_type_id,omitempty" gorm:"type:varchar(36)"`                                                                    // Page Type ID
	IsAuthenticated       *bool            `json:"is_authenticated,omitempty" gorm:"default:false" example:"false"`                                                   // Is Authenticated
	IsHome                *bool            `json:"is_home,omitempty" gorm:"default:false" example:"false"`                                                            // Is Home
	IsFeatured            *bool            `json:"is_featured,omitempty" gorm:"default:false" example:"false"`                                                        // Is Featured
	FeaturedStartDate     *strfmt.DateTime `json:"featured_start_date,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`                     // Featured Start Date
	FeatureEndDate        *strfmt.DateTime `json:"feature_end_date,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`                        // Feature End Date
	HasFeedback           *bool            `json:"has_feedback,omitempty" gorm:"default:false" example:"false"`                                                       // Has Feedback
	CallOutID             *uuid.UUID       `json:"call_out_id,omitempty" gorm:"type:varchar(36)"`                                                                     // Call Out ID
	CallOutStartDate      *strfmt.DateTime `json:"call_out_start_date,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`                     // Call Out Start Date
	CallOutEndDate        *strfmt.DateTime `json:"call_out_end_date,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`                       // Call Out End Date
	IsTemplate            *bool            `json:"is_template,omitempty" gorm:"default:false" example:"false"`                                                        // Is Template
	PublishedAt           *strfmt.DateTime `json:"published_at,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`                            // Published At
	Helpful               *int             `json:"helpful,omitempty" gorm:"type:smallint;" example:"1"`                                                               // Helpful
	Unhelpful             *int             `json:"unhelpful,omitempty" gorm:"type:smallint;" example:"1"`                                                             // Unhelpful
	Excerpt               *string          `json:"excerpt,omitempty" gorm:"type:text"`                                                                                // Excerpt
	AssignAllDestinations *bool            `json:"assign_all_destinations,omitempty" gorm:"default:false" example:"false"`                                            // Assign All Destinations
	Description           *string          `json:"description,omitempty" gorm:"type:text"`                                                                            // Description
	PageAsset             *PageAssetAPI    `json:"page_asset,omitempty" gorm:"-"`                                                                                     // PageAssetAPI
}

// Seed Page
func (page *Page) Seed() *Page {
	seed := Page{
		PageAPI: PageAPI{
			PageCode:              lib.Strptr("contact-us"),
			PageName:              lib.Strptr("Contact Us"),
			Slug:                  lib.Strptr("contact-us"),
			HTMLTitle:             nil,
			PageTypeID:            lib.UUIDPtr(uuid.New()),
			IsAuthenticated:       lib.Boolptr(false),
			IsHome:                lib.Boolptr(false),
			IsFeatured:            lib.Boolptr(false),
			FeaturedStartDate:     nil,
			FeatureEndDate:        nil,
			HasFeedback:           lib.Boolptr(false),
			CallOutID:             lib.UUIDPtr(uuid.New()),
			CallOutStartDate:      nil,
			CallOutEndDate:        nil,
			IsTemplate:            lib.Boolptr(false),
			PublishedAt:           nil,
			Helpful:               nil,
			Unhelpful:             nil,
			Excerpt:               nil,
			AssignAllDestinations: lib.Boolptr(false),
			Description:           nil,
		},
	}
	_ = lib.Merge(seed, &page)
	return page
}
