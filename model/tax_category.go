package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TaxCategory model
type TaxCategory struct {
	basic.Base
	basic.DataOwner
	TaxCategoryAPI
	TaxRate []TaxRate `json:"tax_rate,omitempty" gorm:"foreignKey:TaxCategoryID;references:ID"`
}

// TaxCategoryAPI model
type TaxCategoryAPI struct {
	TaxCategoryCode *string `json:"tax_category_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null"`
	TaxCategoryName *string `json:"tax_category_name,omitempty" gorm:"type:varchar(256);not null;index:idx_tax_category_name_deleted_at,unique,where:deleted_at is null"`
	TaxSymbol       *string `json:"tax_symbol,omitempty" gorm:"type:varchar(3)"`
	Description     *string `json:"description,omitempty" gorm:"type:text"`

	// internal use
	AgentID                          *uuid.UUID `json:"agent_id,omitempty" swaggerignore:"true" gorm:"-"`
	DomesticFlightVAT                *TaxRate   `json:"domestic_flight_vat,omitempty" gorm:"-"`
	DomesticFlightServiceFeeVAT      *TaxRate   `json:"domestic_flight_service_fee_vat,omitempty" gorm:"-"`
	DomesticHotelVAT                 *TaxRate   `json:"domestic_hotel_vat,omitempty" gorm:"-"`
	DomesticHotelServiceFeeVAT       *TaxRate   `json:"domestic_hotel_service_fee_vat,omitempty" gorm:"-"`
	InternationalFlightVAT           *TaxRate   `json:"international_flight_vat,omitempty" gorm:"-"`
	InternationalFlightServiceFeeVAT *TaxRate   `json:"international_flight_service_fee_vat,omitempty" gorm:"-"`
	InternationalHotelVAT            *TaxRate   `json:"international_hotel_vat,omitempty" gorm:"-"`
	InternationalHotelServiceFeeVAT  *TaxRate   `json:"international_hotel_service_fee_vat,omitempty" gorm:"-"`
	OtherVAT                         *TaxRate   `json:"other_vat,omitempty" gorm:"-"`
	OtherServiceFeeVAT               *TaxRate   `json:"other_service_fee_vat,omitempty" gorm:"-"`
}

// AfterCreate TaxCategory
func (t *TaxCategory) AfterCreate(tx *gorm.DB) error {
	return t.QueryHook(tx)
}

// AfterUpdate TaxCategory
func (t *TaxCategory) AfterUpdate(tx *gorm.DB) (err error) {
	return t.QueryHook(tx)
}

