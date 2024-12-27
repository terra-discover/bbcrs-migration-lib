package model

import (
	"strconv"
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProfileStatus Profile Status
type ProfileStatus struct {
	basic.Base
	basic.DataOwner
	ProfileStatusAPI
	ProfileStatusTranslation *ProfileStatusTranslation `json:"profile_status_translation,omitempty"` // Profile Status Translation
}

// ProfileStatusAPI Profile Status API
type ProfileStatusAPI struct {
	ProfileStatusCode *int    `json:"profile_status_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"123"`        // Profile Status Code
	ProfileStatusName *string `json:"profile_status_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"active"` // Profile Status Name
}

// Seed Site
func (p ProfileStatus) Seed() *[]ProfileStatus {
	data := []ProfileStatus{}
	items := []string{
		"0|deleted",
		"1|active",
		"2|draft/ inactive",
		"3|blocked",
		"4|terminated",
		"5|suspended",
		"6|banned",
	}

	for i := range items {
		values := strings.Split(items[i], "|")

		code, _ := strconv.Atoi(values[0])
		name := values[1]

		data = append(data, ProfileStatus{
			ProfileStatusAPI: ProfileStatusAPI{
				ProfileStatusCode: &code,
				ProfileStatusName: &name,
			},
		})
	}

	return &data
}
