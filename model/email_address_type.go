package model

import (
	"strconv"
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EmailAddressType model
type EmailAddressType struct {
	basic.Base
	basic.DataOwner
	EmailAddressTypeAPI
	EmailAddressTypeTranslation *EmailAddressTypeTranslation `json:"email_address_type_translation,omitempty"`
}

// EmailAddressTypeAPI model
type EmailAddressTypeAPI struct {
	EmailAddressTypeCode *int64  `json:"email_address_type_code,omitempty" gorm:"type:smallint;not null;index:unique_email_address_type__email_address_type_code,unique,where:deleted_at is null" example:"1"`
	EmailAddressTypeName *string `json:"email_address_type_name,omitempty" gorm:"type:varchar(256);"`
}

// Seed data
func (s EmailAddressType) Seed() *[]EmailAddressType {
	data := []EmailAddressType{}
	items := []string{
		"1|Personal",
		"2|Business",
		"3|Listserve",
		"4|Internet",
		"5|Property",
		"6|Sales office",
		"7|Reservation office",
		"8|Managing company",
		"1001|Other",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		code, _ := strconv.ParseInt(contents[0], 10, 64)
		var name string = contents[1]
		data = append(data, EmailAddressType{
			EmailAddressTypeAPI: EmailAddressTypeAPI{
				EmailAddressTypeCode: &code,
				EmailAddressTypeName: &name,
			},
		})
	}

	return &data
}
