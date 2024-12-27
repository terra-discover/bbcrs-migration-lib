package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PageWidget struct
type PageWidget struct {
	basic.Base
	basic.DataOwner
	PageWidgetAPI
	Page   *Page   `json:"page,omitempty"` // Page
	Widget *Widget `json:"term,omitempty"` // Widget
}

// PageWidgetAPI struct
type PageWidgetAPI struct {
	PageID   *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Page ID
	WidgetID *uuid.UUID `json:"term_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Page ID
}
