package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerPropertyCategory Integration Partner Property Category
type IntegrationPartnerPropertyCategory struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerPropertyCategoryAPI
	IntegrationPartnerPropertyCategoryTranslation *IntegrationPartnerPropertyCategoryTranslation `json:"integration_partner_property_category_translation,omitempty"` // Integration Partner Property Category Translation
	PropertyCategory                              *PropertyCategory                              `json:"property_category,omitempty"`                                 // Property Category
}

// IntegrationPartnerPropertyCategoryAPI Integration Partner Property Category API
type IntegrationPartnerPropertyCategoryAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_property_category__property_category_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	PropertyCategoryID   *uuid.UUID `json:"property_category_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_property_category__property_category_code,unique,where:deleted_at is null;not null"`   // Property Category ID
	PropertyCategoryCode *string    `json:"property_category_code,omitempty" gorm:"type:varchar(8);index:idx_integration_partner_property_category__property_category_code,unique,where:deleted_at is null;not null"`                                     // Property Category Code
	PropertyCategoryName *string    `json:"property_category_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                                    // Property Category Name
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                                       // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                                           // Comment
}
