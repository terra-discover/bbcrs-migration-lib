package model

import (
	"fmt"
	"strconv"
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EventType Event Type
type EventType struct {
	basic.Base
	basic.DataOwner
	EventTypeAPI
	EventTypeTranslation *EventTypeTranslation `json:"event_type_translation,omitempty"` // Event Type Translation
}

// EventTypeAPI Event Type API
type EventTypeAPI struct {
	EventTypeCode *int    `json:"event_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`         // Event Type Code
	EventTypeName *string `json:"event_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Error"` // Event Type Name
}

type EventTypeCode int

const (
	EventTypeCodeError   EventTypeCode = 1
	EventTypeCodeSuccess EventTypeCode = 2
	EventTypeCodeWarning EventTypeCode = 3
	EventTypeCodeTimeout EventTypeCode = 15
)

func (etc EventTypeCode) Int() int {
	return int(etc)
}

func (etc EventTypeCode) String() string {
	intEtc := etc.Int()
	return strconv.Itoa(intEtc)
}

// Seed data
func (e EventType) Seed() *[]EventType {
	data := []EventType{}
	events := []string{
		fmt.Sprintf("%s|Error", EventTypeCodeError.String()),
		fmt.Sprintf("%s|Success", EventTypeCodeSuccess.String()),
		fmt.Sprintf("%s|Warning", EventTypeCodeWarning.String()),
		"4|Load",
		"5|Open",
		"6|Show",
		"7|Click",
		"8|Submit",
		"9|Close",
		"10|Abort",
		"11|Unload",
		"12|Share",
		"13|Send",
		"14|Print",
		fmt.Sprintf("%s|Timeout", EventTypeCodeTimeout.String()),
		"16|Insert",
		"17|Change",
		"18|Remove",
		"19|Join",
		"20|Login",
		"21|Logout",
		"22|Failed",
		"23|Approve",
		"24|Reject",
	}

	for i := range events {
		contents := strings.Split(events[i], "|")
		code, _ := strconv.Atoi(contents[0])
		var name string = contents[1]
		data = append(data, EventType{
			EventTypeAPI: EventTypeAPI{
				EventTypeCode: &code,
				EventTypeName: &name,
			},
		})
	}

	return &data
}
