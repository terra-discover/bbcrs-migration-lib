package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TermCategoryTerm model
type TermCategoryTerm struct {
	basic.Base
	basic.DataOwner
	TermID         *uuid.UUID `json:"term_id,omitempty" gorm:"type:varchar(36);uniqueIndex:term_id_unique" swaggertype:"string" format:"uuid"`                   // Term id
	TermCategoryID *uuid.UUID `json:"term_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:term_category_id_unique" swaggertype:"string" format:"uuid"` // Term id
}
