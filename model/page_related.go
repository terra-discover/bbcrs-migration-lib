package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PageRelated struct
type PageRelated struct {
	basic.Base
	basic.DataOwner
	PageRelatedAPI
	Page *Page `json:"page,omitempty"` // Page
	//RelatedPage *RelatedPage `json:"related_page,omitempty"` // wait model
}

// PageRelatedAPI struct
type PageRelatedAPI struct {
	PageID        *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`      // Page ID
	RelatedPageID *uuid.UUID `json:"related_page,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Related Page ID
}
