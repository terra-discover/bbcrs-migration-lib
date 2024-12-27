package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProposalProductPriceType struct {
	basic.Base
	basic.DataOwner
	ProposalProductPriceTypeAPI
}

type ProposalProductPriceTypeAPI struct {
	ProposalProductID    *uuid.UUID `json:"proposal_product_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProductPriceTypeID   *uuid.UUID `json:"product_price_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	AgeQualifyingGroupID *uuid.UUID `json:"age_qualifying_group_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	MinAge               *int       `json:"min_age,omitempty" gorm:"type:smallint"`
	MaxAge               *int       `json:"max_age,omitempty" gorm:"type:smallint"`
}
