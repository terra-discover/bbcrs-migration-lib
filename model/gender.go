package model

import (
	"strconv"
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Gender Gender
type Gender struct {
	basic.Base
	basic.DataOwner
	GenderAPI
	GenderTranslation *GenderTranslation `json:"gender_translation,omitempty"`
}

// GenderAPI Gender API
type GenderAPI struct {
	GenderCode *int    `json:"gender_code,omitempty" gorm:"type:smallint;not null;index:idx_gender_code_deleted_at,unique,where:deleted_at is null" example:"1"`       // Gender Code
	GenderName *string `json:"gender_name,omitempty" gorm:"type:varchar(256);not null;index:idx_gender_name_deleted_at,unique,where:deleted_at is null" example:"Man"` // Gender Name
}

// Seed genders
func (gender *Gender) Seed() *[]Gender {
	items := []string{
		"0|Not Known",
		"1|Male",
		"2|Female",
		"9|Not Applicable",
	}

	data := []Gender{}

	for i := range items {
		content := strings.Split(items[i], "|")
		code, _ := strconv.Atoi(content[0])
		item := Gender{}
		item.GenderCode = &code
		item.GenderName = lib.Strptr(content[1])
		data = append(data, item)
	}

	return &data
}

// SeedOne Gender
func (gender *Gender) SeedOne() *Gender {
	seed := Gender{
		GenderAPI: GenderAPI{
			GenderCode: lib.Intptr(1),
			GenderName: lib.Strptr("Man"),
		},
	}
	_ = lib.Merge(seed, &gender)
	return gender
}
