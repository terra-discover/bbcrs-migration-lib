package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PageLink struct
type PageLink struct {
	basic.Base
	basic.DataOwner
	PageLinkAPI
	Page *Page `json:"page,omitempty"` // Page
	//LinkedModule *LinkedModule `json:"linked_module,omitempty"` //wait model
	//LinkedPage   *LinkedPage   `json:"linked_page,omitempty"` // wait model
}

// PageLinkAPI struct
type PageLinkAPI struct {
	PageID         *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Page ID
	LinkedModuleID *uuid.UUID `json:"linked_module_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // Linked Module ID
	LinkedPageID   *uuid.UUID `json:"linked_page_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`   // Linked Page ID
}
