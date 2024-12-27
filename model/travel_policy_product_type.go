package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyProductType Travel Policy Product Type
type TravelPolicyProductType struct {
	basic.Base
	basic.DataOwner
	TravelPolicyProductTypeAPI
	TravelPolicy *TravelPolicy `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"` // travel policy
	ProductType  *ProductType  `json:"product_type,omitempty" gorm:"foreignKey:ProductTypeID;reference:ID"`    // product type
}

// TravelPolicyProductTypeAPI Travel Policy Product Type API
type TravelPolicyProductTypeAPI struct {
	TravelPolicyID *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // travel policy id
	ProductTypeID  *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`  // product type id
}
