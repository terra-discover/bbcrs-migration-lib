package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// Corporate struct
type Corporate struct {
	basic.Base
	basic.DataOwner
	CorporateAPI
	BusinessEntity  *BusinessEntity `json:"business_entity,omitempty"`                       // Business Entity
	CorporateType   *CorporateType  `json:"corporate_type,omitempty"`                        // Corporate Type
	ParentCorporate *Corporate      `json:"parent_corporate,omitempty" swaggerignore:"true"` // Parent Corporate
	CorporateAsset  *CorporateAsset `json:"corporate_asset,omitempty"`                       // Corporate Asset
}

// CorporateAPI struct
type CorporateAPI struct {
	CorporateCode     *string            `json:"corporate_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`  // Corporate Code
	CorporateName     *string            `json:"corporate_name,omitempty" gorm:"type:varchar(512);not null;index:,unique,where:deleted_at is null"` // Corporate Name
	ParentCorporateID *uuid.UUID         `json:"parent_corporate_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`          // Parent Corporate Id
	ProfileID         *uuid.UUID         `json:"profile_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                   // Profile Id
	ProfileStatus     *int               `json:"profile_status,omitempty" gorm:"type:int"`                                                          // Profile Status
	BusinessEntityID  *uuid.UUID         `json:"business_entity_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`           // Business Entity Id
	CorporateTypeID   *uuid.UUID         `json:"corporate_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`            // Corporate Type Id
	Reference         *string            `json:"reference,omitempty" gorm:"type:text"`                                                              // Reference
	Comment           *string            `json:"comment,omitempty" gorm:"type:text"`                                                                // Comment
	Description       *string            `json:"description,omitempty" gorm:"type:text"`                                                            // Description
	CorporateAsset    *CorporateAssetAPI `json:"corporate_asset,omitempty" gorm:"-"`                                                                // Corporate Asset
}

// Seed Corporate
func (c *Corporate) Seed(agentID uuid.UUID) *[]Corporate {
	if lib.IsEmptyUUID(agentID) {
		agentID = uuid.New()
	}
	data := []Corporate{}
	items := []string{
		"1|Bayu Buana",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		code := contents[0]
		name := contents[1]
		data = append(data, Corporate{
			Base: basic.Base{
				ID: &agentID,
			},
			CorporateAPI: CorporateAPI{
				CorporateCode: &code,
				CorporateName: &name,
				ProfileStatus: lib.Intptr(1),
			},
		})
	}

	return &data
}
