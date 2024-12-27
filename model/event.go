package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Event Event Model
type Event struct {
	basic.Base
	basic.DataOwner
	EventAPI
	EventType  *EventType  `json:"event_type,omitempty" gorm:"foreignKey:EventTypeID;references:ID"`
	EventLevel *EventLevel `json:"event_level,omitempty" gorm:"foreignKey:EventLevelID;references:ID"`
}

// EventAPI EventAPI model
type EventAPI struct {
	EventCode        *int       `json:"event_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null"`     // Event Code
	EventName        *string    `json:"event_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Event Name
	EventTypeID      *uuid.UUID `json:"event_type_id,omitempty" swaggertype:"string" format:"uuid"`                                    // Event Type Id
	EventLevelID     *uuid.UUID `json:"event_level_id,omitempty" swaggertype:"string" format:"uuid"`                                   // Event Level Id
	ShortDescription *string    `json:"short_description,omitempty" gorm:"type:text"`                                                  // Short Description
	Description      *string    `json:"description,omitempty" gorm:"type:text"`                                                        // Description
}

type EventCode int

const (
	EventCodeError   EventCode = 1
	EventCodeSuccess EventCode = 2
	EventCodeWarning EventCode = 3
	EventCodeTimeout EventCode = 15
)

func (evc EventCode) Int() int {
	return int(evc)
}

func (evc EventCode) String() string {
	intEvc := evc.Int()
	return strconv.Itoa(intEvc)
}

// Seed data
func (Event) Seed() *[]Event {
	data := []Event{}
	items := []string{
		fmt.Sprintf("%s|Error", EventCodeError.String()),
		fmt.Sprintf("%s|Success", EventCodeSuccess.String()),
		fmt.Sprintf("%s|Warning", EventCodeWarning.String()),
		"4|Load",
		"5|Open",
		"6|Show",
		"7|Click",
		"8|Submit",
		"9|Close",
		"10|Abort",
		"11|Unload",
		"12|Share",
		"13|Send Mail",
		"14|Print",
		fmt.Sprintf("%s|Timeout", EventCodeTimeout.String()),
		"16|Insert",
		"17|Change",
		"18|Remove",
		"19|Join",
		"20|Login",
		"21|Logout",
		"22|Failed Login",
		"23|Approve",
		"24|Reject",
		"101|Flight Searched",
		"102|Hotel Searched",
		"103|No Flights Available",
		"104|No Hotels Available",
		"105|Request for Flight Proposal",
		"106|Request for Hotel Proposal",
		"107|Request for Flight Proposal Review",
		"108|Request for Hotel Proposal Review",
		"109|Review Flight Proposal",
		"110|Review Hotel Proposal",
		"111|Select Flight Proposal",
		"112|Select Hotel Proposal",
		"113|Change Flight Proposal",
		"114|Change Hotel Proposal",
		"115|Book Flight",
		"116|Book Hotel",
		"117|Approve to Issue",
		"118|Issue Ticket",
		"119|Request for Flight Cancellation",
		"120|Request for Hotel Cancellation",
		"121|Cancel Flight Proposal",
		"122|Cancel Hotel Proposal",
		"123|Cancel Flight Booking",
		"124|Cancel Hotel Booking",
		"125|Travel Request Approved",
		"126|Travel Request Rejected",
		"127|Over Credit Approved",
		"128|Over Credit Rejected",
		"129|Request for Hotel Booking Acknowledgment",
		"130|Hotel Booking Acknowledged",
		"131|Request to Reissue Ticket",
		"132|Request to Revalidate Ticket",
		"133|Request to Reissue Ticket (Name Change)",
		"134|Request to Modify Hotel Voucher",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		code, _ := strconv.Atoi(contents[0])
		var name string = contents[1]
		data = append(data, Event{
			EventAPI: EventAPI{
				EventCode: &code,
				EventName: &name,
			},
		})
	}

	return &data
}