// QueryHook TaxCategory
func (t *TaxCategory) QueryHook(tx *gorm.DB) error {
	var agentTaxCategory AgentTaxCategory
	result := tx.Model(&agentTaxCategory).
		Where(tx.Where(AgentTaxCategory{
			AgentTaxCategoryAPI: AgentTaxCategoryAPI{
				AgentID:       t.AgentID,
				TaxCategoryID: t.ID,
			},
		})).
		First(&agentTaxCategory)
	if result.RowsAffected < 1 {
		agentTaxCategory.AgentID = t.AgentID
		agentTaxCategory.TaxCategoryID = t.ID
		err := tx.Create(&agentTaxCategory).Error
		if err != nil {
			return err
		}
	}

	if t.DomesticFlightVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 1).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "18").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.DomesticFlightVAT.ProductTypeID = productType.ID
		t.DomesticFlightVAT.FeeTaxTypeID = feetaxtype.ID

		if t.DomesticFlightVAT.ChargeTypeString != nil {
			t.DomesticFlightVAT.ChargeTypeID = t.ConvertChargeType(tx, t.DomesticFlightVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.DomesticFlightVAT,
			t.ID,
			"domestic",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.DomesticFlightVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.DomesticFlightVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if t.DomesticFlightServiceFeeVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 1).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "23").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.DomesticFlightServiceFeeVAT.ProductTypeID = productType.ID
		t.DomesticFlightServiceFeeVAT.FeeTaxTypeID = feetaxtype.ID
		if t.DomesticFlightServiceFeeVAT.ChargeTypeString != nil {
			t.DomesticFlightServiceFeeVAT.ChargeTypeID = t.ConvertChargeType(tx, t.DomesticFlightServiceFeeVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.DomesticFlightServiceFeeVAT,
			t.ID,
			"domestic",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.DomesticFlightServiceFeeVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.DomesticFlightServiceFeeVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if t.DomesticHotelVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 2).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "20").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.DomesticHotelVAT.ProductTypeID = productType.ID
		t.DomesticHotelVAT.FeeTaxTypeID = feetaxtype.ID
		if t.DomesticHotelVAT.ChargeTypeString != nil {
			t.DomesticHotelVAT.ChargeTypeID = t.ConvertChargeType(tx, t.DomesticHotelVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.DomesticHotelVAT,
			t.ID,
			"domestic",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.DomesticHotelVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.DomesticHotelVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if t.DomesticHotelServiceFeeVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 2).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "25").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.DomesticHotelServiceFeeVAT.ProductTypeID = productType.ID
		t.DomesticHotelServiceFeeVAT.FeeTaxTypeID = feetaxtype.ID
		if t.DomesticHotelServiceFeeVAT.ChargeTypeString != nil {
			t.DomesticHotelServiceFeeVAT.ChargeTypeID = t.ConvertChargeType(tx, t.DomesticHotelServiceFeeVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.DomesticHotelServiceFeeVAT,
			t.ID,
			"domestic",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.DomesticHotelServiceFeeVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.DomesticHotelServiceFeeVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if t.InternationalFlightVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 1).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "19").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.InternationalFlightVAT.ProductTypeID = productType.ID
		t.InternationalFlightVAT.FeeTaxTypeID = feetaxtype.ID
		if t.InternationalFlightVAT.ChargeTypeString != nil {
			t.InternationalFlightVAT.ChargeTypeID = t.ConvertChargeType(tx, t.InternationalFlightVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.InternationalFlightVAT,
			t.ID,
			"international",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.InternationalFlightVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.InternationalFlightVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if t.InternationalFlightServiceFeeVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 1).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "24").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.InternationalFlightServiceFeeVAT.ProductTypeID = productType.ID
		t.InternationalFlightServiceFeeVAT.FeeTaxTypeID = feetaxtype.ID
		if t.InternationalFlightServiceFeeVAT.ChargeTypeString != nil {
			t.InternationalFlightServiceFeeVAT.ChargeTypeID = t.ConvertChargeType(tx, t.InternationalFlightServiceFeeVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.InternationalFlightServiceFeeVAT,
			t.ID,
			"international",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.InternationalFlightServiceFeeVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.InternationalFlightServiceFeeVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if t.InternationalHotelVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 2).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "21").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.InternationalHotelVAT.ProductTypeID = productType.ID
		t.InternationalHotelVAT.FeeTaxTypeID = feetaxtype.ID
		if t.InternationalHotelVAT.ChargeTypeString != nil {
			t.InternationalHotelVAT.ChargeTypeID = t.ConvertChargeType(tx, t.InternationalHotelVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.InternationalHotelVAT,
			t.ID,
			"international",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.InternationalHotelVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.InternationalHotelVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if t.InternationalHotelServiceFeeVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 2).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "26").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.InternationalHotelServiceFeeVAT.ProductTypeID = productType.ID
		t.InternationalHotelServiceFeeVAT.FeeTaxTypeID = feetaxtype.ID
		if t.InternationalHotelServiceFeeVAT.ChargeTypeString != nil {
			t.InternationalHotelServiceFeeVAT.ChargeTypeID = t.ConvertChargeType(tx, t.InternationalHotelServiceFeeVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.InternationalHotelServiceFeeVAT,
			t.ID,
			"international",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.InternationalHotelServiceFeeVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.InternationalHotelServiceFeeVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if t.OtherVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 3).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "17").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.OtherVAT.ProductTypeID = productType.ID
		t.OtherVAT.FeeTaxTypeID = feetaxtype.ID
		if t.OtherVAT.ChargeTypeString != nil {
			t.OtherVAT.ChargeTypeID = t.ConvertChargeType(tx, t.OtherVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.OtherVAT,
			t.ID,
			"other",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.OtherVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.OtherVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	if t.OtherServiceFeeVAT != nil {
		var productType ProductType
		var feetaxtype FeeTaxType
		if err := tx.Where("product_type_code = ?", 3).First(&productType).Error; err != nil {
			return err
		}
		if err := tx.Where("fee_tax_type_code = ?", "22").First(&feetaxtype).Error; err != nil {
			return err
		}

		// t.OtherServiceFeeVAT.ProductTypeID = productType.ID
		t.OtherServiceFeeVAT.FeeTaxTypeID = feetaxtype.ID
		if t.OtherServiceFeeVAT.ChargeTypeString != nil {
			t.OtherServiceFeeVAT.ChargeTypeID = t.ConvertChargeType(tx, t.OtherServiceFeeVAT.ChargeTypeString)
		}
		err := t.InsertTaxRate(
			tx,
			t.OtherServiceFeeVAT,
			t.ID,
			"other",
		)

		// Create or Update Tax Rate Product Type
		var taxRateProductType TaxRateProductType
		taxRateProductType.TaxRateID = t.OtherServiceFeeVAT.ID
		taxRateProductType.ProductTypeID = productType.ID

		//Exist Data
		var existTaxRateProductType TaxRateProductType
		result = tx.Model(&existTaxRateProductType).
			Where(tx.Where(TaxRateProductType{
				TaxRateProductTypeAPI: TaxRateProductTypeAPI{
					TaxRateID: t.OtherServiceFeeVAT.ID,
				},
			})).
			First(&existTaxRateProductType)
		if result.RowsAffected < 1 {
			err := tx.Create(&taxRateProductType).Error
			if err != nil {
				return err
			}

		} else {
			err := tx.Where(&existTaxRateProductType).Updates(&taxRateProductType).Error
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// InsertTaxRate TaxCategory
func (t *TaxCategory) InsertTaxRate(
	tx *gorm.DB,
	taxRate *TaxRate,
	taxCategoryID *uuid.UUID,
	// productTypeID *uuid.UUID,
	destinationCode string,
) error {
	var data *TaxRate
	var currencyID string
	data = taxRate
	data.TaxCategoryID = taxCategoryID
	// db := services.DB
	if data.Amount != nil {
		tx.Select("id").Model(&Currency{}).Where("currency_code = ?", "IDR").Scan(&currencyID)
		data.CurrencyID = lib.UUIDPtr(uuid.MustParse(currencyID))
	}

	var existTaxRate TaxRate
	result := tx.Model(&existTaxRate).
		Where(tx.Where(TaxRate{
			TaxRateAPI: TaxRateAPI{
				TaxCategoryID: t.ID,
				FeeTaxTypeID:  taxRate.FeeTaxTypeID,
			},
		})).
		First(&existTaxRate)
	if result.RowsAffected < 1 {
		err := tx.Create(&data).Error
		if err != nil {
			return err
		}

	} else {
		err := tx.Model(&existTaxRate).Updates(map[string]interface{}{
			"tax_category_id": data.TaxCategoryID,
			"fee_tax_type_id": data.FeeTaxTypeID,
			// "product_type_id": data.ProductTypeID,
			"amount":         data.Amount,
			"currency_id":    data.CurrencyID,
			"percent":        data.Percent,
			"charge_type_id": data.ChargeTypeID,
		}).Error
		if err != nil {
			return err
		}
		data.ID = existTaxRate.ID
	}

	var taxRateProductType TaxRateProductType
	taxRateProductType.TaxRateID = data.ID

	var feeTaxType FeeTaxType
	if err := tx.Where("id = ?", data.FeeTaxTypeID).Take(&feeTaxType).Error; err != nil {
		return err
	}

	var destinationGroup DestinationGroup
	tx.Model(&destinationGroup).
		Where(tx.Where(DestinationGroup{
			DestinationGroupAPI: DestinationGroupAPI{
				DestinationGroupCode: lib.Strptr(destinationCode),
			},
		})).
		First(&destinationGroup)

	var taxRateDestinationGroup TaxRateDestinationGroup
	taxRateDestinationGroup.TaxRateID = data.ID
	taxRateDestinationGroup.DestinationGroupID = destinationGroup.ID
	taxRateDestinationGroup.IsIncluded = lib.Boolptr(true)

	var existTaxRateDestinationGroup TaxRateDestinationGroup
	result = tx.Model(&existTaxRateDestinationGroup).
		Where(tx.Where(TaxRateDestinationGroup{
			TaxRateDestinationGroupAPI: TaxRateDestinationGroupAPI{
				TaxRateID: data.ID,
			},
		})).
		First(&existTaxRateDestinationGroup)
	if result.RowsAffected < 1 {
		err := tx.Create(&taxRateDestinationGroup).Error
		if err != nil {
			return err
		}
	} else {
		err := tx.Where(&existTaxRateDestinationGroup).Updates(&taxRateDestinationGroup).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// ConvertChargeType TaxCategory
func (t *TaxCategory) ConvertChargeType(tx *gorm.DB, chargeType *string) *uuid.UUID {
	var chargeTypeID string
	var chargeTypeCode int

	if *chargeType == "/Ticket" {
		chargeTypeCode = 1001

	} else if *chargeType == "/Person" {
		chargeTypeCode = 7

	} else if *chargeType == "/Transaction" {
		chargeTypeCode = 26

	} else if *chargeType == "/Room Night" {
		chargeTypeCode = 19

	} else if *chargeType == "/Room" {
		chargeTypeCode = 25

	} else if *chargeType == "/Unit" {
		chargeTypeCode = 31
	}

	result := tx.Table("charge_type").
		Select("id").
		Where("charge_type_code = ?", chargeTypeCode).
		Scan(&chargeTypeID)

	if result.RowsAffected < 1 {
		chargeTypeUUID, _ := uuid.Parse(*chargeType)
		return &chargeTypeUUID
	}

	return lib.UUIDPtr(uuid.MustParse(chargeTypeID))
}
