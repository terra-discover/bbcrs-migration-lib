package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TermCategory Term Category
type TermCategory struct {
	basic.Base
	basic.DataOwner
	TermCategoryAPI
	TermCategoryTranslation *TermCategoryTranslation `json:"term_category_translation,omitempty"` // Term Category Translation
}

// TermCategoryAPI Term Category API
type TermCategoryAPI struct {
	TermCategoryCode     *string    `json:"term_category_code,omitempty" gorm:"type:varchar(36)"`                 // Term Category Code
	TermCategoryName     *string    `json:"term_category_name,omitempty" gorm:"type:varchar(256)"`                // Term Category Name
	ParentTermCategoryID *uuid.UUID `json:"parent_term_category_id,omitempty" swaggertype:"string" format:"uuid"` // Parent Term Category ID
	IsSystem             *bool      `json:"is_system,omitempty"`                                                  // Is System
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                               // Description
}
