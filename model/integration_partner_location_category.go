package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerLocationCategory Integration Partner Location Category
type IntegrationPartnerLocationCategory struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerLocationCategoryAPI
	IntegrationPartnerLocationCategoryTranslation *IntegrationPartnerLocationCategoryTranslation `json:"integration_partner_location_category_translation,omitempty"` // Integration Partner Location Category Translation
	LocationCategory                              *LocationCategory                              `json:"location_category,omitempty"`                                 // location_category
}

// IntegrationPartnerLocationCategoryAPI Integration Partner Location Category API
type IntegrationPartnerLocationCategoryAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_location_category__location_category_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	LocationCategoryID   *uuid.UUID `json:"location_category_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_location_category__location_category_code,unique,where:deleted_at is null;not null"`   // Location Category ID
	LocationCategoryCode *string    `json:"location_category_code,omitempty" gorm:"type:varchar(8);index:idx_integration_partner_location_category__location_category_code,unique,where:deleted_at is null;not null"`                                     // Location Category Code
	LocationCategoryName *string    `json:"location_category_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                                    // Location Category Name
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                                       // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                                           // Comment
}
