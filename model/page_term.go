package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PageTerm struct
type PageTerm struct {
	basic.Base
	basic.DataOwner
	PageTermAPI
	Page *Page `json:"page,omitempty"` // Page
	Term *Term `json:"term,omitempty"` // Term
}

// PageTermAPI struct
type PageTermAPI struct {
	PageID *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Page ID
	TermID *uuid.UUID `json:"term_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Page ID
}
