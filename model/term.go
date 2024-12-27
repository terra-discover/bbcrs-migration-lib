package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Term Term
type Term struct {
	basic.Base
	basic.DataOwner
	TermAPI
	TermTranslation *TermTranslation `json:"term_translation,omitempty"` // Term Translation
}

// TermAPI Term API
type TermAPI struct {
	TermCode        *string    `json:"term_code,omitempty" gorm:"type:varchar(256)"`                // Term Code
	TermName        *string    `json:"term_name,omitempty" gorm:"type:varchar(256)"`                // Term Name
	Slug            *string    `json:"slug,omitempty" gorm:"type:varchar(256)"`                     // Slug
	ParentTermID    *uuid.UUID `json:"parent_term_id,omitempty" swaggertype:"string" format:"uuid"` // Parent Term ID
	HTMLTitle       *string    `json:"html_title,omitempty" gorm:"type:varchar(256)"`               // HTML Title
	IsAuthenticated *bool      `json:"is_authenticated,omitempty"`                                  // Is Authenticated
	Description     *string    `json:"description,omitempty" gorm:"type:text"`                      // Description
}
