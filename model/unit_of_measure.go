package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UnitOfMeasure Unit Of Measure
type UnitOfMeasure struct {
	basic.Base
	basic.DataOwner
	UnitOfMeasureAPI
	UnitOfMeasureTranslation *UnitOfMeasureTranslation `json:"unit_of_measure_translation,omitempty"` // Unit Of Measure Translation
}

// UnitOfMeasureAPI Unit Of Measure API
type UnitOfMeasureAPI struct {
	UnitOfMeasureCode   *int    `json:"unit_of_measure_code,omitempty" gorm:"ype:int;not null;index:,unique,where:deleted_at is null"`           // Unit Of Measure Code
	UnitOfMeasureName   *string `json:"unit_of_measure_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Unit Of Measure Name
	UnitOfMeasureSymbol *string `json:"unit_of_measure_symbol,omitempty" gorm:"type:varchar(8)"`                                                 // Unit Of Measure Symbol
	QuantitySymbol      *string `json:"quantity_symbol,omitempty" gorm:"type:varchar(8)"`                                                        // Quantity Symbol
	QuantityName        *string `json:"quantity_name,omitempty" gorm:"type:varchar(256)"`                                                        // Quantity Name
}

// Seed data
func (s UnitOfMeasure) Seed() *[]UnitOfMeasure {
	data := []UnitOfMeasure{}
	items := []string{
		"1|Miles|mi|L|Length|",
		"2|Kilometers|km|L|Length|",
		"3|Meters|m|L|Length|",
		"4|Millimeters||||",
		"5|Centimeters||||",
		"6|Yards||||",
		"7|Feet||||",
		"8|Inches||||",
		"9|Pixels||||",
		"10|Block||||",
		"11|Megabytes||||",
		"12|Gigabytes||||",
		"13|Square feet|sq ft|A|Area|",
		"14|Square meters|m2|A|Area|",
		"15|Pounds|lb|W|Weight|",
		"16|Kilograms|kg|W|Weight|",
		"17|Square inch||||",
		"18|Square yard||||",
		"19|Acre||||",
		"20|Square millimeter||||",
		"21|Square centimeter||||",
		"22|Hectare||||",
		"23|Ounce||||",
		"24|Gram||||",
		"25|Gallons||||",
		"26|Liters||||",
		"27|Kilowatts||||",
		"28|Cubic meters||||",
		"1001|Kilobytes||||",
	}

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		code := lib.StrToInt(c[0])
		name := c[1]
		symbol := c[2]
		quantitySymbol := c[3]
		quantityName := c[4]

		unit := UnitOfMeasure{}
		unit.UnitOfMeasureCode = &code
		unit.UnitOfMeasureName = &name
		if symbol != "" {
			unit.UnitOfMeasureSymbol = &symbol
		}

		if quantitySymbol != "" {
			unit.QuantitySymbol = &quantitySymbol
		}

		if quantityName != "" {
			unit.QuantityName = &quantityName
		}

		data = append(data, unit)
	}

	return &data

}
