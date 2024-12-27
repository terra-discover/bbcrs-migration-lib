package model

import (
	"strconv"
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateRatingType Corporate Rating Type
type CorporateRatingType struct {
	basic.Base
	basic.DataOwner
	CorporateRatingTypeAPI
	CorporateRatingTypeTranslation *CorporateRatingTypeTranslation `json:"corporate_rating_type_translation,omitempty"` // Corporate Rating Type Translation
}

// CorporateRatingTypeAPI Corporate Rating Type API
type CorporateRatingTypeAPI struct {
	CorporateRatingTypeCode *int    `json:"corporate_rating_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`           // Corporate Rating Type Code
	CorporateRatingTypeName *string `json:"corporate_rating_type_name,omitempty" example:"Summary" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Corporate Rating Type Name
	Scale                   *int    `json:"scale,omitempty" gorm:"type:smallint"`                                                                                            // Scale
	RatingSymbol            *string `json:"rating_symbol,omitempty" example:"Star" gorm:"type:varchar(256)"`                                                                 // Rating Symbol
}

// Seed data
func (s CorporateRatingType) Seed() *[]CorporateRatingType {
	data := []CorporateRatingType{}
	items := []string{
		"1|Summary|5|Star",
		"2|Transaction Value|5|Star",
		"3|Transaction Volume|5|Star",
		"4|Payment Punctuality|5|Star",
		"5|Payment Status|5|Star",
		"6|Corporate Type|5|Star",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		code, _ := strconv.Atoi(contents[0])
		scale, _ := strconv.Atoi(contents[2])
		var name string = contents[1]
		var symbol string = contents[3]
		data = append(data, CorporateRatingType{
			CorporateRatingTypeAPI: CorporateRatingTypeAPI{
				CorporateRatingTypeCode: &code,
				CorporateRatingTypeName: &name,
				Scale:                   &scale,
				RatingSymbol:            &symbol,
			},
		})
	}

	return &data
}
