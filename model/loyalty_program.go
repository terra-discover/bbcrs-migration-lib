package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LoyaltyProgram Loyalty Program
type LoyaltyProgram struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramAPI
	LoyaltyProgramTranslation *LoyaltyProgramTranslation `json:"loyalty_program_translation,omitempty"` // Loyalty Program Translation
	LoyaltyProgramAsset       *LoyaltyProgramAsset       `json:"loyalty_program_asset,omitempty"`       // Loyalty Program Asset
}

// LoyaltyProgramDetail Loyalty Program Detail
type LoyaltyProgramDetail struct {
	LoyaltyProgram
	ProductTypeID *uuid.UUID   `json:"product_type_id,omitempty" swaggertype:"string" format:"uuid"` // Product Type ID
	ProductType   *ProductType `json:"product_type,omitempty"`                                       // Product Type
}

// LoyaltyProgramAPI Loyalty Program API
type LoyaltyProgramAPI struct {
	LoyaltyProgramCode  *string                 `json:"loyalty_program_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`  // Loyalty Program Code
	LoyaltyProgramName  *string                 `json:"loyalty_program_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Loyalty Program Name
	Comment             *string                 `json:"comment,omitempty" gorm:"type:text"`                                                                      // Comment
	Description         *string                 `json:"description,omitempty" gorm:"type:text"`                                                                  // Description
	LoyaltyProgramAsset *LoyaltyProgramAssetAPI `json:"loyalty_program_asset,omitempty" gorm:"-"`                                                                // Loyalty Program Asset
	ProductTypeID       *uuid.UUID              `json:"product_type_id,omitempty" validate:"required" gorm:"-" swaggertype:"string" format:"uuid"`               // Product Type ID
}

// BeforeCreate attraction
func (a *LoyaltyProgram) BeforeCreate(tx *gorm.DB) error {
	// calling super class method
	err := a.Base.BeforeCreate(tx)

	if nil != a.ProductTypeID {
		loyaltyProgramProductType := LoyaltyProgramProductType{}
		loyaltyProgramProductType.LoyaltyProgramID = a.ID
		loyaltyProgramProductType.ProductTypeID = a.ProductTypeID
		tx.Create(&loyaltyProgramProductType)
	}

	return err
}

// BeforeUpdate attraction
func (a *LoyaltyProgram) BeforeUpdate(tx *gorm.DB) error {
	// calling super class method
	err := a.Base.BeforeUpdate(tx)

	loyaltyProgramProductType := LoyaltyProgramProductType{}
	res := tx.First(&loyaltyProgramProductType, `loyalty_program_id = ?`, a.ID)
	if res.RowsAffected == 0 {
		a.Base.BeforeCreate(tx)
	} else {
		loyaltyProgramProductType.ProductTypeID = a.ProductTypeID
		tx.Updates(&loyaltyProgramProductType)
	}

	return err
}

// Seed Loyalty Program
func (a *LoyaltyProgram) Seed() *LoyaltyProgram {
	seed := LoyaltyProgram{
		LoyaltyProgramAPI: LoyaltyProgramAPI{
			LoyaltyProgramCode: lib.Strptr("Loyalty Program Code"),
			LoyaltyProgramName: lib.Strptr("Loyalty Program Name"),
			Comment:            nil,
			Description:        nil,
		},
	}
	_ = lib.Merge(seed, &a)
	return a
}
