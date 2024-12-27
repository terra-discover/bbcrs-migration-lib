package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// InformationType struct
type InformationType struct {
	basic.Base
	basic.DataOwner
	InformationTypeCode        *int64                      `json:"information_type_code,omitempty" gorm:"type:smallint;uniqueIndex:,where:deleted_at is null;not null"` // Information Type Code
	InformationTypeName        *string                     `json:"information_type_name,omitempty" gorm:"type:varchar(256)"`                                            // Information Type Name
	InformationTypeTranslation *InformationTypeTranslation `json:"information_type_translation,omitempty"`                                                              // Information Type Translation
}

// Seed data
func (s InformationType) Seed() *[]InformationType {
	data := []InformationType{}
	items := []string{
		"1|Description|",
		"2|Policy|",
		"3|Marketing|",
		"4|Special instructions|",
		"5|Other|",
		"6|Amenities|",
		"7|Attractions|",
		"8|Awards|",
		"9|Corporate locations|",
		"10|Dining|",
		"11|Driving directions|",
		"12|Facilities|",
		"13|Recreation|",
		"14|Safety|",
		"15|Services|",
		"16|Transportation|",
		"17|Short description|",
		"18|Advisory|",
		"19|Geocodes|",
		"20|Location|",
		"21|Address|",
		"22|Contact|",
		"23|Pictures|",
		"24|Descriptive content|",
		"25|Long name|",
		"26|Alias name|",
		"1001|Photo Gallery|",
		"1002|Important Notice|",
		"1003|Company Profile|",
		"1004|Corporate Contract|",
		"1005|Travel Request|",
		"1006|Splash Screen|",
		"1007|Dashboard Screen|",
		"1008|Message Attachment|",
		"1009|Passport|",
		"1010|Import Database Employee|",
	}

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		code := lib.StrToInt64(c[0])
		name := c[1]

		data = append(data, InformationType{
			InformationTypeCode: &code,
			InformationTypeName: &name,
		})
	}

	return &data
}
