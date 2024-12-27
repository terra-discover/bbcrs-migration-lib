package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProfileType Profile Type
type ProfileType struct {
	basic.Base
	basic.DataOwner
	ProfileTypeAPI
	ProfileTypeTranslation *ProfileTypeTranslation `json:"profile_type_translation,omitempty"` // Profile Type Translation
}

// ProfileTypeAPI Profile Type API
type ProfileTypeAPI struct {
	ProfileTypeCode *int    `json:"profile_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`           // Profile Type Code
	ProfileTypeName *string `json:"profile_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Contact"` // Profile Type Name
}

// Seed data
func (s ProfileType) Seed() *[]ProfileType {
	data := []ProfileType{}
	items := []string{
		"1|Travel Consultant|",
		"2|Corporate Booker|",
		"3|Corporate Duty Traveler|",
		"4|Corporate Personal Traveler|",
		"5|Corporate Guest Traveler|",
		"6|Member|",
		"7|Retail Customer|",
		"8|Over Credit Approver|",
		"9|Contact|",
	}
	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		code := lib.StrToInt(c[0])
		name := c[1]

		data = append(data, ProfileType{
			ProfileTypeAPI: ProfileTypeAPI{
				ProfileTypeCode: &code,
				ProfileTypeName: &name,
			},
		})
	}

	return &data
}
