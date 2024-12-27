package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdditionalDetailType Additional Detail Type
type AdditionalDetailType struct {
	basic.Base
	basic.DataOwner
	AdditionalDetailTypeAPI
	AdditionalDetailTypeTranslation *AdditionalDetailTypeTranslation `json:"additional_detail_type_translation,omitempty"`
}

// AdditionalDetailTypeAPI Additional Detail Type API
type AdditionalDetailTypeAPI struct {
	AdditionalDetailTypeCode *int    `json:"additional_detail_type_code,omitempty" gorm:"type:smallint;uniqueIndex:idx_additional_detail_type_code_deleted_at,where:deleted_at is null;not null"`
	AdditionalDetailTypeName *string `json:"additional_detail_type_name,omitempty" gorm:"type:varchar(256);not null"`
}

// Seed data
func (s AdditionalDetailType) Seed() *[]AdditionalDetailType {
	data := []AdditionalDetailType{}
	items := []string{
		"1|Rate description|",
		"2|Property description|",
		"3|Property location|",
		"4|Room information|",
		"5|Guarantee information|",
		"6|Deposit information|",
		"7|Cancellation information|",
		"8|Check in check out information|",
		"9|Extra charge information|",
		"10|Tax information|",
		"11|Service charge information|",
		"12|Package information|",
		"13|Commission information|",
		"14|Miscellaneous information|",
		"15|Promotional information|",
		"16|Inclusion information|",
		"17|Amenity information|",
		"18|Late arrival information|",
		"19|Late departure information|",
		"20|Advanced booking information|",
		"21|Extra person information|",
		"22|Areas served|",
		"23|Onsite facilities information|",
		"24|Offsite facilities information|",
		"25|Onsite services information|",
		"26|Offsite services information|",
		"27|Extended stay information|",
		"28|Corporate booking information|",
		"29|Booking guidelines|",
		"30|Government booking policy|",
		"31|Group booking information|",
		"32|Rate disclaimer information|",
		"33|Visa/travel requirement information|",
		"34|Security information|",
		"35|Onsite recreational activities information|",
		"36|Offsite recreational activities information|",
		"37|General meeting planning information|",
		"38|Group meeting planning information|",
		"39|Contract/negotiated booking information|",
		"40|Travel industry booking information|",
		"41|Meeting room description|",
		"42|Pet policy description|",
		"43|Meal plan description|",
		"44|Family plan description|",
		"45|Children information|",
		"46|Early checkout description|",
		"47|Special offers description|",
		"48|Catering description|",
		"49|Room decor description|",
		"50|Oversold policy description|",
		"51|Last room availability description|",
		"52|Room type upgrade description|",
		"53|Driving directions|",
		"54|Driving directions from the north|",
		"55|Driving directions from the south|",
		"56|Driving directions from the east|",
		"57|Driving directions from the west|",
		"58|Surcharge information|",
		"59|Minimum stay information|",
		"60|Maximum stay information|",
		"61|Check-in policy|",
		"62|Check-out policy|",
		"63|Express check-in policy|",
		"64|Express check-out policy|",
		"65|Facility restrictions|",
		"66|Customs information for material|",
		"67|Seasons|",
		"68|Food and beverage minimums for groups|",
		"69|Deposit policy for master account|",
		"70|Deposit policy for reservations|",
		"71|Restaurant services|",
		"72|Special events|",
		"73|Cuisine description|",
		"1001|Terms and conditions|",
		"1002|Privacy policy|",
		"1003|Special request policy|",
		"1004|Passport/travel requirement information|",
		"1005|Travel information|",
		"1006|Terms of use|",
		"1007|New normal|",
		"1008|Splash screen title|",
		"1009|Flight terms and conditions|",
		"1010|Hotel terms and conditions|",
		"1011|Others terms and conditions|",
		"1012|Passport disclaimer|",
		"1013|Visa disclaimer|",
		"1014|Passport warning|",
		"1015|Visa warning|",
		"2001|Exterior view|",
		"2002|Lobby view|",
		"2003|Pool view|",
		"2004|Restaurant|",
	}

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		code := lib.StrToInt(c[0])
		name := c[1]

		data = append(data, AdditionalDetailType{
			AdditionalDetailTypeAPI: AdditionalDetailTypeAPI{
				AdditionalDetailTypeCode: &code,
				AdditionalDetailTypeName: &name,
			},
		})
	}

	return &data
}
