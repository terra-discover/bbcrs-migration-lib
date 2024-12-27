package model

import (
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FeeTaxType Model
type FeeTaxType struct {
	basic.Base
	basic.DataOwner
	FeeTaxTypeAPI
	FeeTaxTypeTranslation *FeeTaxTypeTranslation `json:"fee_tax_type_translation,omitempty"` // Fee Tax Type Translation
}

// FeeTaxTypeAPI API
type FeeTaxTypeAPI struct {
	FeeTaxTypeCode *string `json:"fee_tax_type_code,omitempty" gorm:"type:varchar(36);not null;index:idx_fee_tax_type_code_deleted_at,unique,where:deleted_at is null"` // Fee Tax Type Code
	FeeTaxTypeName *string `json:"fee_tax_type_name,omitempty" gorm:"type:varchar(256);not null;"`                                                                      // Fee Tax Type Name
	FeeCategory    *string `json:"fee_category,omitempty" gorm:"type:varchar(2);"`                                                                                      // Fee Category
	Description    *string `json:"description,omitempty" gorm:"type:text"`                                                                                              // Description
}

// Seed Seed data
func (f *FeeTaxType) Seed() *[]FeeTaxType {
	data := []FeeTaxType{}

	items := []string{
		"1|Service Fee|SF",
		"2|Reissue Fee (Reissue & Revalidate/ Reroute)|PF",
		"3|Cancellation Fee|PF",
		"4|Refund Fee|PF",
		"5|MDR Fee|PF",
		"6|Emergency Service Assistance 24 Hours Surcharge - Issued Only|PF",
		"7|Late Payment|PF",
		"8|Domestic Flight Service Fee|SF",
		"9|International Flight Service Fee|SF",
		"10|Domestic Hotel Service Fee|SF",
		"11|International Hotel Service Fee|SF",
		"12|Markup|MU",
		"13|Domestic Flight Markup|MU",
		"14|International Flight Markup|MU",
		"15|Domestic Hotel Markup|MU",
		"16|International Hotel Markup|MU",
		"17|VAT|TX",
		"18|Domestic Flight VAT|TX",
		"19|International Flight VAT|TX",
		"20|Domestic Hotel VAT|TX",
		"21|International Hotel VAT|TX",
		"22|Service Fee VAT|TX",
		"23|Domestic Flight Service Fee VAT|TX",
		"24|International Flight Service Fee VAT|TX",
		"25|Domestic Hotel Service Fee VAT|TX",
		"26|International Hotel Service Fee VAT|TX",
		"27|Processing Fee|PF",
		"28|RFP Fee (Contact Fee)|PF",
		"29|Void Fee (Same Day)|PF",
		"30|Modify Hotel Fee|PF",
		"31|Hotel Refund Fee|PF",
		"32|Booking Processing Fee|PF",
		"33|Domestic Reissue Fee (Reissue & Reroute)|PF",
		"34|International Reissue Fee (Reissue & Reroute)|PF",
		"35|Domestic Refund Fee|PF",
		"36|International Refund Fee|PF",
		"37|Domestic Void Fee (Same Day)|PF",
		"38|International Void Fee (Same Day)|PF",
		"39|Domestic RFP Fee (Contact Fee)|PF",
		"40|International RFP Fee (Contact Fee)|PF",
		"41|Domestic Non-GDS Flight Booking Process Fee|PF",
		"42|International Non-GDS Flight Booking Process Fee|PF",
		"43|Domestic Modify Hotel Fee|PF",
		"44|International Modify Hotel Fee|PF",
		"45|Domestic Hotel Refund Fee|PF",
		"46|International Hotel Refund Fee|PF",
		"47|Domestic Non-GDS Hotel Booking Process Fee|PF",
		"48|International Non-GDS Hotel Booking Process Fee|PF",
		"49|Domestic Revalidate Fee|PF",
		"50|International Revalidate Fee|PF",
		"1001|Passenger Service Charge|AX",
		"1002|Service Fee - Carrier-Imposed Misc|AX",
		"1003|Mandatory Travel Insurance Tax IWJR|AX",
		"1004|Service Fee - Carrier-Imposed Fuel|AX",
		"1005|Aviation Levy|AX",
		"1006|Service Fee - Carrier-Imposed Misc|AX",
		"1007|Passenger Service and Security Fee PSSF|AX",
	}

	for i := range items {
		content := strings.Split(items[i], "|")
		code := content[0]
		name := content[1]
		cat := content[2]
		desc := content[1] + " - " + content[2]
		data = append(data, FeeTaxType{
			FeeTaxTypeAPI: FeeTaxTypeAPI{
				FeeTaxTypeCode: &code,
				FeeTaxTypeName: &name,
				FeeCategory:    &cat,
				Description:    &desc,
			},
		})
	}

	return &data
}
